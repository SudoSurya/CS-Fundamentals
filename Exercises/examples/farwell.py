
# farewell.py

# Import the greet module
import greet

# Get the name from the user
name = input("Enter the senior's name: ")

# Print the message variable from greet module
print(greet.message, end=' ')

# Call the greet function from greet module
greet.greet(name)

# Print the documentation string of the greet function
print("Documentation string:", greet.greet.__doc__)
