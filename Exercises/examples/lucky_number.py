# Function to find if a number is lucky
def find_lucky(number):
    sm = 0
    while number != 0:
        temp = number % 10
        sm += temp
        number = number // 10

    if sm % 2 == 0:
        return 1
    else:
        return 0

# Get the number from the user
number = int(input("Enter the Number: "))

# Check if the number is less than or equal to 0
if number <= 0:
    print("Invalid Number")
else:
    # Call the find_lucky function
    luck = find_lucky(number)

    # Display the result based on the return value of find_lucky
    if luck == 1:
        print(str(number) + " is lucky")
    else:
        print(str(number) + " is not lucky")

