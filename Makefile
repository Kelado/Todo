all: build run

build:
	./bin/templ generate
	go build -o bin/app

run:
	./bin/app

css-watch:
	npx tailwindcss -i ./public/input.css -o ./public/output.css -w