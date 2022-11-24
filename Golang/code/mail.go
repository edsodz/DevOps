package main

import "fmt"

func main() {
	/* This is my first program in Golang */
	fmt.Println("Hello World!")
}

/*
The first line of the program package main defines the package name in which this program should lie. It is a mandatory statement, as Go programs run in packages. The main package is the starting point to run the program. Each package has a path and name associated with it.

The next line import "fmt" is a preprocessor command which tells the Go compiler to include the files lying in the package fmt.

The next line func main() is the main function where the program execution begins.

The next line comment is ignored by the compiler and it is there to add comments in the program. Comments are also represented using // similar to Java or C++ comments.

The next line fmt.Println(...) is another function available in Go which causes the message "Hello, World!" to be displayed on the screen. Here fmt package has exported Println method which is used to display the message on the screen.
*/
