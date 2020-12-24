# go-expandstrict

A replacement for [`os.Expand`](https://golang.org/pkg/os/#Expand) that returns an error if any of the variables to be expanded are undefined.

## Install

`$ go get github.com/njkleiner/go-expandstrict`

## Usage

```go
package example

import (
    "os"
    "fmt"

    "github.com/njkleiner/go-expandstrict"
)

func Example() {
    exp, err := expandstrict.Expand("Hello, is anybody $HOME?", os.Getenv)

    if err != nil {
        // $HOME is undefined
        panic(err)
    }

    fmt.Printf("expanded: %s", exp)
}
```

## Contributing

You can contribute to this project by [sending patches](https://git-send-email.io) to `noah@njkleiner.com`. Pull Requests are also welcome.

## Authors

* [Noah Kleiner](https://github.com/njkleiner)

See also the list of [contributors](https://github.com/njkleiner/go-expandstrict/contributors) who participated in this project.

## License

This project is licensed under the MIT License. See the [LICENSE.md](LICENSE.md) file for details.
