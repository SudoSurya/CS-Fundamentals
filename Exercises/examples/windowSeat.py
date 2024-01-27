# Get the number of seats per row
seats_per_row = int(input("Enter the number of seats per row\n"))

# Check if seats_per_row is less than or equal to 0
if seats_per_row <= 0:
    print("Invalid Input")
else:
    # Get the seat number
    seat_number = int(input("Enter the seat number\n"))

    # Check if the seat number is valid
    if seat_number <= 0 or seat_number > 11 * seats_per_row:
        print("Invalid Seat Number")
    else:
        # Check if the seat is a window seat
        if seat_number % seats_per_row == 1 or seat_number % seats_per_row == 0:
            print("Window Seat")
        else:
            print("Not a Window Seat")
