.PHONY: build

build:
	@echo removing
	rm -rf bootstrap bootstrap.zip
	@echo building
	GOOS=linux GOARCH=amd64 go build -o bootstrap cmd/transactions/*.go
	zip bootstrap.zip bootstrap
	@echo Done
