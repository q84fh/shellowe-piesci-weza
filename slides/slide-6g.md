# Minimum naszego CLI - Flagi
Bez Vipera trzeba sobie radzić przez 'bindowanie' do zmiennych czy struktur np.:
```go
var caps, low bool

type Poem struct {
	Actor1, Actor2 string
	Say            int
}

var myPoem Poem
```

Wyróżniamy dwa rodzaje flag:
- Persistent Flags - dostępna dla rodzica jak i każdego dziecka tego rodzica. Ustawiona w `rootCmd` staje się globalnie dostępna.
- Local Flags - lokalna flaga, która dotyczy tylko tego konkretnego polecenia.

Flagi można grupować, wzajemnie wykluczać, oznaczać jako wymagane itp.

<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>