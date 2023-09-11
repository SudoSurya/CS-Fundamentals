def average(numbers):
    sum = 0.0
    count_nums = 0
    for num in numbers:
        if(num % 2 == 0):
            sum = sum + num
            count_nums = count_nums + 1
    return sum / count_nums

# Path: oddMean.py

print(average([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]))
