def quickselect(arr,left,right,k):
    if right - left <= 0:
        return arr[left]

    pIdx = partition(arr,left,right)

    if k < pIdx:
        return quickselect(arr,left,pIdx-1,k)
    elif k > pIdx:
        return quickselect(arr,pIdx+1,right,k)
    else:
        return arr[pIdx]

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

arr = [1,2,31,21,1]
print(quickselect(arr,0,len(arr),len(arr)-2))
