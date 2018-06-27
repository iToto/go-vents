.PHONY: run build-run

build:
	go build -o bin/go-vents

clean:
	rm -r bin

run:
	heroku local

build-run: build run