# [github.com/bitcrook/cycull/pkg/noauth/userlookup](https://github.com/bitcrook/cycull/tree/main/pkg/noauth/userlookup) - no authentication required


## Types

``` go
type Website struct {
	Title  string
	Domain string
	Delete string
	Valid  bool
	Extra  string
}
```

## Functions


UserLookup takes a username (string) and a write flag (soon to be deprecated) as its parameters and returns a slice of type Website ([]Website)
``` go
func UserLookup(userres string, write bool) []Website {
	var wg sync.WaitGroup
	var arrNo []Website = noRedirSites(userres)
	for _, v := range arrNo {
		go checkUser(v, userres, write, false, &wg)
	}
	var arrYes []Website = redirSites(userres)
	for _, v := range arrYes {
		go checkUser(v, userres, write, true, &wg)
	}

	wg.Wait()
	return results
}
```

## Usage

Reading from the returned struct can be done like so:
``` go
package main

import (
	"fmt"
	"strconv"

	"github.com/bitcrook/cycull/pkg/noauth/userlookup"
)

func main() {
	x := userlookup.UserLookup("maraudery", false)
	for _, v := range x {
		fmt.Println(v.Domain + " - " + strconv.FormatBool(v.Valid))
	}
}
```
Part of output from the above example:
```
https://en.gravatar.com/maraudery - true
https://medium.com/@maraudery - false
https://pypi.org/user/maraudery - false
```