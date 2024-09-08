# Expense Tracker CLI Application - GO

A basic Expense Tracker CLI application written in Go

## Run Locally

Clone the project

```bash
  git clone https://github.com/hamidhasibul/expense-tracker.git
```

Go to the project directory

```bash
  cd expense-tracker
```

The application supports the following commands:

#### `help`

Displays all available commands and their usage.

```bash
./expense-tracker help
```

#### `list`

Lists all the expenses stored in the system.

```bash
./expense-tracker list
```

#### `add`

Adds a new expense to the system. Requires `--description` and `--amount` flags.

```bash
./expense-tracker add --description "Groceries" --amount 300
```

#### `delete`

Deletes an expense by ID. Requires `--id` flag.

```bash
./expense-tracker delete --id 1
```

#### `update`

Updates an expense by ID. Requires `--id`, `--description`, and `--amount` flags.

```bash
./expense-tracker update --id 1 --description "Updated Description" --amount 400
```

#### `summary`

Shows a summary of total expenses. Optionally accepts a `--month` flag to filter expenses by month (1-12).

```bash
./expense-tracker summary
./expense-tracker summary --month 7
```

## Project URL

- [roadmap.sh | expense-tracker](https://roadmap.sh/projects/expense-tracker)
