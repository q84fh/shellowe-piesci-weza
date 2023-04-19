## Viper - dobieranie się do wartości

Płaskie:
```go
viper.GetString("logfile") // case-insensitive Setting & Getting
if viper.GetBool("verbose") {
	fmt.Println("verbose enabled")
}
```

Zagnieżdżone:
```json
{
    "host": {
        "address": "localhost",
        "ports": [5799, 1337],
    }
}
```
```go
GetString("host.address") // (zwraca localhost)
GetInt("host.ports.2") // (zwraca 1337)
```

<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>