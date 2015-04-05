TEST=.
default: build

# http://cloc.sourceforge.net/
cloc:
	@cloc --not-match-f='Makefile|_test.go' .

fmt:
	@go fmt ./...

build:
	@mkdir -p bin
	@go build -o bin/codebook codebook/main.go


doc:
	@godoc -http=:6060

test:
	@go test $(TEST)

clean:
	rm -rf tmp

.PHONY: cloc fmt build doc test clean
