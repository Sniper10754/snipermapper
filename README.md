# snipermapper
Golang port of snipermapper


[![asciicast](https://asciinema.org/a/lnvEYEUsoCWiSu1HgjkaYpH7T.svg)](https://asciinema.org/a/lnvEYEUsoCWiSu1HgjkaYpH7T)


## Install
```sh
go get github.com/Sniper10754/snipermapper
go install github.com/Sniper10754/snipermapper
```

Ready to go!

## Usage examples

main.go

```go
package main

import (
  "fmt"
  "strconv"
  snipermapper "github.com/Sniper10754/snipermapper/api"
)

func main() {
  isportopen := snipermapper.ScanPort("tcp", "amazon.com", 80, 1)
      
  fmt.Println("Is port open? " + strconv.FormatBool(isportopen))
}
```
