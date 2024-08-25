<h1 align="center">goundo</h1>
<h3 align="center">A CLI tool for recovering files accidentally deleted with commands like `rm`.</h3>

## Overview

**goundo** is a command-line utility designed to help users recover files that were accidentally deleted using commands like `rm`.

Instead of permanently losing important files, **goundo** provides a way to undo the deletion by creating backups that can be restored when needed.

## Features

- **Undo File Deletion**: Easily recover files or directories that were accidentally deleted.
- **Automatic Backups**: **goundo** automatically backs up files before deletion, ensuring that they can be restored if needed.
- **Comprehensive Command Interface**: **goundo** offers a range of commands, including `list`, `restore` for file recovery and various others to suit different needs.
- **Configurable Backup Location**: Specify where your backups are stored, giving you control over your recovery options.
- **Lightweight and Fast**: **goundo** is designed to have minimal impact on your system while providing powerful recovery capabilities.

## Installation

To get started with **goundo**, follow these steps to install the CLI app and configure it to intercept `rm` commands.

### 1. Install the `goundo` CLI

First, install the **goundo** CLI application by running the following command:
```bash
go install github.com/alicse3/goundo@latest
```
This command will download and install the **goundo** binary to your Go bin directory, making it accessible from your terminal.

### 2. Configure rm Alias
To ensure that all file deletions are handled by **goundo**, you can create an alias for the `rm` command. This alias will redirect any `rm` command to **goundo**, allowing you to take advantage of its file recovery features automatically.

Add the following line to your shell configuration file (e.g., **.bashrc**, **.zshrc**, or .**bash_profile**):
```bash
alias rm="goundo rm"
```

After adding the alias, refresh your shell session by running:
```bash
source ~/.zshrc  # or source ~/.bashrc depending on your shell
```

Now, whenever you use the `rm` command, it will be intercepted by **goundo**, providing you with the ability to recover deleted files easily.

### 3. Verify Installation
To confirm that the installation was successful and the alias is working correctly, you can check the version of **goundo**:
```bash
goundo version
```

## Usage

### version
Check the current version of the **goundo** CLI:
```bash
goundo version
```

### Help
Get a list of available commands and options:
```bash
goundo help
```

### Configure
Configure the app settings, such as specifying the backup directory or other preferences:
```bash
goundo configure
```

### List
Display information about existing backups, including file names, backup IDs, and timestamps:
```bash
goundo list
```

### Restore
Restore files from backups using one or more backup IDs. You can find the backup IDs by running the `list` command:
```bash
goundo restore <backup_id1>,<backup_id2>,<backup_id3> ...
```

## Contributing

Contributions to **goundo** are welcome and encouraged! If you'd like to contribute, please follow these steps:

1. **Fork the Repository**  
   Create a fork of the **goundo** repository on GitHub to work on your changes.

2. **Make Your Changes**  
   Make your changes in your forked repository. 

3. **Submit a Pull Request**  
   Once you've made your changes, open a pull request against the original **goundo** repository. Provide a clear description of what you've done and any relevant information.

We appreciate your contributions to **goundo** and look forward to reviewing your pull requests!

## License
This project is licensed under the MIT License. See the `LICENSE` file in the project root directory for more details.

## Contact
For any questions or support, please contact alicse3@gmail.com.

