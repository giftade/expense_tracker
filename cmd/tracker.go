package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	log.SetFlags(log.Lshortfile)
	file, err := os.ReadFile("assets/expense.json")
	if err != nil {
		log.Fatalf("err: %v", err)
		return err
	}

	var Expenses []Expense

	err = json.Unmarshal(file, &Expenses)
	if err != nil {
		log.Fatalf("err: %v", err)
		return err
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
	file, err := os.OpenFile("assets/expense.json", os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("err: %v", err)
		return err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		log.Fatalf("err: %v", err)
		return err
	}

	var Expenses []Expense

	if fileStat.Size() != 0 {
		err = json.NewDecoder(file).Decode(&Expenses)
		if err != nil {
		log.Fatalf("err: %v", err)
		return err
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
		log.Fatalf("err: %v", err)
		return err
	}

	fmt.Printf("Expense added successfully (ID: %v)", newId)
	return nil
}
