# Minimum naszego CLI

Tu się wszysto zaczyna:
```go{5}
// main.go
package main
import "cobra-demo/cmd"
func main() {
	cmd.Execute()
}
```
Wywoływana jest główna funkcja naszego CLI: `cmd.Execute()`


```go
// cmd/root.go
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
```
`rootCmd` reprezentuje polecenie podstawowe, gdy jest wywoływane bez żadnych podpoleceń.<br>
Execute dodaje wszystkie podpolecenia do polecenia głównego i odpowiednio ustawia flagi.<br>
Wykonuje się to tylko raz dla `rootCmd`.


<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>