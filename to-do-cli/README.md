# To-Do CLI

A terminal-based task management application built with **Go**. This tool allows you to manage your daily tasks directly from your command line with persistent storage and a clean tabular interface.

## 📺 Credits & Walkthrough
This project was built following the comprehensive walkthrough by **Coding with Patrik**.
- **Video Link:** [How to Build A CLI Todo App in Go](https://youtu.be/g16Zf0KQEWI?list=PLCSt4Vg6dyVqY50J2GIr7yAYS3EWvoPLL)
- **Channel:** [Coding with Patrik](https://www.youtube.com/@codingwithpatrik)

## 🚀 Features
- **Task Management:** Add, edit, delete, and toggle completion status for tasks.
- **Persistent Storage:** Tasks are automatically saved to a `todos.json` file, ensuring data persists between sessions.
- **Clean UI:** Uses the `table` package to display task IDs, titles, status, and timestamps in a professional grid.
- **Batch Operations:** Includes a custom `drop` command to quickly clear the entire list using Go's efficient reslicing.
- **Auto-Help:** Built-in `-help` flag to guide users through available commands.

## 🛠 Commands & Usage

| Flag | Format | Description |
| :--- | :--- | :--- |
| `-add` | `title:hours` | Add a task with an estimated completion time. |
| `-list` | `N/A` | List all tasks in a table. |
| `-edit` | `id:new_title` | Change the title of a specific task. |
| `-toggle`| `id` | Mark a task as complete or incomplete. |
| `-del` | `id` | Delete a specific task by its index. |
| `-drop` | `N/A` | **Warning:** Deletes all tasks from the list. |
| `-help` | `N/A` | Show the list of all available commands. |

### Example Usage:
```bash
# Adding a task
./todo -add "Revise C++ Sliding Window:2"

# Marking task 0 as done
./todo -toggle 0

# Viewing the list
./todo -list