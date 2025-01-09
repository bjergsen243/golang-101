# Stack and Heap

## Stack

A stack is a consecutive block of memory, and every function call in thread of execution shares the same stack. Allocating memory on the stack is fast and simple. A stack pointer tracks the last location where memory was allocated; allocating additional memory is done by moving the stack pointer. When a function is invoked, a new stack frame is created for the function's data. Local variables are stored on the stack, along with paremeters passed into a function. Each new variable moves the stack pointer by the size is moved back to the beginning of the stack frame for the excited function, deallocating all of the stack memory that was used by that function's local variables and parameters

To store something on the stack, you have to know exactly how big it is at compile time. This is why the size is considered part of the type for an array. Because their sizes are known, they can be allocated on the stack instead of the heap. The size of a pointer type is also, and it is also stored on the stack
