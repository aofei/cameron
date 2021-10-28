# Cameron

[![GitHub Actions](https://github.com/aofei/cameron/workflows/Main/badge.svg)](https://github.com/aofei/cameron)
[![codecov](https://codecov.io/gh/aofei/cameron/branch/master/graph/badge.svg)](https://codecov.io/gh/aofei/cameron)
[![Go Report Card](https://goreportcard.com/badge/github.com/aofei/cameron)](https://goreportcard.com/report/github.com/aofei/cameron)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/aofei/cameron)](https://pkg.go.dev/github.com/aofei/cameron)

An avatar generator for Go.

Oh, by the way, the name of this project came from the
[Avatar](https://en.wikipedia.org/wiki/Avatar_(2009_film))'s director
[James Cameron](https://en.wikipedia.org/wiki/James_Cameron).

## Features

* [Identicon](https://en.wikipedia.org/wiki/Identicon)

## Installation

Open your terminal and execute

```bash
$ go get github.com/aofei/cameron
```

done.

> The only requirement is the [Go](https://golang.org), at least v1.13.

## Quick Start

Create a file named `cameron.go`

```go
package main

import (
	"bytes"
	"image/png"
	"net/http"

	"github.com/aofei/cameron"
)

func main() {
	http.ListenAndServe("localhost:8080", http.HandlerFunc(handleIdenticon))
}

func handleIdenticon(rw http.ResponseWriter, req *http.Request) {
	buf := bytes.Buffer{}
	png.Encode(&buf, cameron.Identicon([]byte(req.RequestURI), 540, 60))
	rw.Header().Set("Content-Type", "image/jpeg")
	buf.WriteTo(rw)
}
```

and run it

```bash
$ go run cameron.go
```

then visit `http://localhost:8080` with different paths.

## Community

If you want to discuss Cameron, or ask questions about it, simply post questions
or ideas [here](https://github.com/aofei/cameron/issues).

## Contributing

If you want to help build Cameron, simply follow
[this](https://github.com/aofei/cameron/wiki/Contributing) to send pull requests
[here](https://github.com/aofei/cameron/pulls).

## TODOs

* [ ] Add support for cartoon avatar
* [ ] Add support for simulation avatar

## License

This project is licensed under the MIT License.

License can be found [here](LICENSE).
