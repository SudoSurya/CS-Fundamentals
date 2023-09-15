def double_array(arr,index):
    if index > len(arr) - 1:
        return
    arr[index] = arr[index] * 2
    double_array(arr,index+1)

array = [1, 2, 3, 4]
double_array(array, 0)
print(array)
