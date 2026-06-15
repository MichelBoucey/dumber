COMMIT_SHORT_SHA := $(shell git rev-parse --short HEAD)
INSTALL_BIN_PATH=~/.cargo/bin/
INSTALL=install -Dm755 dumber ${INSTALL_BIN_PATH}

h:
	@echo "Makefile of dumber"

c:
	@rm -f dumber test/*sections*

b:
	cargo build

e:
	vim src/main.rs

r: b
	./target/debug/dumber test/test.md

.PHONY: test
t: c b
	@echo
	test/run
	@echo

f:
	rustfmt src/main.rs

i:
	cargo install --path $(INSTALL_BIN_PATH)

w:
	bacon

l:
	cargo clippy
