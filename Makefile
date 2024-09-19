
COMMIT_SHORT_SHA := $(shell git rev-parse --short HEAD)
BUILD=go build -o dumber -ldflags "-X 'main.commitShortHash=$(COMMIT_SHORT_SHA)'"
MV=sudo mv dumber ~/go/bin/dumber

help:
	@echo "Usage:"
	@echo
	@echo "    make [build|clean|distclean|test|watch]"
	@echo

build:
	@echo
	${BUILD}

clean:
	rm -f test/*sections*

distclean: clean
	rm -f dumber ~/go/bin/dumber
	
install: build
	${MV}
	@echo && dumber -v || echo

watch:
	@which CompileDaemon > /dev/null 2>&1 || (echo "CompileDaemon is required to watch (https://github.com/githubnemo/CompileDaemon)."; exit 1)
	CompileDaemon -build "${BUILD}" -command "${MV}"

.PHONY: test
test: clean install
	test/run
	@echo

