DUMBER_CUR_VER := $(shell cargo pkgid | cut -d "#" -f2)

help:
	@echo "Usage:"
	@echo
	@echo "    edit"
	@echo "    build"
	@echo "    build-release"
	@echo "    dist"
	@echo "    test"
	@echo "    fmt"
	@echo "    install"
	@echo "    watch"
	@echo "    lint"
	@echo "    clean"
	@echo

edit:
	vim src/main.rs

dist: build-release
	rm -rf ~/.tmp/dumber-*
	mkdir ~/.tmp/dumber-$(DUMBER_CUR_VER)
	cp LICENSE ~/.tmp/dumber-$(DUMBER_CUR_VER)/
	cp target/release/dumber ~/.tmp/dumber-$(DUMBER_CUR_VER)/
	cd ~/.tmp/ && tar -czvf dumber-$(DUMBER_CUR_VER)-linux-86_64.tar.gz dumber-$(DUMBER_CUR_VER)
	rm -rf ~/.tmp/dumber-$(DUMBER_CUR_VER)/

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
