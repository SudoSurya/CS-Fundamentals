def is_armstrong_number(num):
    # initialize sum
    sum = 0

    # find the sum of the cube of each digit
    temp = num
    while temp > 0:
       digit = temp % 10
       sum += digit ** 3
       temp //= 10
    return num == sum

def display_armstrong_numbers(start, end):
    if start < 0 or end < 0:
        print("Starting and ending numbers must be greater than or equal to zero")
        return
    elif start > end:
        print("Invalid input!! Ending number should be greater than starting number")
        return
    armstrong_numbers = []
    for num in range(start, end + 1):
        if is_armstrong_number(num):
            armstrong_numbers.append(num)

    if armstrong_numbers:
        print(f"Armstrong numbers between {start} and {end} are:")
        for armstrong_num in armstrong_numbers:
            print(armstrong_num)
    else:
        print(f"There is no Armstrong number between {start} and {end}")

# Get input from the user
n1, n2 = map(int, input("Enter the starting and ending numbers:\n").split())

# Call the function to display Armstrong numbers
display_armstrong_numbers(n1, n2)
