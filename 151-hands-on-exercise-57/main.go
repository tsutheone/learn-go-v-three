package main

import "fmt"

func main() {
	ixi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(sum(ixi))
}

func sum(xi []int) (total int) {
	for _, v := range xi {
		total += v
	}
	return total
}

/* Review of function concepts:
functions -> block of code that performs a task.
purpose of functions -> to perform a task while:
                            being modular -> clean rehusable code (no spaguetti).
							being abstract -> summarizes code.
							allowing DRY -> Don't repeat yourself (coding).
parameters vs arguments -> Parameters are the input(s) that are defined when
                           writing the function signature while arguments are
						   the actual values that are passed to a function
						   as inputs when the funcion is called.
syntax of functions -> func (receiver) identifier (parameter(s)) return(s) {
						<code>
						}
named returns -> to put an identifier to a return of a function, therefore the
                 variable that returns is declared in the code block created by
				 the function and the function can just return without redeclaring
				 that the return will be that variable. This hands on contains an
				 example of a named return in a function.
variadic parameters -> To be able to be of multiple parameters without
						specifying how many.
unfurling a slice -> remove the envelope of slice so that all elements that
                     where contained in the slice can be used as singular
					 entities through the code.
defer -> set up a function to be executed just before the callback of the
         surronding (above) function goes to the caller function. That's it just
		 after the surrounding function executes it's return.
methods -> functions are bounded to a type are called methods of that type.
polymorphism -> Object Oriented Programming concept that is the ability
                of a block of code to be used in more than one type. This is
				achieved through interfaces.
interfaces -> custom types that are used to specify a set of method(s) signatures
              that could be used in other types if the other types implement
			  a method signature with same identifier. This allow to method
			  abstraction and customization of that method.
			  Examples of interfaces are: Stringer, Logging, Writer. All provided
			  in the built-in golang pkgs.
anonymous func -> a function without identifier, it can can 0 to 1 parameter and
                  multiple return values.
first class citzien entities -> functional programming concept. In go it refers
								to entities (such as functions) that are allowed
								to have privileged behaviours. For instance
								passing such functions to as arguments of other
								functions. Another privileged behaviour is to be
								able to assign a function to a variable as a
								value of that variable.
callbacks -> Ability to pass a funciont into another function as an arguemnt.
closures {} -> limits the scope of variables (which could be functions hehe)
           creating inner scopes and outer scopes.
recursion -> Ability of a function that calls itself, an example of this is
             the factorial function.
life philosophy tip -> focus on what's important. not what's apparently urgent
                       Humans are hardwired to have urgencies some of which
					   are not enough justified in modern world be classed as
					   such. Therefore apparently unknown anxiety and stress
					   response appear. The tip here is to focus on importancy
					   rather of apparently urgency of things, staying calm and
					   reevaluatingg the apparentyl urgency and general situation.

*/
