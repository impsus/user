buildapi:
	go build -o out/api cmd/api/main.go

dev: buildapi
	./out/api