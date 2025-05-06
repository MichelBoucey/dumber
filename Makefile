
COMMIT_SHORT_SHA := $(shell git rev-parse --short HEAD)
BUILD=go build -o dumber -ldflags "-X 'main.commitShortHash=$(COMMIT_SHORT_SHA)'"
INSTALL_BIN_PATH=/usr/local/bin/dumber
INSTALL=install -Dm755 dumber ${INSTALL_BIN_PATH}

help:
	@echo "Usage:"
	@echo
	@echo "    make [build|install|clean|distclean|test]"
	@echo

build:
	@echo
	${BUILD}

clean:
	rm -f dumber test/*sections*

distclean: clean
	rm -f ${INSTALL_BIN_PATH}
	
install: build
	${INSTALL}
	@echo && dumber -v || echo

.PHONY: test
test: clean install
	test/run
	@echo

watch:
	@which CompileDaemon > /dev/null 2>&1 || (echo "CompileDaemon is required to watch (https://github.com/githubnemo/CompileDaemon)."; exit 1)
	CompileDaemon -build "${BUILD}" -command "${INSTALL}"

