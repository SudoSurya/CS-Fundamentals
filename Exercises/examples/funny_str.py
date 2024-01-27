def funny_string(s):
    # Check if the length is less than 2
    if len(s) < 2:
        return "Invalid string"

    # Check if the string is funny
    for i in range(1, len(s)):
        if abs(ord(s[i]) - ord(s[i - 1])) != abs(ord(s[::-1][i]) - ord(s[::-1][i - 1])):
            return "Not funny"

    return "Funny"

# Get input from the user
input_string = input("Enter the string:")

# Call the function and display the result
result = funny_string(input_string)
print(result)
