
COMMIT_SHORT_SHA := $(shell git rev-parse --short HEAD)
BUILD=go build -o dumber -ldflags "-w -s -buildid= -X 'main.commitShortHash=$(COMMIT_SHORT_SHA)'" -trimpath
INSTALL_BIN_PATH=/usr/local/bin/dumber
INSTALL=install -Dm755 dumber ${INSTALL_BIN_PATH}

help:
	@echo "Usage:"
	@echo
	@echo "    make [build|install|clean|distclean|test]"
	@echo

build:
	${BUILD}

clean:
	rm -f dumber test/*sections*

distclean: clean
	rm -f ${INSTALL_BIN_PATH}
	
install: build
	${INSTALL}

.PHONY: test
test: clean build
	test/run
	@echo

watch:
	@which CompileDaemon > /dev/null 2>&1 || (echo "CompileDaemon is required to watch (https://github.com/githubnemo/CompileDaemon)."; exit 1)
	CompileDaemon -build "${BUILD}" -command "${INSTALL}"

