PORT?=8000
PACKAGE:=github.com/philip-bui/bit-ranges
COVERAGE:=coverage.txt

godoc:
	echo "localhost:${PORT}/pkg/${PACKAGE}"
	godoc -http=:${PORT}

test:
	go test -race -coverprofile=${COVERAGE} -covermode=atomic

coverage: test
	go tool cover -html=${COVERAGE}

