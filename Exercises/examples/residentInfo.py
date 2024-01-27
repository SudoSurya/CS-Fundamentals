# Function to validate age
def validate_age(age):
    return 21 <= age <= 58

# Function to validate band
def validate_band(band):
    return band in ['A', 'B', 'C', 'D']

# Function to display residents and filter by band
def display_residents(residents, band_of_interest):
    header = ('NAME', 'AGE', 'DESIGNATION', 'BAND')
    print(header)

    if band_of_interest == '':
        for resident in residents:
            print(resident)
        return
    residents_filtered = []
    for resident in residents:
        if resident[-1] == band_of_interest:
            residents_filtered.append(resident)

    if len(residents_filtered) <= 0:
        print(f"No resident under this band")
    else:
        for resident in residents_filtered:
            print(resident)

# Get the number of residents from the user
num_of_residents = int(input("No of Residents: "))

# Validate the number of residents
if num_of_residents <= 0:
    print("Invalid")
    exit()

# Initialize the list to store residents
List_of_Residents = []

# Get details of each resident
for i in range(1, num_of_residents + 1):
    print(f"\nResident {i}:")
    name = input("Name: ")
    age = int(input("Age: "))
    designation = input("Designation: ")
    band = input("Band: ")

    # Validate age and band
    if not validate_age(age) or not validate_band(band):
        print("Invalid")
        exit()

    # Create a tuple and insert into the list
    resident_tuple = (name, age, designation, band)
    List_of_Residents.append(resident_tuple)

# Display all residents
display_residents(List_of_Residents, '')

# Get the band of interest from the user
band_of_interest = input("\nEnter your band of interest: ")

# Validate the band of interest
if not validate_band(band_of_interest):
    print("Invalid")
    exit()

# Display residents filtered by band of interest
display_residents(List_of_Residents, band_of_interest)
