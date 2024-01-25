# Step 1: Get the number of participants from the user
n = int(input("Enter the no. of participants details to be created: "))

# Step 2: Create an empty student list for storing details
student_list = []

# Step 3: Check if the number of participants is greater than 0
if n > 0:
    # Step 4: Get name, state, and age from the user using a loop
    for i in range(n):
        student = {}
        name = input("Name: ")
        state = input("State: ")
        age = int(input("Age: "))

        # Step 4: Check if age is valid
        if age <= 10 or age > 80:
            print("Invalid input")
            continue

        # Step 5: Create a dictionary and append it to the student list
        student["Name"] = name
        student["State"] = state
        student["Age"] = age
        student_list.append(student)

    # Step 6: Display the list of participants' details
    print("\nHere's the list of participants' details:")
    for participant in student_list:
        print(participant)

    # Step 7: Count the participants from each state
    state_count = {}
    for participant in student_list:
        if participant["State"] not in state_count.keys():
            state_count[participant["State"]] = 1
        else:
            state_count[participant["State"]] += 1

    # Step 8: Display the count of participants from each state
    print("\nCount of participants from each state:")
    for state, count in state_count.items():
        print(f"State: {state} Count: {count}")

# Step 9: If the number of participants is less than or equal to 0, display "Invalid input"
else:
    print("Invalid input")
