package cmd

import (
	"encoding/json"
	"fmt"
	"log"
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
  
   fmt.Fprintf(w,"%v\t%v\t%v\t%v\n", "ID", "Description", "Amount", "Date")
  for _, expense := range Expenses {
    fmt.Fprintf(w,"%v\t%v\t$%v\t%v\n", expense.ID, expense.Description, expense.Amount, expense.Date)
  }
  w.Flush()
	return nil
}

func AddExpenses() error{
  file, err := os.OpenFile("assets/expense.json", os.O_RDWR, 0644)
 if err != nil {
		log.Fatalf("1err: %v", err)
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
		log.Fatalf("2err: %v", err)
		return err
	}
}
  


  newExpense := &Expense{
    ID: 3,
    Description: "School",
    Amount: 80,
    Date:"2024-12-06",
  }

  Expenses = append(Expenses, *newExpense)

  file.Seek(0, 0)

  encoder := json.NewEncoder(file)
  encoder.SetIndent("", " ")
  err = encoder.Encode(Expenses)
  
  if err != nil {
		log.Fatalf("3err: %v", err)
		return err
	}


  fmt.Println("Expense added")
  return nil
}