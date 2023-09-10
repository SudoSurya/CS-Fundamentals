def biggest_number(arr):
    mx= arr[0]

    for i in range(len(arr)):
        mx= max(arr[i], mx)
    return mx

# Path: smallest.py
arr = [1, 2, 3, 4, 5]
print(biggest_number(arr))
