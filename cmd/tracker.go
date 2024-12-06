package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
)

type Expense struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Amount      float32 `json:"amount"`
	Date        string  `json:"date"`
}

func ListExpenses() error {
	file, err := os.ReadFile("assets/expense.json")
	if err != nil {
		fmt.Printf("err: %v", err)
		return err
	}
	var Expenses []Expense

	err = json.Unmarshal(file, &Expenses)
  if err != nil {
    fmt.Printf("err: %v", err)
		return err
  }
  w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
  
   fmt.Fprintf(w,"%v\t%v\t%v\t%v\n", "ID", "Description", "Amount", "Date")
  for _, expense := range Expenses {
    fmt.Fprintf(w,"%v\t%v\t$%v\t%v\n", expense.ID, expense.Description, expense.Amount, expense.Date)
  }
  w.Flush()
	return nil
}
