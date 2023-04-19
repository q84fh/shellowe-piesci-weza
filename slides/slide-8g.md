 ## Viper - zmienne środowiskowe

### Prefix 
Można ustawić prefix przed zmiennymi środowiskowymi, żeby zapewnić ich unikalność: `SetEnvPrefix("mojprefix)` i nie wejść 
w konflikt z istniejącymi zmiennymi w środowisku.

### Ręcznie
`BindEnv(string...) : error`  - pierwszy argument to klucz w Viperze, kolejne (opcjonalne) to nazwy zmiennych środowiskowych, kolejność jest ważna. Jeżeli podany tylko pierwszy argument Viper zakłada, że powiązana zmienna środowiskowa ma format `PREFIX + "_" + NAZWA` (z wielkich liter)

### Automatycznie
`AutomaticEnv()` - powoduje, że przy każdym odwołaniu do wartości środowisko jest przeszukiwane pod kątem obecności zmiennych środowiskowych w formacie `PREFIX + "_" + NAZWA` (z wielkich liter) - nie trzeba ręcznie bindować.

<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>