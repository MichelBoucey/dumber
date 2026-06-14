
COMMIT_SHORT_SHA := $(shell git rev-parse --short HEAD)
BUILD=go build -o dumber -ldflags "-w -s -buildid= -X 'main.commitShortHash=$(COMMIT_SHORT_SHA)'" -trimpath
INSTALL_BIN_PATH=~/.cargo/bin/
INSTALL=install -Dm755 dumber ${INSTALL_BIN_PATH}

help:
	@echo "Usage:"
	@echo
	@echo "    make [build|install|clean|distclean|test]"
	@echo

c:
	rm -f dumber test/*sections*

b:
	cargo build

e:
	vim src/main.rs

r: b
	./target/debug/dumber

.PHONY: test
t: c b
	test/run
	@echo

f:
	rustfmt src/main.rs

i:
	cargo install --path $(INSTALL_BIN_PATH)

w:
	bacon
