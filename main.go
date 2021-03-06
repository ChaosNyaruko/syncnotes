package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

var Version = "v0.0.3"

var p = flag.Bool("p", false, "whether automatically push commit to remote")
var m = flag.String("m", time.Now().Format("2006/01/02 15:04:05"), "message you want to add as git commit messages")
var a = flag.Bool("a", false, "whether to launch your editor when commit, analogous to git commit -a, it OVERWRITES -m message")
var f = flag.Bool("f", false, "whether fetch/pull from remote first")
var v = flag.Bool("v", false, "show version")

var objs myList

func main() {
	flag.Var(&objs, "o", "specify where you want to execute the sync process")
	flag.Parse()
	// fmt.Printf(" push:'%v' objs:'%v' message:'%v'\n", *p, objs, *m)
	if *v {
		fmt.Printf("syncnotes version: %v\n", Version)
		return
	}
	path, err := os.Getwd()
	// path, err := os.Executable()
	if err != nil {
		log.Fatal(path)
	}
	if home, err := os.UserHomeDir(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("$HOME:", home)
	}

	if len(objs) == 0 {
		log.Fatal("sync paths must be given!")
	}

	for _, obj := range objs {
		fmt.Printf("current dir:'%v' push:'%v' obj:'%v' message:'%v' dasha: %t\n", path, *p, obj, *m, *a)
		if *f {
			if err := pull(obj); err != nil {
				log.Println(err)
				continue
			}
		}
		if err := commit(obj, *m, *a); err != nil {
			log.Println(err)
			continue
		}
		if *p {
			if err := push(obj); err != nil {
				log.Println(err)
				continue
			}
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

func commit(path string, message string, a bool) error {
	add := exec.Command("git", "add", ".")
	add.Dir = path
	add.Stdout = os.Stdout
	add.Stderr = os.Stderr
	if err := add.Run(); err != nil {
		return fmt.Errorf("git add err: %v '%v'", commit, err)
	}
	log.Println("committing..")
	commit := exec.Command("git", "commit", "-m", message)
	if a {
		commit = exec.Command("git", "commit", "-a")
	}
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
