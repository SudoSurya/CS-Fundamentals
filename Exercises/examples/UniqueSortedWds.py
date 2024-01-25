# Get input from the user
input_string = input("Enter the string:\n")

# Split the input string into a list of words/numbers
words_list = input_string.split()

# Use set to remove duplicates and then sort the data
unique_sorted_words = sorted(set(words_list), key=lambda x: x.lower())

# Display the result
print(" ".join(unique_sorted_words))
