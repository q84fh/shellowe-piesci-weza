## 12factors - WTF?

[Dwanaście aspektów aplikacji](https://12factor.net/pl/) jest metodologią budowania aplikacji.

### Aspekt 3: Przechowuj konfigurację w środowisku
Konfiguracja to jedyny element, który może się różnić pomiędzy wdrożeniami aplikacji (staging, produkcja, środowisko developerskie, etc). 

W jej skład wchodzą wartości różne dla każdego wdrożenia, jak np. kanoniczna nazwa hosta, dane połączenia do usług, dane uwierzytelniające etc.

Zasady:
 - konfiguracja jest ściśle oddzielona od kodu aplikacji,
 - aplikacja przechowuje konfigurację w zmiennych środowiskowych, 
 - zmienne środowiskowe służą do precyzyjnej kontroli poszczególnych ustawień, posiadając różne, nie mylące się ze sobą nazwy. Nigdy nie są zgrupowane w “środowiskach”, tylko niezależnie ustawiane dla każdego wdrożenia. 


<!-- Copy this block for every slide -->
<BarBottom  title="Goat - Poznań Go Devs #7">
  <Item text="Meetup">
    <a href="https://www.meetup.com/pl-PL/goat-poznan-go-devs/"><img src="/images/meetup-icon.svg" class="w-5"/></a>
  </Item>
</BarBottom>