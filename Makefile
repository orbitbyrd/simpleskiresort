build:
	go build -o bin/simpleskiresort main.go

run: build
	go run main.go

all: build