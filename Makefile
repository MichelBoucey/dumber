h:
	@echo "Makefile of dumber"

c:
	@cargo clean
	@rm -f dumber test/*sections*

b:
	cargo build

br:
	cargo build --release

e:
	vim src/main.rs

.PHONY: test
t: c b
	test/run

f:
	rustfmt src/main.rs

i: br
	cp target/release/dumber /usr/local/bin/

w:
	bacon

l:
	cargo clippy
