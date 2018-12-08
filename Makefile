sources:=$(shell find . -name '*.go' -and -not -name '*_test.go')
binary:=hound

all: $(binary)

$(binary): $(sources)
	go build

test:
	go test ./...

clean:
	rm -rf $(binary)
