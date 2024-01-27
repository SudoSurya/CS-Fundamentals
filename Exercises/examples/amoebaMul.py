def fibonacci(month):
    if month < 1 or month > 12:
        print("Invalid month")
        return

    # Initialize the Fibonacci series
    fib_series = [0, 1]

    # Calculate Fibonacci series up to the given month
    for i in range(2, month + 1):
        fib_series.append(fib_series[i - 1] + fib_series[i - 2])

    # Display the amoeba size for the given month
    print(f"The amoeba size is {fib_series[month - 1]}")

# Get user input for the month
month = int(input("Enter the month as numeric value:\n"))

# Call the function
fibonacci(month)
