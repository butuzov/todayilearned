all:
	cargo test
	cargo bench

	cargo binstall cargo-fuzz 
	# will fail on arm 
	cargo fuzz run --no-default-features test_target
	# will fail on arm 