package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

// TODO: define my own var to implement list values (e.g. -o dotfile dollop )
// check https://pkg.go.dev/flag#Value.
var obj = flag.String("o", "", "where you want to launch the sync process")
var p = flag.Bool("p", false, "whether automatically push commit to remote")
var m = flag.String("m", time.Now().Format("2006/01/02 15:04:05"), "message you want to add as git commit messages")

var list myList

func main() {
	flag.Var(&list, "list", "shit")
	flag.Parse()
	fmt.Println(list)
	return
	path, err := os.Getwd()
	// path, err := os.Executable()
	if err != nil {
		log.Fatal(path)
	}
	// var objs  = []string{}
	if home, err := os.UserHomeDir(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("$HOME:", home)
	}

	if *obj == "" {
		log.Fatal("a sync path must be given!")
	}

	fmt.Printf("current dir:'%v' push:'%v' obj:'%v' message:'%v'\n", path, *p, *obj, *m)
	if err := pull(*obj); err != nil {
		log.Println(err)
	}
	if err := commit(*obj, *m); err != nil {
		log.Println(err)
	}
	if *p {
		if err := push(*obj); err != nil {
			log.Println(err)
		}
	}
}

func pull(path string) error {
	log.Println("git pulling..")
	pull := exec.Command("git", "pull", "origin", "main")
	pull.Dir = path
	pull.Stdout = os.Stdout
	pull.Stderr = os.Stderr
	if err := pull.Run(); err != nil {
		return fmt.Errorf("git pull err: %v '%v'", pull, err)
	}
	return nil
}

func commit(path string, message string) error {
	add := exec.Command("git", "add", ".")
	add.Dir = path
	add.Stdout = os.Stdout
	add.Stderr = os.Stderr
	if err := add.Run(); err != nil {
		return fmt.Errorf("git add err: %v '%v'", commit, err)
	}
	log.Println("committing..")
	commit := exec.Command("git", "commit", "-m", message)
	// TODO: specify the commit msg without launch editor
	commit.Dir = path
	commit.Stdout = os.Stdout
	commit.Stderr = os.Stderr
	if err := commit.Run(); err != nil {
		return fmt.Errorf("git commit err: %v '%v'", commit, err)
	}
	return nil
}

func push(path string) error {
	log.Println("pushing..")
	push := exec.Command("git", "push")
	push.Dir = path
	push.Stdout = os.Stdout
	push.Stderr = os.Stderr
	if err := push.Run(); err != nil {
		return fmt.Errorf("git push err: %v '%v'", push, err)
	}
	return nil
}
