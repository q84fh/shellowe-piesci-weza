## Viper - wartości domyślne

Można ustawić je wołając po prostu `viper.SetDefault()`:
```go
viper.SetDefault("ContentDir", "/srv/app/content")
```

Lub można przypiąć klucz vipera do flagi Cobry i ustawić domyślną wartość używając Cobry:
```go
rootCmd.PersistentFlags().String("contentdir", "/srv/app/content", "directory for ur content")
viper.BindPFlag("contentdir", rootCmd.PersistentFlags().Lookup("contentdir"))
```

Można też wyliczyć domyślną wartość dynamicznie i zamiast `"/srv/app/content"` wstawić wywołanie funkcji.


<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>