# Step 1: Get the number of students from the user
no = int(input("Enter number of students: "))

# Step 2: Open the file in write mode to write the details
file = open('output_data.txt', 'w')

# Step 3: Write student details to the file
for i in range(1, no + 1):
    print("For student", i)
    name = input("Enter name: ")
    score = input("Enter the score: ")
    data_format = "Name: " + name + " Score: " + score
    file.write(data_format + "\n")

# Close the file
file.close()

# Step 4: Read data from the file and display it
read_file = open("output_data.txt", "r")
data = read_file.read()
print(data)
read_file.close()
