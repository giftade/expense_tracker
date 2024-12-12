package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

type Expense struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Amount      float32 `json:"amount"`
	Date        string  `json:"date"`
}

func ListExpenses() error {

	file, err := os.OpenFile("assets/expense.json", os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file stats: %v", err)
	}

	if fileStat.Size() <= 0 {
		fmt.Println("No expenses to list")
		return nil
	}

	var Expenses []Expense

	err = json.NewDecoder(file).Decode(&Expenses)
	if err != nil {
		return fmt.Errorf("failed to Decode file: %v", err)
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)

	fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", "ID", "Description", "Amount", "Date")
	for _, expense := range Expenses {
		fmt.Fprintf(w, "%v\t%v\t$%v\t%v\n", expense.ID, expense.Description, expense.Amount, expense.Date)
	}
	w.Flush()
	return nil
}

func AddExpenses(description string, amount float32) error {
	description = strings.TrimSpace(description)
	if amount < 1 {
		return fmt.Errorf("amount must be greater than 0, got: %v", amount)
	}
	if description == "" {
		return fmt.Errorf("description cannot be empty")
	}

	file, err := os.OpenFile("assets/expense.json", os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file stats: %v", err)
	}

	var Expenses []Expense

	if fileStat.Size() != 0 {
		err = json.NewDecoder(file).Decode(&Expenses)
		if err != nil {
			return fmt.Errorf("failed to decode file: %v", err)
		}
	}
	var newId int
	if len(Expenses) > 0 {
		lastExpense := Expenses[len(Expenses)-1]
		newId = lastExpense.ID + 1
	} else {
		newId = 1
	}

	newExpense := &Expense{
		ID:          newId,
		Description: description,
		Amount:      amount,
		Date:        time.Now().Format("2006-01-02"),
	}

	Expenses = append(Expenses, *newExpense)

	file.Seek(0, 0)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(Expenses)

	if err != nil {
		return fmt.Errorf("failed to encode file: %v", err)
	}

	fmt.Printf("Expense added successfully (ID: %v, Amount: %v, Description: %s)\n", newId, amount, description)
	return nil
}

func ExpenseSummary(month int) error {
	file, err := os.OpenFile("assets/expense.json", os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to get open file: %v", err)
	}

	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file stats: %v", err)
	}

	if fileStat.Size() <= 0 {
		fmt.Println("No expenses yet")
		return nil
	}

	var Expenses []Expense

	err = json.NewDecoder(file).Decode(&Expenses)
	if err != nil {
		return fmt.Errorf("failed to decode file: %v", err)
	}
if month < 0 ||month > 12 {
	fmt.Printf("month %v does not exist", month)
	return nil
}
	if month >= 1 && month <= 12 {
		var totalExpense float32
		
		for _, expense := range Expenses {
			expenseDate := strings.Split(expense.Date, "-")

			expenseMonth := expenseDate[1]
			expenseMonthInt, _ := strconv.Atoi(expenseMonth)

			if month == expenseMonthInt {
				totalExpense += expense.Amount
			}
	}

	var newMonth string

	switch month {
	case 1:
		newMonth = "January"
	case 2:
		newMonth = "February"
	case 3:
		newMonth = "March"
	case 4:
		newMonth = "April"
	case 5:
		newMonth = "May"
	case 6:
		newMonth = "June"
	case 7:
		newMonth = "July"
	case 8:
		newMonth = "August"
	case 9:
		newMonth = "September"
	case 10:
		newMonth = "October"
	case 11:
		newMonth = "November"
	case 12:
		newMonth = "December"
	}

	fmt.Printf("Total expenses for %s: $%v\n", newMonth, totalExpense)
	return nil
}

	var totalExpense float32

	for _, expense := range Expenses {

		totalExpense += expense.Amount
	}

	fmt.Printf("Total expenses: $%v", totalExpense)
	return nil
}

func DeleteExpense(id int) error {
	file, err := os.OpenFile("assets/expense.json", os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file stats: %v", err)
	}

	if fileStat.Size() <= 0 {
		fmt.Println("No expenses to delete yet")
		return nil
	}

	var expenses []Expense
	if err := json.NewDecoder(file).Decode(&expenses); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	var newExpenses []Expense
	found := false
	for _, expense := range expenses {
		if expense.ID == id {
			found = true
			continue
		}
		newExpenses = append(newExpenses, expense)
	}

	if !found {
		fmt.Printf("Expense with ID: %v not found\n", id)
		return nil
	}

	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate file: %v", err)
	}
	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek to start of file: %v", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(newExpenses); err != nil {
		return fmt.Errorf("failed to encode JSON: %v", err)
	}

	fmt.Printf("Expense with ID: %v deleted successfully\n", id)
	return nil
}
