.PHONY: noise

noise:
	CGO_ENABLED=0 go build -o ./bin/noise ./main.go
