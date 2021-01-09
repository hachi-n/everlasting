get-depends:
	go get github.com/markbates/pkger/cmd/pkger

clean:
	rm -f ./bin/everlasting

go-build:
	pkger -o pack --include /resources
	go build  -o ./bin/everlasting ./cmd/everlasting

build:
	$(MAKE) clean
	$(MAKE) go-build
run: 
	$(MAKE) clean
	$(MAKE) go-build
	./bin/everlasting tui
