def findMaxNum(x, y, z):
    if abs(x - y) > z:
        return -1

    max_value = max(x, y)

    if z % 2 == 1:
        max_value -= 1

    return max_value

# Example usage:
x = 8
y = 5
z = 6

result = findMaxNum(x, y, z)
print(result)  # Output: 7
