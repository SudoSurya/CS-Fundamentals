# Step 1: Get the weight from the user
weight = int(input("Enter the weight of the person(kg):"))

# Step 2: Get the height from the user
height = float(input("Enter the height of the person(m):"))

# Step 3: Check if both height and weight are greater than 0
if weight > 0 and height > 0:
    # Calculate BMI and round to 1 decimal place
    bmi = weight / (height * height)
    BMI = round(bmi, 1)

    # Step 4: Check the BMI range and display the associated risk
    if BMI >= 27.5:
        print("Your BMI is {:.1f} (High Risk)".format(BMI))
    elif 23 <= BMI <= 27.4:
        print("Your BMI is {:.1f} (Moderate Risk)".format(BMI))
    elif 18.5 <= BMI <= 22.9:
        print("Your BMI is {:.1f} (Low Risk)".format(BMI))
    elif BMI < 18.5:
        print("Your BMI is {:.1f} (Risk of nutritional deficiency diseases)".format(BMI))

# Step 9: If height or weight is not greater than 0, display the error message
else:
    print("Provide a valid input")
