package main

import "fmt"

type myList []string

func (v *myList) String() string {
	return fmt.Sprintf("%q", *v)
}

func (v *myList) Set(value string) error {
	*v = append(*v, value)
	return nil
}
