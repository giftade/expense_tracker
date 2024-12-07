package cmd

import (
	"github.com/spf13/cobra"
)

var description string
var amount float32

// AddExpensesCmd represents the AddExpenses command
var AddExpensesCmd = &cobra.Command{
	Use:     "Add Expenses",
	Aliases: []string{"add"},
	Short:   "Add expenses",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		AddExpenses(description, amount)
	},
}

func init() {
	AddExpensesCmd.Flags().StringVar(&description, "description", "", "description of expense")

	AddExpensesCmd.Flags().Float32Var(&amount, "amount", 0, "cost of expense")

	
	rootCmd.AddCommand(AddExpensesCmd)

}
