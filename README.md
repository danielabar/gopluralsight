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

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Go Fundamentals

> Learning Go with Pluralsight [course](https://app.pluralsight.com/library/courses/go-fundamentals/table-of-contents)

## Hello World

Go encourages modular programming via small packages being composed as applications.

### Setting up Workspace

Go needs a *workspace*, root folder structure for apps. Under that need 3 sub-directories for: src, pkg, bin.

### First Program

[hello-world.go](hello/hello-world.go)

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

[hello-vars.go](vars/hello-vars.go)

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

[hello-types.go](vars/hello-types.go)

Can check types with runtime reflection using `reflect` package.

```go
func main() {
	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name)) //Name is  and is of type string
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module)) //Module is 0 and is of type float64
}
```

Can also *infer* types:

[infer-types.go](vars/infer-types.go)

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

[shorg-asg.go](vars/short-asg.go)

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
