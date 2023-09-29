package main

import (
	"fmt"
)

func main() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	fmt.Println(0)
}

/*
Defer -> After the return function executes of any non-deferred func and
		 just before the callback to the above func, the runtime will execute
		 any deferred func that are in defer LIFO (last in first out order) queue,
		 that is defer funcs run in LIFO order, last defer func will callback to
		 above func in this case main.
Tasks
● “defer” multiple functions in main
	○ show that a deferred func runs after the func containing it exits.
	○ determine the order in which the multiple defer funcs run
*/
