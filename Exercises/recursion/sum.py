def sum (array):
    if len(array) == 1 :
        return array[0]

    return array[0] + sum(array[1:])

def number_of_paths(n):
    if n < 0:
        return 0
    if n == 1 or  n == 0:
        return 1
    return number_of_paths(n-1) + number_of_paths(n-2) + number_of_paths(n-3)
print(number_of_paths(3))
print(sum([1,2,3]))
