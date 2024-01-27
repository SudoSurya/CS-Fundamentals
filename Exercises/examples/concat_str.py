
def concat_string(string1, string2):
    new_string = string1[2:] + string2[2:]
    length = len(new_string)
    return new_string, length

# Get user input
string1 = input("Enter String1: ")
string2 = input("Enter String2: ")

# Call the function
result, length = concat_string(string1, string2)

# Display the result
print("The concatenated string:", result)
print("The length of the new string is:", length)
