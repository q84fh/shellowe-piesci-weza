## Viper - pliki konfiguracyjne odczyt

Potrzebna jest minimalna konfiguracja, żeby Viper wiedział gdzie i czego szukać - nie ma domyślnej.
```go
viper.SetConfigName("config") // nazwa pliku bez rozszerzenia
viper.SetConfigType("yaml") // typ 
viper.AddConfigPath("/etc/appname/")   // ścieżki gdzie możemy spodziewać się konfiguracji
viper.AddConfigPath("$HOME/.appname")  // można dodać wiele
viper.AddConfigPath(".")               // i opcjonalnie katalog roboczy też, kolejność ma znaczenie
if err := viper.ReadInConfig(); err != nil {
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		//  konfiguracji nie ma, ale możemy to zignorować jeżeli mamy dobre domyślne wartości
	} else {
		// konfiguracja znaleziona, ale coś innego nie poszło
	}
}
```

<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>