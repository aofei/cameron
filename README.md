# Cameron

[![Test](https://github.com/aofei/cameron/actions/workflows/test.yaml/badge.svg)](https://github.com/aofei/cameron/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/aofei/cameron/branch/master/graph/badge.svg)](https://codecov.io/gh/aofei/cameron)
[![Go Report Card](https://goreportcard.com/badge/github.com/aofei/cameron)](https://goreportcard.com/report/github.com/aofei/cameron)
[![Go Reference](https://pkg.go.dev/badge/github.com/aofei/cameron.svg)](https://pkg.go.dev/github.com/aofei/cameron)

An avatar generator for Go.

Fun fact, Cameron is named after [James Cameron](https://en.wikipedia.org/wiki/James_Cameron), the director of
[Avatar](https://en.wikipedia.org/wiki/Avatar_(2009_film)).

## Features

- Pixel-perfect [GitHub Identicons](https://github.blog/news-insights/company-news/identicons/)
- Zero third-party dependencies

## Installation

To use this project programmatically, `go get` it:

```bash
go get github.com/aofei/cameron
```

## Quickstart

Create a file named `cameron.go`:

```go
package main

import (
	"image/png"
	"net/http"

	"github.com/aofei/cameron"
)

func main() {
	http.ListenAndServe("localhost:8080", http.HandlerFunc(identicon))
}

func identicon(rw http.ResponseWriter, req *http.Request) {
	img := cameron.Identicon([]byte(req.RequestURI), 70)
	rw.Header().Set("Content-Type", "image/png")
	png.Encode(rw, img)
}
```

Then run it:

```bash
go run cameron.go
```

Finally, visit `http://localhost:8080` with different paths.

## Community

If you have any questions or ideas about this project, feel free to discuss them
[here](https://github.com/aofei/cameron/discussions).

## Contributing

If you would like to contribute to this project, please submit issues [here](https://github.com/aofei/cameron/issues) or
pull requests [here](https://github.com/aofei/cameron/pulls).

When submitting a pull request, please make sure its commit messages adhere to
[Conventional Commits 1.0.0](https://www.conventionalcommits.org/en/v1.0.0/).

## License

This project is licensed under the [MIT License](LICENSE).
