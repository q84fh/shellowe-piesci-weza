## Viper: kolejność

Kolejność na wypadek, gdyby klucz występował w wielu źródłach (najważniejsze na górze):

- bezpośrednie zawołanie `viper.Set()`
- flagi z CLI
- zmienne środowiskowe
- pliki konfiguracyjne
- systemy klucz/wartość (Etcd etc.)
- domyślne wartości

Gwiazdki:
- (obecnie) klucze są nieczułe na wielkość liter
- można wczytać jeden plik konfiguracyjny na raz
  - ... ale można go poszukiwać w wielu ścieżkach posortowanych ważnością
  - ... można mieć też wiele Viperów w jednej aplikacji

<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>