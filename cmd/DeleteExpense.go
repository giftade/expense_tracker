package cmd

import (
	"github.com/spf13/cobra"
)

var id int

// DeleteExpenseCmd represents the DeleteExpense command
var DeleteExpenseCmd = &cobra.Command{
	Use:   "DeleteExpense",
	Aliases: []string{"delete"},
	Short: "Delete expense by ID",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		DeleteExpense(id)
	},
}

func init() {
	rootCmd.AddCommand(DeleteExpenseCmd)

	DeleteExpenseCmd.Flags().IntVar(&id, "id", 0, "id of expense")
}
