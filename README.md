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

Cannot add integer and float (type mismatch error).
