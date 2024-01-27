def find_not_working_doctors(all_doctors, working_doctors):
    # Use list comprehension to find the doctors who are not working
    not_working_doctors = [doctor_id for doctor_id in all_doctors if doctor_id not in working_doctors]
    return not_working_doctors

# Get the size of the first list from the user
size_all_doctors = int(input())

# Check for valid list size
if size_all_doctors <= 0:
    print("Invalid list size")
    exit()

# Get the elements of the first list
all_doctors = []
for _ in range(size_all_doctors):
    doctor_id = int(input())

    # Check for valid doctor id
    if doctor_id <= 0:
        print("Invalid Id")
        exit()

    all_doctors.append(doctor_id)

# Get the size of the second list from the user
size_working_doctors = int(input())

# Check for valid list size
if size_working_doctors < 0:
    print("Invalid list size")
    exit()

# Get the elements of the second list
working_doctors = []
for _ in range(size_working_doctors):
    doctor_id = int(input())

    # Check for valid doctor id
    if doctor_id <= 0:
        print("Invalid Id")
        exit()

    working_doctors.append(doctor_id)

# Call the function to find not working doctors
not_working_doctors = find_not_working_doctors(all_doctors, working_doctors)

# Display the result
print("Not Working Doctors' IDs are:")
print(" ".join(map(str, not_working_doctors)))
