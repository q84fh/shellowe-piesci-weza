package cmd

import (
	"cobra-demo/cmd/validator"
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var additionCmd = &cobra.Command{
	Use:   "addition",
	Short: "Komenda, która dodaje.",
	Args: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Argumenty: %v\n", args)
		if len(args) < 1 {
			return errors.New("wymagany minimum jeden argument")
		}
		isNum, invalidArg := validator.AreNumbers(args)
		if !isNum {
			return fmt.Errorf("nie podano liczby: %s", invalidArg)
		}

		isInt, invalidArg := validator.AreInts(args)
		if !isInt {
			return fmt.Errorf("nie podano liczby całkowitej: %s", invalidArg)
		}
		return nil
	},
	Aliases: []string{"dodaj"},
	Run: func(cmd *cobra.Command, args []string) {
		res := addition(args)
		fmt.Printf("Wynik dodawania to: %d\n", res)
	},
}

func init() {
	mathCmd.AddCommand(additionCmd)
}

func addition(args []string) (result int) {
	ints := make([]int, len(args))
	for i, s := range args {
		ints[i], _ = strconv.Atoi(s)
	}
	for _, v := range ints {
		result = result + v
	}
	return
}
