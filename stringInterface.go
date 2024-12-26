package StringerInterface

import (
	"fmt"
	"strconv"
)

type Book struct {
	title string
}

func (b Book) String() string {
	return fmt.Sprint("The title of the Book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("Tis is te Number ", strconv.Itoa(int(c)))
}
