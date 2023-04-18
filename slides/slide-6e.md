# Minimum naszego CLI - rootCmd

```go
var rootCmd = &cobra.Command{
	Use:   "cobra-demo",
	Short: "Ta aplikacja prezentuje demo biblioteki Cobra",
	Long: `Dłuży opis aplikacji może zawierać wcięcia, lepsze formatowanie
i jakieś przykłady. Na przykład:

	cobra-demo math <num1> <num2> --operation addition
	cobra-demo poem --actor1 --actor2

Ta aplikacja mogła by więcej, ale daliśmy z siebie całe 30%.`,
	// Run: func(cmd *cobra.Command, args []string) {},
}
```
Funkcje typu `Run` są uruchamiane w następującej kolejności:

- PersistentPreRun
- PreRun
- Run
- PostRun
- PersistentPostRun

<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>