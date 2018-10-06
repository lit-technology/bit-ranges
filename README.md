# Bit Ranges

[![BuildStatus Widget]][BuildStatus Result]
[![CodeCov Widget]][CodeCov Result]
[![GoReport Widget]][GoReport Status]
[![GoDoc Widget]][GoDoc]

[BuildStatus Result]: https://travis-ci.org/philip-bui/bit-ranges
[BuildStatus Widget]: https://travis-ci.org/philip-bui/bit-ranges.svg?branch=master

[CodeCov Result]: https://codecov.io/gh/philip-bui/bit-ranges
[CodeCov Widget]: https://codecov.io/gh/philip-bui/bit-ranges/branch/master/graph/badge.svg

[GoReport Status]: https://goreportcard.com/report/github.com/philip-bui/bit-ranges
[GoReport Widget]: https://goreportcard.com/badge/github.com/philip-bui/bit-ranges

[GoDoc]: https://godoc.org/github.com/philip-bui/bit-ranges
[GoDoc Widget]: https://godoc.org/github.com/philip-bui/bit-ranges?status.svg

Library for setting and unsetting `n` least significant bits as ranges, and merging intersecting ranges.

## Usage

```go
import (
	"github.com/philip-bui/bit-ranges"
)

func main() {
	l, h := bits.RangeUint64(5, 777) // 1100001001.
	// l = 768, 1100000000.
	// h = 799, 1100011111.
	arr := bits.RangesUint64(5, 700, 799, 777, 768, 800, 77)
	// arr = [831, 768, 703, 672, 95, 64]. Arr will always be even amount (pairs).
}
```

## License

Bit Ranges is available under the MIT license. [See LICENSE](https://github.com/philip-bui/bit-ranges/blob/master/LICENSE) for details.
