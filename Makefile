
BUILD_COMMAND=go build -o dumber cmd/dumber/main.go

help:
	@echo "Usage:"
	@echo ""
	@echo "    make [build|test|watch]"
	@echo ""

build:
	${BUILD_COMMAND}
	mv dumber ${HOME}/go/bin/dumber

watch:
	@which CompileDaemon > /dev/null 2>&1 || (echo "CompileDaemon is required to watch (https://github.com/githubnemo/CompileDaemon)."; exit 1)
	CompileDaemon -build "${BUILD_COMMAND}" -command "mv dumber ${HOME}/go/bin/dumber"

.PHONY: test
test: build
	test/run

