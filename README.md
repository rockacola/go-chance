# go-chance - Random generator helper for golang

A golang random generator helper inspired by [`chance.js`](https://github.com/chancejs/chancejs).

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
// Instantiate with specified seed
c := chance.NewChanceWithSeed(123456)

// All methods had a 'WithParams' variant allows feature customizations and error capture
num, err := c.NaturalWithParams(10, 100)
if (err == nil) {
    fmt.Println("Natural():", num)
}
```

Check out [`go-chance-examples`](https://github.com/rockacola/go-chance-examples) for more example usages.

## Features

| Done | Category      | Basic Method          | Advanced Method                       | Description       |
| :--: | ------------- | --------------------- | ------------------------------------- | ----------------- |
| ✓    | Basics        | `Bool()`              | `BoolWithParams(likelihood int)`      | Generate a random boolean value, either 'true' or 'false'. |
| ✓    | Basics        | `Character()`         | `CharacterWithParams(lowerCaseAlphabets bool, upperCaseAlphabets bool, numerics bool, symbols bool)` | Generate a single length string within a pool of common characters. |
| ✓    | Basics        | `Floating()`          | `FloatingWithParams(min int, max int)` | Generate a random float32 number between MinInt32 (-2,147,483,648) and MaxInt32 (2,147,483,647) |
| ✓    | Basics        | `Integer()`           | `IntegerWithParams(min int, max int)` | Generate a random int32 number between MinInt32 (-2,147,483,648) and MaxInt32 (2,147,483,647). |
| ✓    | Basics        | `Letter()`            | `LetterWithParams(lowerCaseAlphabets bool, upperCaseAlphabets bool)` | Generate a single length string within a pool of alphabetical characters. |
| ✓    | Basics        | `Natural`             | `NaturalWithParams(min int, max int)` | Generate a random int32 number between 0 and MaxInt32 (2,147,483,647). |
| ✓    | Basics        | `Prime()`             | `PrimeWithParams(min int, max int)`   | Generate a prime number between 1 and 10,000 inclusive. |
| ✓    | Basics        | `String()`            | `StringWithParams(length int, lowerCaseAlphabets bool, upperCaseAlphabets bool, numerics bool, symbols bool)` | Generate a random string with a pool of common characters. |
| ✗    | Text          | `Paragraph()`         | `TBA`                                 | TBA               |
| ✗    | Text          | `Sentence()`          | `SentenceWithParams(minWords int, maxWords int)` | Generate a semi-pronounceable nonsense word. |
| ✓    | Text          | `Syllable()`          | N/A                                   | Generate a semi-speakable syllable with 2 or 3 letters. |
| ✓    | Text          | `Word()`              | `WordWithParams(minSyllables int, maxSyllables int)` | Generate a semi-pronounceable nonsense word. |
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
| ✓    | Things        | `Animal()`            | `AnimalWithParams(category string)`   | Generate a random animal name within a pool of available animals. |
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
| ✓    | Miscellaneous | `Coin()`              | N/A                                   | Flip a coin and returns 'head' or tail'. |
| ✗    | Miscellaneous | `Guid()`              | `TBA`                                 | TBA               |
| ✓    | Miscellaneous | `Hash()`              | `HashWithParams(length int, toUpperCase bool)` | Generate a random hex hash. |
| ✓    | Miscellaneous | `Normal()`            | `NormalWithParams(mean int, deviation int)` | Return a normally-distributed random variate. |
| ✓    | Miscellaneous | `Radio()`             | `RadioWithParams(side string)`        | Generate a random radio call sign. |
| ✗    | Miscellaneous | `Rpg()`               | `TBA`                                 | TBA               |
| ✓    | Miscellaneous | `Tv()`                | `TvWithParams(side string)`           | Generate a TV station call sign. This is an alias for Radio() since they both follow the same rules. |
