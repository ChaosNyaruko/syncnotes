package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

// TODO: define my own var to implement list values (e.g. -o dotfile dollop )
// check https://pkg.go.dev/flag#Value.
var obj = flag.String("o", "./tmp", "where you want to launch the sync process")
var push = flag.Bool("p", false, "whether automatically push commit to remote")

func main() {
	flag.Parse()
	log.Println(*push)
	// var objs  = []string{}
	var dotfilesPath string
	if home, err := os.UserHomeDir(); err != nil {
		log.Fatal(err)
	} else {
		dotfilesPath = home + "/GitPrjs/dotfiles"
		log.Println(home, dotfilesPath)
	}
	pull := exec.Command("git", "pull", "origin", "main")
	pull.Dir = dotfilesPath
	pull.Stdout = os.Stdout
	pull.Stderr = os.Stderr
	log.Println("git pulling..")
	if err := pull.Run(); err != nil {
		log.Fatalf("git pull err: '%v'", err)
	}
	log.Println("committing..")
	commit := exec.Command("git", "commit", "-a")
	// TODO: specify the commit msg without launch editor
	log.Println(commit)
	commit.Dir = dotfilesPath
	commit.Stdout = os.Stdout
	commit.Stderr = os.Stderr
	if err := commit.Run(); err != nil {
		log.Fatal(err)
	}
	if *push {
		log.Println("pushing..")
		push := exec.Command("git", "push")
		log.Println(push)
		push.Dir = dotfilesPath
		push.Stdout = os.Stdout
		push.Stderr = os.Stderr
		if err := push.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
