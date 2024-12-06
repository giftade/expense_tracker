package cmd

import (
	"github.com/spf13/cobra"
)

// ListExpensesCmd represents the ListExpenses command
var ListExpensesCmd = &cobra.Command{
	Use:     "List Expenses",
	Aliases: []string{"list"},
	Short:   "List all expenses",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		ListExpenses()
	},
}

func init() {
	rootCmd.AddCommand(ListExpensesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ListExpensesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ListExpensesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
