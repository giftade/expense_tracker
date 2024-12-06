package cmd

import (
	"github.com/spf13/cobra"
)

// AddExpensesCmd represents the AddExpenses command
var AddExpensesCmd = &cobra.Command{
	Use:     "Add Expenses",
	Aliases: []string{"add"},
	Short:   "Add expenses",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		AddExpenses()
	},
}

func init() {
	rootCmd.AddCommand(AddExpensesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// AddExpensesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// AddExpensesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
