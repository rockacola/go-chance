# go-chance - Random generator helper for golang

Unofficial golang port of [`chance.js`](https://github.com/chancejs/chancejs).

## Get Started

Basic usage:

```go
// Create an instance of chance
c := chance.NewChance()

// Obtain a random natural number
fmt.Println("Natural():", c.Natural())
```

Advanced methods:

```go
// Instantiate with your specified seek
c := chance.NewChanceWithSeed(123456)

// All methods had a 'WithParams' variant to allow feature customizations
fmt.Println("Natural():", c.NaturalWithParams(10, 100))
```

Check out [`go-chance-examples`](https://github.com/rockacola/go-chance-examples) for more example usages.

## Features

| Done | Category      | Basic Method          | Advanced Method                       | Description       |
| :--: | ------------- | --------------------- | ------------------------------------- | ----------------- |
| ✓    | Basics        | `Bool()`              | `BoolWithParams(likelihood int)`      | Return a random boolean value. |
| ✗    | Basics        | `Character()`         | `TBA`                                 | TBA               |
| ✗    | Basics        | `Floating()`          | `TBA`                                 | TBA               |
| ✗    | Basics        | `Integer()`           | `TBA`                                 | TBA               |
| ✗    | Basics        | `Letter()`            | `TBA`                                 | TBA               |
| ✓    | Basics        | `Natural`             | `NaturalWithParams(min int, max int)` | Return a natural number. |
| ✗    | Basics        | `Prime()`             | `TBA`                                 | TBA               |
| ✗    | Basics        | `String()`            | `TBA`                                 | TBA               |
| ✗    | Text          | `Paragraph()`         | `TBA`                                 | TBA               |
| ✗    | Text          | `Sentence()`          | `TBA`                                 | TBA               |
| ✗    | Text          | `Syllable()`          | `TBA`                                 | TBA               |
| ✗    | Text          | `Word()`              | `TBA`                                 | TBA               |
| ✗    | Person        | `Age()`               | `TBA`                                 | TBA               |
| ✗    | Person        | `Birthday()`          | `TBA`                                 | TBA               |
| ✗    | Person        | `Cf()`                | `TBA`                                 | TBA               |
| ✗    | Person        | `Cpf()`               | `TBA`                                 | TBA               |
| ✗    | Person        | `First()`             | `TBA`                                 | TBA               |
| ✗    | Person        | `Gender()`            | `TBA`                                 | TBA               |
| ✗    | Person        | `Last()`              | `TBA`                                 | TBA               |
| ✗    | Person        | `Name()`              | `TBA`                                 | TBA               |
| ✗    | Person        | `Prefix()`            | `TBA`                                 | TBA               |
| ✗    | Person        | `Ssn()`               | `TBA`                                 | TBA               |
| ✗    | Person        | `Suffix()`            | `TBA`                                 | TBA               |
| ✓    | Things        | `Animal()`            | `TBA`                                 | TBA               |
| ✗    | Mobile        | `AndroidId()`         | `TBA`                                 | TBA               |
| ✗    | Mobile        | `AppleToken()`        | `TBA`                                 | TBA               |
| ✗    | Mobile        | `BlackBerryPin()`     | `TBA`                                 | TBA               |
| ✗    | Mobile        | `W7Anid()`            | `TBA`                                 | TBA               |
| ✗    | Mobile        | `W8Anid2()`           | `TBA`                                 | TBA               |
| ✗    | Web           | `Avatar()`            | `TBA`                                 | TBA               |
| ✗    | Web           | `Color()`             | `TBA`                                 | TBA               |
| ✗    | Web           | `Company()`           | `TBA`                                 | TBA               |
| ✗    | Web           | `Domain()`            | `TBA`                                 | TBA               |
| ✗    | Web           | `Email()`             | `TBA`                                 | TBA               |
| ✗    | Web           | `FacebookId()`        | `TBA`                                 | TBA               |
| ✗    | Web           | `GoogleAnalytics()`   | `TBA`                                 | TBA               |
| ✗    | Web           | `HashTag()`           | `TBA`                                 | TBA               |
| ✗    | Web           | `Ip()`                | `TBA`                                 | TBA               |
| ✗    | Web           | `Ipv6()`              | `TBA`                                 | TBA               |
| ✗    | Web           | `Klout()`             | `TBA`                                 | TBA               |
| ✗    | Web           | `Profession()`        | `TBA`                                 | TBA               |
| ✗    | Web           | `Tld()`               | `TBA`                                 | TBA               |
| ✗    | Web           | `Twitter()`           | `TBA`                                 | TBA               |
| ✗    | Web           | `Url()`               | `TBA`                                 | TBA               |
| ✗    | Location      | `Address()`           | `TBA`                                 | TBA               |
| ✗    | Location      | `Altitude()`          | `TBA`                                 | TBA               |
| ✗    | Location      | `AreaCode()`          | `TBA`                                 | TBA               |
| ✗    | Location      | `City()`              | `TBA`                                 | TBA               |
| ✗    | Location      | `Coordinates()`       | `TBA`                                 | TBA               |
| ✗    | Location      | `Country()`           | `TBA`                                 | TBA               |
| ✗    | Location      | `Depth()`             | `TBA`                                 | TBA               |
| ✗    | Location      | `GeoHash()`           | `TBA`                                 | TBA               |
| ✗    | Location      | `Latitude()`          | `TBA`                                 | TBA               |
| ✗    | Location      | `Longitude()`         | `TBA`                                 | TBA               |
| ✗    | Location      | `Phone()`             | `TBA`                                 | TBA               |
| ✗    | Location      | `Postal()`            | `TBA`                                 | TBA               |
| ✗    | Location      | `Province()`          | `TBA`                                 | TBA               |
| ✗    | Location      | `State()`             | `TBA`                                 | TBA               |
| ✗    | Location      | `Street()`            | `TBA`                                 | TBA               |
| ✗    | Location      | `Zip()`               | `TBA`                                 | TBA               |
| ✗    | Time          | `AmPm()`              | `TBA`                                 | TBA               |
| ✗    | Time          | `Date()`              | `TBA`                                 | TBA               |
| ✗    | Time          | `HammerTime()`        | `TBA`                                 | TBA               |
| ✗    | Time          | `Hour()`              | `TBA`                                 | TBA               |
| ✗    | Time          | `Millisecond()`       | `TBA`                                 | TBA               |
| ✗    | Time          | `Minute()`            | `TBA`                                 | TBA               |
| ✗    | Time          | `Month()`             | `TBA`                                 | TBA               |
| ✗    | Time          | `Second()`            | `TBA`                                 | TBA               |
| ✗    | Time          | `Timestamp()`         | `TBA`                                 | TBA               |
| ✗    | Time          | `Timezone()`          | `TBA`                                 | TBA               |
| ✗    | Time          | `Weekday()`           | `TBA`                                 | TBA               |
| ✗    | Time          | `Year()`              | `TBA`                                 | TBA               |
| ✗    | Finance       | `CreditCard()`        | `TBA`                                 | TBA               |
| ✗    | Finance       | `CreditCardType()`    | `TBA`                                 | TBA               |
| ✗    | Finance       | `Currency()`          | `TBA`                                 | TBA               |
| ✗    | Finance       | `CurrencyPair()`      | `TBA`                                 | TBA               |
| ✗    | Finance       | `Dollor()`            | `TBA`                                 | TBA               |
| ✗    | Finance       | `Euro()`              | `TBA`                                 | TBA               |
| ✗    | Finance       | `CreditCardExp()`     | `TBA`                                 | TBA               |
| ✗    | Finance       | `CreditCardExpMonth()`| `TBA`                                 | TBA               |
| ✗    | Finance       | `CreditCardExpYear()` | `TBA`                                 | TBA               |
| ✗    | Miscellaneous | `Coin()`              | `TBA`                                 | TBA               |
| ✗    | Miscellaneous | `Dice()`              | `TBA`                                 | TBA               |
| ✗    | Miscellaneous | `Guid()`              | `TBA`                                 | TBA               |
| ✗    | Miscellaneous | `Hash()`              | `TBA`                                 | TBA               |
| ✗    | Miscellaneous | `Normal()`            | `TBA`                                 | TBA               |
| ✗    | Miscellaneous | `Radio()`             | `TBA`                                 | TBA               |
| ✗    | Miscellaneous | `Rpg()`               | `TBA`                                 | TBA               |
| ✗    | Miscellaneous | `Tv()`                | `TBA`                                 | TBA               |
