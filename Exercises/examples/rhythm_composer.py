def is_prime(num):
    if num < 2:
        return False
    for i in range(2, int(num**0.5) + 1):
        if num % i == 0:
            return False
    return True

def find_prime(start, end):
    prime_numbers = [num for num in range(start, end + 1) if is_prime(num)]
    return prime_numbers

# Get user input
start = int(input())
end = int(input())

# Check for invalid input
if start > end:
    print("Invalid range")
elif start < 0 or end < 0:
    print("Invalid range")
else:
    result = find_prime(start, end)
    if not result:
        print("There are no prime numbers in this range")
    else:
        print(" ".join(map(str, result)))
