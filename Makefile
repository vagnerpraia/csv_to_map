export GO111MODULE = on
export GOOS = linux
export GOARCH = amd64

run-load-publisher:
	go run cmd/csv_to_map/main.go -file "bisac.csv" -separator ";" -quote "\""
