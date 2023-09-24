def partition(arr,left,right):
    pIdx = left
    pivot = arr[pIdx]
    right -= 1

    while arr[left] < pivot:
        left += 1
    while arr[right] > pivot:
        right -= 1

    if(left >= right):
        return pIdx
    else:
        arr[left],arr[right] = arr[right],arr[left]

    arr[pIdx],arr[left] = arr[left],arr[pIdx]

    return left

def quickSort(arr,left,right):
    if right - left <= 0:
        return

    pIdx = partition(arr,left,right)
    quickSort(arr,left,pIdx-1)
    quickSort(arr,pIdx+1,right)


# Path: QS.py

arr = [5,4,3,2,1]
print(arr)
quickSort(arr,0,len(arr))
print(arr)
