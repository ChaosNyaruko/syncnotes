run:
	go run main.go -p -o "${HOME}/GitPrjs/dotfiles"
message:
	go run main.go  -o "${HOME}/GitPrjs/dotfiles" -m "test-${RAND}"
