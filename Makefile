export GO111MODULE = on
export GOOS = linux
export GOARCH = amd64

run-load-publisher:
	go run cmd/transformer/main.go -file "codigos_bisac.csv" -separator ";" -quote "\""
