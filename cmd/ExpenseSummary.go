package cmd

import (
	"github.com/spf13/cobra"
)
var month int

// ExpenseSummaryCmd represents the ExpenseSummary command
var ExpenseSummaryCmd = &cobra.Command{
	Use:     "Expense Summary",
	Aliases: []string{"summary"},
	Short:   "Gives an expense  summary",
	Long: `summary for expenses, can give summary for a particular month with the --month flag eg:  expense-tracker summary --month 8 with a response like: # Total expenses for August: $20
`,
	Run: func(cmd *cobra.Command, args []string) {
		ExpenseSummary(month)
	},
}

func init() {
	
	ExpenseSummaryCmd.Flags().IntVar(&month, "month", 0, "expense for the month")

	rootCmd.AddCommand(ExpenseSummaryCmd)

}
