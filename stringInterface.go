package github.com/Kishore545/StringerInterface

import (
	"fmt"
	"strconv"
)

type Book struct {
	title string
}

func Books(b Book) String() string {
	return fmt.Sprint("The title of the Book is ", b.title)
}

type count int

func Counts(c count) String() string {
	return fmt.Sprint("Tis is te Number ", strconv.Itoa(int(c)))
}


	

