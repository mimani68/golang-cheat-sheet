# Stack and Heap in Golang

In Golang, memory is allocated in two main places: the stack and the heap. Understanding the difference between these two memory regions is important for efficient memory management in Golang.

## Stack

The stack is a region of memory that is used for the execution of a thread. It is a scratch space where function calls and local variables are stored. The stack memory is managed automatically by the compiler and runtime, and it is allocated and deallocated as functions are called and return. The stack memory is fast to access and deallocate, making it ideal for managing temporary data associated with function calls 


## Heap

The heap is another region of memory used to store data in a program. Unlike the stack, the heap is not managed automatically by the compiler or runtime. Instead, the programmer is responsible for allocating and freeing memory on the heap as needed. The heap can grow dynamically at runtime and is often used to store data that is too large to fit on the stack or needs to persist beyond the lifetime of a function call. However, because the heap is not managed automatically, it is more prone to memory leaks and other types of errors if not used carefully 

## Stack vs Heap Allocation

In Golang, the compiler determines whether a variable should be allocated on the stack or the heap based on its lifetime and memory footprint. Variables with a known lifetime and memory footprint at compile time are allocated on the stack, while variables with an unknown or dynamic lifetime are allocated on the heap at runtime 

Stack allocation is generally cheaper and faster than heap allocation. Allocating memory on the stack only requires a few CPU instructions, while heap allocation involves more complex operations. However, not all data can be allocated on the stack. If the lifetime and memory footprint of a variable cannot be determined at compile time, a dynamic allocation on the heap occurs at runtime 