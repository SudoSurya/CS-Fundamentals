def is_sorted(arr):
    for i in range(len(arr)-1):
        if arr[i] > arr[i+1]:
            return False
    return True

def linear_search(arr,needle):
    for i in range(len(arr)):
        if arr[i] == needle:
            return i
    return -1
arr = [1,1,2,4,3]
needle = 3
print(linear_search(arr,needle))
print(is_sorted(arr))
