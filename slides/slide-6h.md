# Minimum naszego CLI - Argumenty

Argumenty przekazane do polecenia są dostępne pod `args - []string{}`.

Wbudowane są następujące walidatory:

- NoArgs - polecenie zgłosi błąd, jeśli istnieją jakiekolwiek argumenty pozycyjne.
ArbitraryArgs - polecenie przyjmie dowolne argumenty.
- OnlyValidArgs - polecenie zgłosi błąd, jeśli istnieją jakiekolwiek argumenty pozycyjne, które nie znajdują się w polu ValidArgs polecenia.
- MinimumNArgs(int) - polecenie zgłosi błąd, jeśli nie ma co najmniej N argumentów pozycyjnych.
- MaximumNArgs(int) - polecenie zgłosi błąd, jeśli istnieje więcej niż N argumentów pozycyjnych.
- ExactArgs(int) - polecenie zgłosi błąd, jeśli nie istnieje dokładnie N argumentów pozycyjnych.
- ExactValidArgs(int) = polecenie zgłosi błąd, jeżeli nie istnieje dokładnie N argumentów pozycyjnych LUB jeżeli istnieją jakiekolwiek argumenty pozycyjne, które nie znajdują się w polu ValidArgs polecenia
- RangeArgs(min, max) - polecenie zgłosi błąd, jeśli liczba argumentów nie mieści się pomiędzy minimalną i maksymalną liczbą oczekiwanych argumentów.

<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>