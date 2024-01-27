# Get user inputs
dead_count = int(input("Dead Count:\n"))

# Check if dead count is negative
if dead_count < 0:
    print("Invalid input")
    exit()

injured_count = int(input("Injured Count:\n"))

# Check if injured count is negative
if injured_count < 0:
    print("Invalid input")
    exit()

safe_count = int(input("Safe Count:\n"))

# Check if safe count is negative
if safe_count < 0:
    print("Invalid input")
    exit()

# Generate and print the report
print("\nTSUNAMI REPORT OF JAPAN\n\nThe number of people\n\nDead: {}\nInjured: {}\nSafe: {}".format(dead_count, injured_count, safe_count))
print("\nPlease help the people who are suffering!!!")
