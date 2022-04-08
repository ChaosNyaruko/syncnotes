run:
	go run main.go model.go -p -o "${HOME}/GitPrjs/dotfiles"
message:
	go run main.go model.go -o "${HOME}/GitPrjs/dotfiles" -m "test-${RAND}"
objs:
	go run model.go main.go  -o "${HOME}/GitPrjs/dotfiles" -o "${HOME}/GitPrjs/symmetrical-dollop"  -m "test-${RAND}"

