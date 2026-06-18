DUMBER_CUR_VER := 4.0.0

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
	rm -rf ~/.tmp/dumber/
	mkdir ~/.tmp/dumber/
	cp LICENSE ~/.tmp/dumber/
	cp target/release/dumber ~/.tmp/dumber/
	cd ~/.tmp/ && tar -czvf dumber-$(DUMBER_CUR_VER)-linux-86_64.tar.gz dumber/

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
