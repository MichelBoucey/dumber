
help:
	@echo "Usage:"
	@echo ""
	@echo "    make [build|watch]"
	@echo ""

build:
	go build -o dumber cmd/dumber/main.go

watch:
	@which CompileDaemon > /dev/null 2>&1 || (echo "CompileDaemon is required to watch (https://github.com/githubnemo/CompileDaemon)."; exit 1)
	CompileDaemon -build "go build -o dumber cmd/dumber/main.go" -command "mv dumber ${HOME}/go/bin/dumber"

