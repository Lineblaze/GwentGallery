.PHONY:run
run:
	go run app/cmd/main.go

.PHONY:build-w
build:
	go build -o build/gallery.exe cmd/main.go