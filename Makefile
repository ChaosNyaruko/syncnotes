installAs = "/usr/local/bin/syncnotes"
build:
	go build -o output/syncnotes main.go model.go
install:
	cp output/syncnotes ${installAs}
uninstall:
	rm ${installAs}
clean:
	rm output/*
run:
	go run main.go model.go -p -o "${HOME}/GitPrjs/dotfiles" -a
message:
	go run main.go model.go -o "${HOME}/GitPrjs/dotfiles" -m "test-${RAND}"
objs:
	go run model.go main.go  -o "${HOME}/GitPrjs/dotfiles" -o "${HOME}/GitPrjs/symmetrical-dollop"  -m "test-${RAND}"

