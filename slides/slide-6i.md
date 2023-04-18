# Sugestie przy błędnych komendach

Cobra będzie drukować automatyczne sugestie używające implementacji [odległości Levenshteina](https://pl.wikipedia.org/wiki/Odleg%C5%82o%C5%9B%C4%87_Levenshteina), gdy wystąpią błędy "nieznanego polecenia".

```bash
$ ./cobra-demo dance
Error: unknown command "dance" for "cobra-demo"
Run 'cobra-demo --help' for usage
```

Wyłączenie sugestii:
```go
rootCmd.DisableSuggestions = true
```

Dostosowanie odległości łańcuchów w swoim poleceniu:
```go
rootCmd.SuggestionsMinimumDistance = 1
```

<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>