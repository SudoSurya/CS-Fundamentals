def calculate_password(plots):
    even_plots = 0
    odd_plots = 0
    for i in range(len(plots)):
        if isEven(i):
            even_plots += plots[i]
        else:
            odd_plots += plots[i]

    return (even_plots + odd_plots) / 2


def isEven(num):
    return num % 2 == 0
# Get the total number of plots from the user
total_plots = int(input("Enter the total no.of plots: "))

# Check for valid input size
if 1 <= total_plots <= 20:
    # Get the plot numbers from the user
    plots = []
    print("Enter the numbers of each plot: ")
    for _ in range(total_plots):
        plot = float(input())

        # Check for valid input (positive values only)
        if plot <= 0:
            print("Invalid Input")
            exit()

        plots.append(plot)

    # Call the function to calculate the password
    password = calculate_password(plots)

    # Display the result with 2 decimal places
    print(f"The password for the file is: {password:.2f}")
else:
    print("Invalid Input")
