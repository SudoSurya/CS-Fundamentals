def average_celsius(numbers):
    celcius = []
    for num in numbers:
        celcius.append((num - 32) * 5/9)

    sum = 0.0
    for num in celcius:
        sum = sum + num

    return sum / len(celcius)

print(average_celsius([32, 212]))

# Path: average_celsius.py
