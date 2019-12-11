# Notes

In go a pointer is represented by the *(asterisk) character followed by the type of the stored value.

**\*** is also used to *dereference* pointer variables. Dereferencing a pointer gives us access to the value the pointer points to.

When we write **\*xPtr = 0** we are saying **store the int 0 in the memory location xPtr refers to**.

If we try 'xPtr = 0' instead, we will get a compiler error because xPtr is not an int, its a '*int', which can be given another '*int'

## New

**new** takes a **type** as an argument, allocates enough memory to fit a value of that **type** and returns a **pointer** to it.

