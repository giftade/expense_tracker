# Expense Tracker

Expense Tracker is a CLI tool for tracking expenses. It allows you to add, list, delete, and summarize expenses.

## Features

- Add new expenses
- List all expenses
- Delete an expense by ID
- Summarize expenses by month

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/giftade/expense_tracker
   ```
2. Navigate to the project directory:
   ```sh
   cd expense-tracker
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage

### Add an Expense

To add a new expense, use the `AddExpenses` command:

```sh
go run main.go add --description "Description" --amount 100.50
```
### List Expenses
To list all expenses, use the ListExpenses command:
```sh
go run main.go list
```

### Delete an Expense
To delete an expense by ID, use the DeleteExpense command:
```sh
go run main.go delete --id 1
```

### Summarize Expenses
To summarize expenses by month, use the ExpenseSummary command:
```sh
go run main.go summary --month 1
```

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any changes.