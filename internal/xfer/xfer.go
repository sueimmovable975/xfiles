// Package xfer holds the file-transfer core shared by xftp's interactive REPL
// and the xcp one-shot command: a crash-safe download (temp file renamed on
// success), a chunked upload, and a progress display. Keeping it here means both
// binaries move bytes the same way and a fix lands in both at once.
package xfer

import (
	"context"
	"fmt"
	"io"
	"mime"
	"os"
	"path"
	"path/filepath"

	"github.com/excelano/xftp/internal/drive"
	"github.com/excelano/xftp/internal/spauth"
)

// progressThreshold is the file size above which a transfer prints a progress
// line. Below it, transfers are quick enough that progress would just flicker.
const progressThreshold = 50 * 1024 * 1024 // 50 MiB

// Download streams the remote file at library-relative path `remote` to
// localPath. It writes through a temp file in the destination directory and
// renames into place only on success, so an interrupted or failed transfer
// never leaves a corrupt file at the real name. Progress prints to stderr for
// files over progressThreshold. A remote folder is rejected. When the context
// is cancelled (Ctrl-C), the returned error wraps ctx.Err().
func Download(ctx context.Context, g *spauth.GraphClient, d *drive.Drive, remote, localPath string) error {
	item, err := d.Stat(ctx, g, remote)
	if err != nil {
		return err
	}
	if item.IsFolder {
		return fmt.Errorf("/%s is a folder", remote)
	}

	tmp, err := os.CreateTemp(filepath.Dir(localPath), "."+filepath.Base(localPath)+".part-*")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()

	var w io.Writer = tmp
	if item.Size > progressThreshold {
		w = &progressWriter{w: tmp, total: item.Size, label: path.Base(remote)}
		defer fmt.Fprintln(os.Stderr)
	}

	dlErr := d.Download(ctx, g, remote, w)
	if cerr := tmp.Close(); cerr != nil && dlErr == nil {
		dlErr = cerr
	}
	if dlErr != nil {
		os.Remove(tmpName)
		if ctx.Err() != nil {
			return fmt.Errorf("interrupted: %w", ctx.Err())
		}
		return dlErr
	}
	if err := os.Rename(tmpName, localPath); err != nil {
		os.Remove(tmpName)
		return err
	}
	return nil
}

// Upload streams the local file at localPath to the library-relative path
// `remote`. Files at or below the simple-upload ceiling go in a single PUT;
// larger ones use a chunked, resumable session (see drive.Upload). Progress
// prints to stderr for files over progressThreshold. It returns the number of
// bytes sent. A local directory is rejected. On ctx cancellation the returned
// error wraps ctx.Err().
func Upload(ctx context.Context, g *spauth.GraphClient, d *drive.Drive, localPath, remote string) (int64, error) {
	f, err := os.Open(localPath)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return 0, err
	}
	if info.IsDir() {
		return 0, fmt.Errorf("%s is a directory", localPath)
	}

	var r io.Reader = f
	if info.Size() > progressThreshold {
		r = &progressReader{r: f, total: info.Size(), label: path.Base(remote)}
		defer fmt.Fprintln(os.Stderr)
	}

	ctype := mime.TypeByExtension(filepath.Ext(localPath))
	if err := d.Upload(ctx, g, remote, ctype, r, info.Size()); err != nil {
		if ctx.Err() != nil {
			return 0, fmt.Errorf("interrupted: %w", ctx.Err())
		}
		return 0, err
	}
	return info.Size(), nil
}

// progressReader wraps a reader and prints a single rewritten line of upload
// progress to stderr as bytes flow through it. The label is the remote file name.
type progressReader struct {
	r     io.Reader
	total int64
	read  int64
	label string
}

func (p *progressReader) Read(b []byte) (int, error) {
	n, err := p.r.Read(b)
	p.read += int64(n)
	pct := 0.0
	if p.total > 0 {
		pct = float64(p.read) / float64(p.total) * 100
	}
	fmt.Fprintf(os.Stderr, "\ruploading %s: %d/%d bytes (%.0f%%)", p.label, p.read, p.total, pct)
	return n, err
}

// progressWriter is the download counterpart of progressReader: it prints a
// rewritten progress line as bytes are written to the local file.
type progressWriter struct {
	w     io.Writer
	total int64
	wrote int64
	label string
}

func (p *progressWriter) Write(b []byte) (int, error) {
	n, err := p.w.Write(b)
	p.wrote += int64(n)
	pct := 0.0
	if p.total > 0 {
		pct = float64(p.wrote) / float64(p.total) * 100
	}
	fmt.Fprintf(os.Stderr, "\rdownloading %s: %d/%d bytes (%.0f%%)", p.label, p.wrote, p.total, pct)
	return n, err
}
