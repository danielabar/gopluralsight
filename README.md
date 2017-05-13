<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Go Fundamentals](#go-fundamentals)
  - [Hello World](#hello-world)
    - [Setting up Workspace](#setting-up-workspace)
    - [First Program](#first-program)
    - [Functions Primer](#functions-primer)
    - [Printing to screen](#printing-to-screen)
    - [Running program](#running-program)
    - [More detail](#more-detail)
  - [Variables and Constants](#variables-and-constants)
    - [Declaring at the Package Level](#declaring-at-the-package-level)
    - [Determining Types](#determining-types)
    - [Short Assignment](#short-assignment)
    - [Pointers](#pointers)
    - [Passing by Value](#passing-by-value)
    - [Passing by Reference](#passing-by-reference)
    - [Constants](#constants)
    - [Accessing Environment Variables](#accessing-environment-variables)
  - [Functions](#functions)
    - [Go Function Syntax](#go-function-syntax)
    - [Function Basics](#function-basics)
    - [Variadic Functions](#variadic-functions)
  - [Conditionals](#conditionals)
    - [if Syntax](#if-syntax)
    - [Simple Initialization Statements](#simple-initialization-statements)
    - [Switch Syntax](#switch-syntax)
    - [Breaking and Fall-through](#breaking-and-fall-through)
    - [The Role of if in Error Handling](#the-role-of-if-in-error-handling)
  - [Loops](#loops)
    - [for Syntax](#for-syntax)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Go Fundamentals

> Learning Go with Pluralsight [course](https://app.pluralsight.com/library/courses/go-fundamentals/table-of-contents)

## Hello World

Go encourages modular programming via small packages being composed as applications.

### Setting up Workspace

Go needs a *workspace*, root folder structure for apps. Under that need 3 sub-directories for: src, pkg, bin.

### First Program

[Example](hello/hello-world.go)

Every go program needs a *package declaration*, for example:

```go
package main
```

`main` is special. Tells compiler to compile program as a stand-alone executable, and not as a shared library. i.e. every executable needs `package main` as its first line of code.

Then must have a main function:

```go
func main() {

}
```

`func main()` is executable program's entrypoint. First function that will be run in a go executable program.

### Functions Primer

Functions are first class citizens, can be passed as args, assign to vars, etc.

Start with `func` keyword, then can name it anything, except `main` which has special meaning.

All functions have parentheses `()`, which list any arguments that get passed to function, then list any returns.

Function code written inside pair of curly braces `{}`.

Main function doesn't take any arguments and doesn't return any value. When main exits, entire program exits. If there were no errors, it returns 0 exit code to OS.

### Printing to screen

Idiomatic way is to use `fmt.println` from the format package. Therefore need to import it:

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello from", runtime.GOOS)
}
```

`fmt` ships as part of the standard library.


### Running program

```shell
cd $GOPATH
cd /path/to/src
go run hello-world.go
```

Run command compiles source into executable, puts it in a temp dir and runs it.

### More detail

All go source available on github, for example, [fmt package](https://github.com/golang/go/blob/master/src/fmt/print.go).

Note that declared package name `package fmt` is the import name `import ( "fmt" )`.

Note that shared library `fmt` does not have package main or a main func.

Go is case sensitive! For the most part, all go code is written in lower case. But note function name with capital `fmt.Println`.

In order for functions in a package to be exposed, must start with capital.

Go is not whitespace sensitive but recommend spacing for legibility.

Comments are generally line comments `// this is a line comment`, even when span multiple lines, rather than block comments.

## Variables and Constants

### Declaring at the Package Level

[Example](vars/hello-vars.go)

Variables are *statically typed*. Eg, variable declared as string, its type set at compile time and can never change.

When declaring variable at package level, i.e. outside of any functions, must use `var keyword`. Go does not allow non-declaration statements at package level.

Variable names must start with underscore or letter. Name of variable goes on the left and type goes on the right.

```go
var (
	name   string  //Name of subscriber
	course string  //Name of current course
	module float64 //Current place in course
)
```

Variables of the same type can also go on the same line:

```go
var (
  name, course string
  module float64
)
```

Generally when variables are declared but not initialized in code, they will be initialized with a 0 value (0 for numeric, empty string for strings).

```go
func main() {
	fmt.Println("Name is set to", name)  //Name is set to
	fmt.Println("Course is set to", course) //Course is set to
	fmt.Println("Module is set to", module) //MOdule is seto to 0
}
```

### Determining Types

[Example](vars/hello-types.go)

Can check types with runtime reflection using `reflect` package.

```go
func main() {
	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name)) //Name is  and is of type string
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module)) //Module is 0 and is of type float64
}
```

Can also *infer* types:

[Example](vars/infer-types.go)

```go
var (
	inferredName, inferredCourse, inferredModule = "Nigel", "Docker Deep Dive", 3.2
)

func main() {
	fmt.Println("Name is", inferredName, "and is of type", reflect.TypeOf(inferredName)) //Name is Nigel and is of type string
	fmt.Println("Module is", inferredModule, "and is of type", reflect.TypeOf(inferredModule)) //Module is 3.2 and is of type float64
}
```

For legibility, better to initialize multiple variables each on its own line.

Cannot add integer and float (type mismatch error):

```go
a := 10.0000000000
b := 3
c := a + b //invalid operation: a + b (mismatched types float64 and int)
```

But can convert float to int:

```go
a := 10.0000000000
b := 3
c := int(a) + b
fmt.Println("\nC has value:", c, "and is of type:", reflect.TypeOf(c)) //C has value: 13 and is of type: int
```

Note that `a` is still a float, calling `int(a)` does not change `a` to int.

### Short Assignment

[Example](vars/short-asg.go)

Variables declared at package level are available to all functions in the package, i.e. they're *global* in scope.

Can also declare variables within a function, then use a shorthand declare-initialize construct:

```go
func main() {
  name := "Nigel"
}
```

Shorthand only works in functions, and when declaring and initializing variables on the same line. This is idiomatic Go.

Note that a variable is declared in the shorthand notation in a function and not used, the program will not compile. Will get error like "somevar declared and not used". Whereas when variable declared at package level and not used, program will compile.

### Pointers

Go passes arguments by value, not by reference. When passing an argument to a function, Go makes a *copy* of the value being passed, places copy on the stack for use by the function. Variable itself is not placed on the stack.

**Behind the scenes**

When variable is created, Go sets aside memory for it, for eg: at memory addres `0xAA`, place a value "Docker Deep Dive".
When this variable passed as argument to a function, Go makes a copy, eg: at memory address `0xBB`, value "Docker Deep Dive".
Copy at `0xBB` is placed on stack. Achieves immutability, any changes made to var by function will only affect the copy, not the original.

To workaround this default behaviour, use *pointers*. Ampersand before variable name represents memory address (aka pointer value) rather than value of the variable. Can also create a pointer variable using ampersand.

In this example, the `ptr` variable holds the pointer value of the `module` variable. Asterisk in front of pointer variable dereferences it so we get the contents at that memory address:

```go
func main() {
	module := 3.2
	ptr := &module
	fmt.Println("Ptr is", ptr, "and is of type", reflect.TypeOf(ptr)) //Ptr is 0xc42000a2c8 and is of type *float64
	fmt.Println("Memory address of *module* variable is", ptr, "and the value of *module* is", *ptr) //Memory address of *module* variable is 0xc42000a2c8 and the value of *module* is 3.2
}
```

Summary:
* `&` references a pointer
* `*` de-references a pointer

### Passing by Value

[Example](vars/pass-by-value.go)

To demonstrate, a function to change the value of a variable:

```go
// changeCourse gets a COPY of the course variable
// course argument is a string and this function also returns a string
func changeCourse(course string) string {
  // Use = rather than := because not declaring a new variable,
  // just assigning a new value to existing variable.
	course = "First Look: Native Docker Clustering"

	fmt.Println("Trying to change your cousrse to", course)
	return course
}
```

Return values can be named or unnamed. Above is example of unnamed.

Now use the function. Note that any changes made to variable inside of function have no impact on original variable passed in.

```go
func main() {
	name := "Nigel"
	course := "Docker Deep Dive"

	fmt.Println("\nHi", name, "you're currently watching", course) // Docker Deep Dive
	changeCourse(course)
	fmt.Println("\nHi", name, "you're currently watching", course) // Still Docker Deep Dive
}
```

### Passing by Reference

[Example](vars/pass-by-ref.go)

Use pointers. `&` on variable passed in to function, and `*` on variable used within function.

```go
package main

import "fmt"

func main() {
	name := "Nigel"
	course := "Docker Deep Dive"

	fmt.Println("\nHi", name, "you're currently watching", course) // Docker Deep Dive

	// pass pointer to course variables location in memory
	changeCourse(&course)

	fmt.Println("\nHi", name, "you're currently watching", course) // First Look: Native Docker Clustering
}

// askterisk tells Go that `course` is a POINTER to a string variable
func changeCourse(course *string) string {
	// asterisk tells Go that we're assigning to location in memory that the course pointer is referencing
	*course = "First Look: Native Docker Clustering"

	fmt.Println("Trying to change your cousrse to", *course)
	return *course
}
```

### Constants

Constants are immutable, once assigned a value, can never change. Declared with `const` keyword, but cannot use shorthand `:=` notation.

```go
const speedOfLightMph = 186000
```

### Accessing Environment Variables

[Example](vars/env-vars.go)

Use `os` package for access to various operating system stuff, including environment variables.

Print out all environment variables with their values, one per line:

```go
for _, env := range os.Environ() {
  fmt.Println(env)
}
```

Get an environment variable by key and use it:

```go
name := os.Getenv("USER")
fmt.Println("Hello", name)
```

## Functions

### Go Function Syntax

Start with `func` keyword, followed by name of function, then parentheses. Any code executed inside of function goes in pair of curly braces.

Function parameters are specified inside the parentheses. Includes parameter name and its type. `func` line is called *function signature*. Multiple parameters separated by commas.

Return type also defined in function signature, after the parens but before the curly braces. Multiple return values must be enclosed in parens.

Use `return` keyword to return a value back to caller.

```go
func titleCase(text string) string {
  <code>
  return convertedText
}
```

`main` function is special
* gets called automatically
* does not take any input parameters
* does not return any values

### Function Basics

[Example](functions/func-basics.go)

### Variadic Functions

[Example](functions/variadic.go)

Used when don't know how many values will be passed to a function. Use ellipsis `...` before the type:

```go
func bestLeageFinishes(finishes ...int) {

}
```

Values passed into the function get made into a slice of ints:

```go
bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 16)
```

Slice is a list, in above example with 6 entries:

0: 13
1: 10
2: 13
3: 17
4: 14
5: 16

Given a slice, can loop over it using `range`:

```go
for _, f := range finishes {
  // do something with value f
}
```

## Conditionals

### if Syntax

[Example](conditionals/if-syntax.go)

Works exactly the same as other languages. Evaluate a condition, if it's true branch one way, otherwise the other way.

However, in Go, condition to evaluate **must evaluate to a Boolean expression**. Cannot use integers or strings. Strict approach keeps things clean.

**Boolean Comparison Operators**

`==` Equal to

`!=` Not equal to

`<` Less than

`<=` Less than or equal to

`>` Greater than

`>=` Greater than or qual to

`&&` AND

`||` OR

Anatomy of `if`. Note curly brace *MUST be on same line* as `if` expression. Because of how compiler inserts semicolons at end of each line. You don't type those in yourself, compiler inserts them.

Can also use `else if` and `else`:

```go
if <Boolean expression> {
  <code block>
} else if <Boolean expression> {
  <code block>
} else {
  <code block>
}
```

If first `if` condition evaluates to true, the code jumps out of the entire `if` block, the subsequent `else if` and `else` statements will not be evaluated.

Can have multiple `else if`'s but only a single `else` and must be the last statement.

`if` statements can also be nested.

### Simple Initialization Statements

[Example](conditionals/if-syntax.go)

These are optional and execute before the Boolean expression is evaluated.

```go
if <simple statement> ; <Boolean expression> {
  <code block>
}
```

Idiomatic go is to use this to initialize variables that will be used in the if block.

Variables declared in the initialization statement are *scoped* to the `if` statement. When code finishes the `if` statement, these vars are out of scope and will be garbage collected.

### Switch Syntax

Both `simple statement` and `expression` are optional. But if do use `simple statement`, then following `;` is *mandatory*, even if no `expression` follows.

Variables declared in `simple statement` are only scoped to the `switch` block.

```go
switch <simple statement> ; <expression> {
case <expression>: <code>
case <expression>: <code>
...
default: <code>
}
```

`default` block executes if no matches in any of the `case` statements. Lexically it doesn't have to be the last option but more readable to have it at the bottom.

Expression could be a string, eg:

```go
switch "Docker Deep Dive" {
case "Docker Deep Dive": <code block will execute>
case "Go Fundamentals": <code block will NOT execute>
default: <code block also will not execute>
}
```

`switch` type and `case` type must be the same in order for them to be comparable.

After match is found, `switch` block is exited and code continues executing after the curly brace closing the `switch` block. Unlike other languages that have *implicit fallthrough*, after match is found, all cases below it also get executed. This does not happen with Go.

### Breaking and Fall-through

[Example](conditionals/fallthrough.go)

Each `case` statement has an implicit `break`. But if want fallthrough behaviour, just add `fallthrough` keyword as last line in case statement.

The following will output "Here are some recommended Docker Courses" AND "Here are some recommended Windows courses":

```go
topic := "docker"

switch topic {
case "linux":
  fmt.Println("Here are some recommended Linux courses...")
case "docker":
  fmt.Println("Here are some recommended Docker courses...")
  fallthrough
case "windows":
  fmt.Println("Here are some recommended Windows courses...")
default:
  fmt.Println("Sorry we couldn't find a match, " +
    " why not try out Top 100 list!")
}
```

`fallthrough` only applies on a case by case basis. If want them all to fallthrough, must specify the `falthrough` keyword on all of them.

However, idiomatic GO is not to use `fallthrough`, but to make multiple matches in the same `case` statement using a comma separted list:

```go
switch tmpNum := random(); tmpNum {
case 0, 2, 4, 6, 8:
  fmt.Println("We got an even number -", tmpNum)
case 1, 3, 5, 7, 9:
  fmt.Println("We got an odd number -", tmpNum)
}
```

[Full example](conditionals/not-fallthrough.go).

### The Role of if in Error Handling

[Example](conditionals/error-handling.go)

Idiomatic to return an `error` as the last return from functions and methods. For example a function with two return values that tests connectivity to a host. Note that `error` is a standard type defined in Go:

```go
func testConn(target string) (respTime float 64 err error) {
  ...
}
```

If all went well in the function, it should return `nil` for the value of `err`. Otherwise, it should be not `nil`. i.e. `nil` is used to indicate success.

Idiomatic Go is to *always* check the value of returned errors:

```go
if err != nil {
  <error handling code>
}
<code...>
```

## Loops

### for Syntax

[Example]

`expression` can be a Boolean (eg: `i < 10`) or a range. Can also be blank (which Go assumes to be `true`), which generates an infinite loop.

```go
for <expression> {
  <code>
}
```

Example using range. Range takes a list (either slice or a map) and iterates over the list. In every loop iteration, the current value of `courseList` is assigned to the variable `i`:

```go
for i := range courseList {
  <code>
}
```

For loop can be given simple pre and post expressions. eg:

```go
for i := 0; i < 10; i++ {
  <code>
}
```

pre: declare and initialize `i` to 0
boolean: test for `i` less than 10
post: increment `i` by one

Variables declared in pre statement only available within scope of for loop.

Pre statement runs only once, before first execution of loop, before evaluating expression.

Post statement runs at end of every iteration through the loop.
