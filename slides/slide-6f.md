# Minimum naszego CLI - init()

```go
func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
```

Tutaj definiujemy swoje flagi i ustawienia konfiguracyjne.<br>
Cobra wspiera trwałe flagi, które jak zostaną ustawione na poziomie `rootCmd` będą globalne dla aplikacji.
Cobra obsługuje również lokalne flagi, które będą uruchamiane tylko gdy ta akcja zostanie wywołana bezpośrednio.

TODO: DODAĆ FLAGI I KOD

<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>