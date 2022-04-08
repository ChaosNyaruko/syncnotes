package main

type myList []string

func (v *myList) String() string {
	return "this is mylist"
}

func (v *myList) Set(value string) error {
	*v = append(*v, value)
	return nil
}
