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
| ✓    | Basics        | `Floating()`          | `FloatingWithParams(min int, max int)` | Generate a random float64 number between MinInt32 (-2,147,483,648) and MaxInt32 (2,147,483,647) |
| ✓    | Basics        | `Integer()`           | `IntegerWithParams(min int, max int)` | Generate a random int32 number between MinInt32 (-2,147,483,648) and MaxInt32 (2,147,483,647). |
| ✓    | Basics        | `Letter()`            | `LetterWithParams(lowerCaseAlphabets bool, upperCaseAlphabets bool)` | Generate a single length string within a pool of alphabetical characters. |
| ✓    | Basics        | `Natural`             | `NaturalWithParams(min int, max int)` | Generate a random int32 number between 0 and MaxInt32 (2,147,483,647). |
| ✓    | Basics        | `Prime()`             | `PrimeWithParams(min int, max int)`   | Generate a prime number between 1 and 10,000 inclusive. |
| ✓    | Basics        | `String()`            | `StringWithParams(length int, lowerCaseAlphabets bool, upperCaseAlphabets bool, numerics bool, symbols bool)` | Generate a random string with a pool of common characters. |
| ✓    | Text          | `Paragraph()`         | `ParagraphWithParams(minSentences int, maxSentences int)` | Generate a paragraph of semi-pronounceable nonsense word. |
| ✓    | Text          | `Sentence()`          | `SentenceWithParams(minWords int, maxWords int)` | Generate a sentence of semi-pronounceable nonsense word. |
| ✓    | Text          | `Syllable()`          | N/A                                   | Generate a semi-speakable syllable with 2 or 3 letters. |
| ✓    | Text          | `Word()`              | `WordWithParams(minSyllables int, maxSyllables int)` | Generate a semi-pronounceable nonsense word. |
| ✓    | Person        | `Age()`               | `AgeWithParams(category string)`      | Generate a random age. |
| ✗    | Person        | `Birthday()`          | `TBA`                                 | TBA               |
| ✗    | Person        | `Cf()`                | `TBA`                                 | TBA               |
| ✓    | Person        | `Cpf()`               | N/A                                   | Generate a random Brazilian tax identifier. |
| ✓    | Person        | `FirstName()`         | `FirstNameWithParams(gender string, nationality string)` | Generate a random first name. |
| ✓    | Person        | `Gender()`            | `GenderWithparams(extraGenders []string)` | Generate a random gender. |
| ✓    | Person        | `LastName()`          | `LastNameWithParams(nationality string)` | Generate a random last name. |
| ✓    | Person        | `Name()`              | `NameWithParams(middleName bool, prefix bool, suffix bool, gender string, nationality string)` | Generate a random name. |
| ✓    | Person        | `NamePrefix()`        | `NamePrefixWithParams(full bool, gender string)` | Generate a random name prefix. |
| ✓    | Person        | `Ssn()`               | `SsnWithParams(ssnFour bool, dashes bool)` | Generate a random social security number. |
| ✓    | Person        | `NameSuffix()`        | `NamePrefixWithParams(full bool)`     | Generate a random name suffix. |
| ✓    | Things        | `Animal()`            | `AnimalWithParams(category string)`   | Generate a random animal name within a pool of available animals. |
| ✓    | Mobile        | `AndroidId()`         | N/A                                   | Generate a random Android GCM Registration identifier. |
| ✓    | Mobile        | `AppleToken()`        | N/A                                   | Generate a random Apple Push Token. |
| ✓    | Mobile        | `BlackBerryPin()`     | N/A                                   | Generate a random BlackBerry Device PIN. |
| ✓    | Mobile        | `Wp7Anid()`           | N/A                                   | Generate a random Windows Phone 7 ANID. |
| ✓    | Mobile        | `Wp8Anid2()`          | N/A                                   | Generate a random Windows Phone 8 ANID2. |
| ✗    | Web           | `Avatar()`            | `TBA`                                 | TBA               |
| ✗    | Web           | `Color()`             | `TBA`                                 | TBA               |
| ✓    | Web           | `Company()`           | N/A                                   | Generate a random company name. |
| ✓    | Web           | `Domain()`            | `DomainWithParams(tld string)`        | Generate a random domain name with TLD. |
| ✓    | Web           | `Email()`             | `EmailWithParams(domain string)`      | Generate a random email address. |
| ✓    | Web           | `FacebookId()`        | N/A                                   | Generate a random Facebook ID. |
| ✓    | Web           | `GoogleAnalytics()`   | N/A                                   | Generate a random Google Analytics tracking code. |
| ✗    | Web           | `HashTag()`           | `TBA`                                 | TBA               |
| ✗    | Web           | `Ip()`                | `TBA`                                 | TBA               |
| ✗    | Web           | `Ipv6()`              | `TBA`                                 | TBA               |
| ✗    | Web           | `Klout()`             | `TBA`                                 | TBA               |
| ✗    | Web           | `Profession()`        | `TBA`                                 | TBA               |
| ✓    | Web           | `Tld()`               | N/A                                   | Generate a random TLD. |
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
| ✓    | Miscellaneous | `Guid()`              | `GuidWithParams(version int)`         | Generate a random Globally Unique Identifier. |
| ✓    | Miscellaneous | `Hash()`              | `HashWithParams(length int, toUpperCase bool)` | Generate a random hex hash. |
| ✓    | Miscellaneous | `Normal()`            | `NormalWithParams(mean int, deviation int)` | Return a normally-distributed random variate. |
| ✓    | Miscellaneous | `Radio()`             | `RadioWithParams(side string)`        | Generate a random radio call sign. |
| ✓    | Miscellaneous | `Tv()`                | `TvWithParams(side string)`           | Generate a TV station call sign. This is an alias for Radio() since they both follow the same rules. |
