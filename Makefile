default: test

build:
	go build -i -v

test:
	go test -v

clean:
	go clean

tags:
	gotags -f tags -R .

.PHONY: default build test clean tags
