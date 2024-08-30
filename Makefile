generate:
	@go generate ./...

dev:
	@templ generate --watch --proxy="http://localhost:5000" & wgo run . backoffice


build: generate
	go build -o bin/app .

