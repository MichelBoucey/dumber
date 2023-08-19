
BUILD=go build -o dumber
MV=mv dumber ${HOME}/go/bin/dumber

help:
	@echo "Usage:"
	@echo
	@echo "    make [build|clean|distclean|test|watch]"
	@echo

build:
	@echo
	${BUILD}
	${MV}
	@echo && dumber -v || echo

clean:
	rm -f test/*numbered-sections*

distclean: clean
	rm -f ${HOME}/go/bin/dumber
	
watch:
	@which CompileDaemon > /dev/null 2>&1 || (echo "CompileDaemon is required to watch (https://github.com/githubnemo/CompileDaemon)."; exit 1)
	CompileDaemon -build "${BUILD}" -command "${MV}"

.PHONY: test
test: build
	test/run
	@echo

