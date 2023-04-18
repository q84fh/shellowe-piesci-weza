# Budowa CLI

Wywołujemy plik wykonywalny, a następnie rzeczywiste polecenie (czasownik), podkomendy, argumenty, flagi.

<br>

```
appctl {komenda} {podkomenda} {argumenty...} {flagi...}
```
albo też:
```
appctl {czasownik} {zasób} {nazwa-zasobu} {parametry...} {flagi...}
```
<br>

* Przykłady:
```
kubectl autoscale deployment foo --min = 2 --max = 10

kubectl config delete-cluster <nazwa klastra>

gh repo clone https://github.com/cli/cli

gh pr create --title "Pull request title" --body "Pull request body"

http://github.com/owner/repo/pull/1
```

<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>