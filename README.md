# phpfuncs

[![PkgGoDev](https://pkg.go.dev/badge/serkanalgur/phpfuncs)](https://pkg.go.dev/serkanalgur/phpfuncs)

Php Functions in Go

:bangbang: **THIS MODULE IS UNDER DEVELOPMENT. PLEASE USE CAUTIOUSLY** :bangbang:

## Installation

```bash
  go get github.com/serkanalgur/phpfuncs

  # For update

  go get -u github.com/serkanalgur/phpfuncs
```

Description and other info will be here.


### Proper Documentation

Please Visit [https://pkg.go.dev/badge/serkanalgur/phpfuncs](https://pkg.go.dev/badge/serkanalgur/phpfuncs)

<!-- ## Functions

- [InArray (PHP in_array)](#inarray-php-in_array)
- [Time (PHP time)](#time-php-time)

### InArray (PHP in_array)

#### Usage

  ```go

    //needle str|int
    //haystack array []string, []int

    phpfuncs.InArray(needle,haystack)
    //Return will be true or false like in PHP
  ```

#### Example

You can see in [examples folder](examples/InArray.go)

```go
  package main
  import (
    "fmt"
    "strings"

    "github.com/serkanalgur/phpfuncs"
  )

  func main(){

    test := strings.Split("Denee için serkan algur kişisinden birşeyler seçmeliyiz."," ")

    arrayo := phpfuncs.InArray("serkaaaan",test)

    fmt.Printf("Result is : %v\n",arrayo)
  }
````

### Time (PHP time)

#### Usage of Time Function

  ```go
    phpfuncs.Time()
    //Return current Unix timestamp
  ```

#### Example of Time Function

You can see in [examples folder](examples/Time.go)

```go
  package main

  import (
    "fmt"

    "github.com/serkanalgur/phpfuncs"
  )

  func main(){
    currentTime := phpfuncs.Time()
    fmt.Printf("Current Timestamp is: %d",currentTime)
  }
``` -->
