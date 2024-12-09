package cmd

import (
	"github.com/spf13/cobra"
)

// DeleteExpenseCmd represents the DeleteExpense command
var DeleteExpenseCmd = &cobra.Command{
	Use:   "DeleteExpense",
	Short: "Delete expense by ID",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(DeleteExpenseCmd)

	// DeleteExpenseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
