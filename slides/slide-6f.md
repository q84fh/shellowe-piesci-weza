# Minimum naszego CLI - init()

Tutaj definiujemy swoje flagi i ustawienia konfiguracyjne.

```go
func init() {
	poemCmd.PersistentFlags().StringVarP(&myPoem.Actor1, "actor1", "b", "Baba", "Aktor grający Babę")
	poemCmd.PersistentFlags().StringVar(&myPoem.Actor2, "actor2", "Dziad", "Aktor grający Dziada")
	poemCmd.Flags().BoolVarP(&caps, "caps", "c", false, "Wierz dużymi literami")
	poemCmd.Flags().BoolVarP(&low, "low", "l", false, "Wierz małymi literami")
	poemCmd.MarkFlagsMutuallyExclusive("caps", "low")
	poemCmd.Flags().IntVarP(&myPoem.Say, "say", "s", 1, "Ile razy napisać wiersz?")
	rootCmd.AddCommand(poemCmd)
}
```

Cobra wspiera trwałe flagi, które jak zostaną ustawione na poziomie `rootCmd` będą globalne dla aplikacji.<br>

Cobra obsługuje również lokalne flagi, które będą uruchamiane tylko gdy dana akcja zostanie wywołana bezpośrednio.

<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>