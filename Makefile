help:
	@echo "Usage:"
	@echo
	@echo "    edit"
	@echo "    build"
	@echo "    build-release"
	@echo "    test"
	@echo "    fmt"
	@echo "    install"
	@echo "    watch"
	@echo "    lint"
	@echo "    clean"
	@echo

edit:
	vim src/main.rs

build:
	cargo build

build-release:
	cargo build --release

.PHONY: test
test: build
	test/run

fmt:
	rustfmt src/main.rs

install:
	cargo install --path .

watch:
	bacon

lint:
	cargo clippy

clean:
	@cargo clean
	@rm -f dumber test/*sections*
