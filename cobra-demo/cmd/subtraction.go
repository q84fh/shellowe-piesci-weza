package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sub1, sub2, sub3 int

var subtractionCmd = &cobra.Command{
	Use:   "subtraction",
	Short: "Komenda która odejmuje",
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Będzie odejmowane zaraz!")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("subtraction called")
		res := subtraction(sub1, sub2, sub3)
		fmt.Printf("Wynik odejmowania to: %d\n", res)
	},
}

func init() {
	subtractionCmd.Flags().IntVar(&sub1, "sub1", 0, "Pierwsza liczba")
	subtractionCmd.Flags().IntVar(&sub2, "sub2", 0, "Druga liczba")
	subtractionCmd.Flags().IntVar(&sub3, "sub3", 0, "Trzecia ukryta liczba")
	subtractionCmd.Flags().MarkHidden("sub3") // tego schowamy
	mathCmd.AddCommand(subtractionCmd)
}

func subtraction(sub1 int, sub2 int, sub3 int) (result int) {
	return sub1 - sub2 - sub3
}
