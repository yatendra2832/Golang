package main

import "fmt"

func main() {
	// that willwork in the lastInFirstOut
	defer fmt.Println("Hello World !")     //4
	defer fmt.Println("Hello World One !") //3
	defer fmt.Println("Hello World Two!")  //2
	fmt.Println("Defer in Golang")
	myDefer() //1
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

/* Defer in golang
In Go, defer is a keyword used to schedule a function call to be executed at the end of the surrounding function's execution, just before it returns. It's often used to ensure that certain cleanup or finalization tasks are performed regardless of how the function exits, whether normally or due to a panic.

Syntax:

defer functionCall(arguments)
functionCall: The function or method call that you want to defer.
arguments: The arguments to pass to the deferred function.
Example:

package main

import "fmt"

func main() {
    defer fmt.Println("World")

    fmt.Println("Hello")
}
Output:
Hello
World
In this example, fmt.Println("World") is deferred until the end of the main() function. So, "World" is printed after "Hello", even though it appears before it in the code.

Use Cases for defer:
Resource Cleanup: Defer is commonly used to ensure that resources such as files or network connections are properly closed when they are no longer needed.
Unlocking Mutexes: If you acquire locks or mutexes in a function, you can defer their release to ensure they are always unlocked, even in case of errors.
Logging and Tracing: Defer can be used to log entry and exit points of functions, helping with debugging and tracing program execution.
Panic Recovery: In combination with recover(), defer can be used to recover from panics and gracefully handle errors.
Execution Timing: Defer can also be used to delay the execution of a function until a specific point in time, such as after a loop or conditional block.
Order of Execution:
Deferred function calls are executed in Last-In-First-Out (LIFO) order, meaning the most recently deferred function is executed first, followed by the next most recent, and so on.
Notes:
Arguments to deferred functions are evaluated immediately, but the function call is deferred until the surrounding function returns.
Deferred functions are executed even if the surrounding function panics.
Be cautious when deferring functions with loops or large iterations, as they may lead to increased memory consumption.
In summary, defer in Go provides a convenient way to ensure that cleanup or finalization tasks are performed reliably, regardless of how a function exits. It's a powerful tool for improving code clarity and reliability.*/
