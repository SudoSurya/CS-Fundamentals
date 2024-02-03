# CS50 Memory Lecture

## Hexadecimal aka Base 16

- Hexadecimal: 0 1 2 3 4 5 6 7 8 9 A B C D E F

## Pointers

- Pointers: variables that store addresses of other variables
- `int *p` declares a pointer to an integer
- `p = &n` assigns the address of `n` to `p`
- `*p` dereferences `p` to get the value at the address it points to

## Strings

- Strings: arrays of characters
- char *s = "HI!" declares a pointer to a string

## Pointer Arithmetic

- Pointer arithmetic: `p++` increments `p` by the size of the type it points to
- `*(p + 1)` gets the value at the address `p + 1`
-  `p[i]` is equivalent to `*(p + i)`

