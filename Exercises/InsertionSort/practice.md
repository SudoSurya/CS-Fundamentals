# example

A = [31, 41, 59, 26, 41, 58]

- illustrate the operation of insertion sort on A

[]: # (1) [31, 41, 59, 26, 41, 58]
[]: # (2) [31, 41, 59, 26, 41, 58]
[]: # (3) [26, 31, 41, 59, 41, 58]
[]: # (4) [26, 31, 41, 41, 59, 58]
[]: # (5) [26, 31, 41, 41, 58, 59]
[]: # (6) [26, 31, 41, 41, 58, 59]

# practice

for i = 1 ... n:
    insert A[i] into sorted array A[1 ... i - 1]
    key = A[i]
    j = i - 1
    while j > 0 and A[j] > key:
        A[j + 1] = A[j]
        j = j - 1
    A[j + 1] = key


