/*
Copyright © 2023 Arboc Repiv arbocrepiv@example.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobra-demo",
	Short: "Ta aplikacja prezentuje demo biblioteki Cobra",
	Long: `Dłuży opis aplikacji może zawierać wcięcia, lepsze formatowanie
i jakieś przykłady. Na przykład:

	cobra-demo math <num1> <num2> --operation addition
	cobra-demo poem --actor1 --actor2

Ta aplikacja mogła by więcej, ale daliśmy z siebie całe 30%.`,
	Version: "1.0.0",
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var customVersion = "1.2.3"

func init() {
	rootCmd.SetVersionTemplate(fmt.Sprintf("Custom version for Goat: %s\n", customVersion))
}
