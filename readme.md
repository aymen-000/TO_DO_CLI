# TodoCLI

A simple, user-friendly Command Line Interface (CLI) tool for managing a Todo list. This project is built with Go and includes a MySQL backend for managing tasks. The `TodoCLI` script allows you to build, run, install, and manage the application effortlessly.

---

## Features

- Manage todos (add, view, update, delete) via a CLI tool.
- Persistent storage using a MySQL database.
- Simple setup with automated database creation and table setup.
- Includes build, install, and uninstall scripts for ease of use.

---

## Prerequisites

Before you start, ensure you have the following installed:

- [Go](https://golang.org/) (v1.16 or later)
- [MySQL](https://www.mysql.com/)
- Bash shell (Unix-based systems)

---
## Database Structure

The `todo` table includes the following fields:

| Field         | Type         | Description                        |
|---------------|--------------|------------------------------------|
| `id`          | INT          | Unique identifier for each task.  |
| `title`       | VARCHAR(255) | Title of the todo task.           |
| `completed`   | BOOLEAN      | Task status (default: `FALSE`).   |
| `created_at`  | DATETIME     | Timestamp when the task was added.|
| `completed_at`| DATETIME     | Timestamp when the task was completed. |

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/aymen-000/TO_DO_CLI.git
cd TO_DO_CLI
``` 

### 2. Set Env varaibles 
```  bash 

export DBUSER="your_database_user"
export DBPASS="your_database_password"
``` 
### 3. Setup database 
```  bash 
mysql -u root -p 
sql > CREATE DATABASE TODOS
sql > use TODOS 
sql > CREATE USER 'username'@'localhost' IDENTIFIED BY 'password';
sql > GRANT ALL PRIVILEGES ON TODOS.* TO 'username'@'localhost';
sql > FLUSH PRIVILEGES;

```
### 4. Run  
```  bash 
mv install.bash /usr/local/bin/todocli 
sudo chmod +x /usr/local/bin/todocli
todocli help 
todocli setup-db
todocli build
todocli install
todocli run -help 
``` 

