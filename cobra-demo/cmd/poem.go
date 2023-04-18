package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var caps, low bool

type Poem struct {
	Actor1, Actor2 string
	Say            int
}

var myPoem Poem

var poemCmd = &cobra.Command{
	Use:     "poem",
	Short:   "Zostań bohaterem opowieści",
	Long:    `Możesz zostać bohaterem opowieści Ty i ktoś inny podając imiona aktorów.`,
	Example: `cobra-demo poem --actor1 Janina --actor2 Jan`,
	Run: func(cmd *cobra.Command, args []string) {
		myPoem.poem(caps, low)
	},
	SuggestFor: []string{"wiersz"},
}

func init() {
	poemCmd.PersistentFlags().StringVarP(&myPoem.Actor1, "actor1", "b", "Baba", "Aktor grający Babę")
	poemCmd.PersistentFlags().StringVar(&myPoem.Actor2, "actor2", "Dziad", "Aktor grający Dziada")

	poemCmd.Flags().BoolVarP(&caps, "caps", "c", false, "Wierz dużymi literami")
	poemCmd.Flags().BoolVarP(&low, "low", "l", false, "Wierz małymi literami")
	poemCmd.MarkFlagsMutuallyExclusive("caps", "low")

	poemCmd.Flags().IntVarP(&myPoem.Say, "say", "s", 1, "Ile razy napisać wiersz?")

	rootCmd.AddCommand(poemCmd)
}

func (p *Poem) poem(caps bool, low bool) {
	// Przed tym powinna ochronić nas MarkFlagsMutuallyExclusive na flagi "caps" i "low"
	if caps && low {
		fmt.Println("Dziad zapomniał, a Baba nie chce nawet wiedzieć.")
		return
	}

	for i := 0; i < p.Say; i++ {
		wiersz := fmt.Sprintf("Siała %s mak\nNie wiedziała jak\nA %s wiedział\nNie powiedział\nA to było tak\n", p.Actor1, p.Actor2)
		if caps {
			wiersz = strings.ToUpper(wiersz)
		}
		if low {
			wiersz = strings.ToLower(wiersz)
		}
		fmt.Print(wiersz)
	}
}
