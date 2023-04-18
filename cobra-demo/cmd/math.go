package cmd

import (
	"github.com/spf13/cobra"
)

var mathCmd = &cobra.Command{
	Use:   "math",
	Short: "Obsługa funkcji matematycznych",
}

func init() {
	rootCmd.AddCommand(mathCmd)
}
