## Viper - pliki konfiguracyjne zapis i obserwowanie

### Zapis
4 metody - te oznaczone *Safe* nie nadpisują istniejących plików, ale tworzą je, jeśli nie istnieją:
```go
viper.WriteConfig() // zapisuje do ścieżki ustawionej przez 'viper.AddConfigPath()' i 'viper.SetConfigName'
viper.SafeWriteConfig() // to samo, tylko nie nadpiszę pliku, jeśli już istnieje
viper.WriteConfigAs("/path/to/my/.config") // zapisze do konkretnej ścieżki
viper.SafeWriteConfigAs("/path/to/my/.config") // rzuci błędem, bo plik został stworzony w poprzednij linii i istnieje
viper.WriteConfigAs("/path/to/my/.config")  // radośnie nadpiszę plik
```

### Obserwacja
Można kazać viperowi obserwować i wczytywać dynamicznie zmiany w pliku konfiguracyjnym (linuksowy `inotify` lub odpowiednik w innych systemach):
```go
viper.OnConfigChange(func(e fsnotify.Event) {
	fmt.Println("Config file changed:", e.Name)
})
viper.WatchConfig()
```

<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>