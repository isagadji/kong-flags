.PHONY: test
test:
	env CGO_ENABLED=0 go test -cover ./...
