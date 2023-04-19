# Struktura, instalacja i import

Aby użyć Cobry, należy odwzorować strukturę poleceń w plikach go. Dla klasycznej aplikacji CLI może to wyglądać na przykład tak:

```
.
├── cmd
│   ├── root.go
│   ├── cmdsub1.go
│   └── cmdsub2.go
└── main.go
```

W `root.go` ulokujemy *najniższy* poziom naszych komend, czyli `rootCmd` naszej aplikacji CLI. W sąsiadujących plikach zawrzemy podkomendy.

***

Instalacja najnowszej wersji:
```
go get -u github.com/spf13/cobra/cobra
```
Import w aplikacji:
```
import "github.com/spf13/cobra"
```


<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>