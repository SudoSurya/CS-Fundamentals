while True:
    # Get total number of people in the area
    total_people = int(input("Enter the total no of
                             people in the area: "))

    # Check if total_people is less than or equal to 0
    if total_people <= 0:
        print("Invalid Input")
        break

    # Get single-dose count
    single_dose_count = int(input("Single-dose count: "))

    # Check if single_dose_count is less than 0 or greater than total_people
    if single_dose_count < 0 or single_dose_count >
    total_people:
        print("Invalid Input")
        break

    # Get double-dose count
    double_dose_count = int(input("Double-dose count: "))

    # Check if double_dose_count is less than 0 or greater than total_people
    if double_dose_count < 0 or
    double_dose_count > total_people:
        print("Invalid Input")
        break

    # Check if total vaccinated counts are greater than total_people
    if single_dose_count + double_dose_count > total_people:
        print("Invalid Input")
        break

    # Calculate not vaccinated people count
    not_vaccinated_count = total_people -
    (single_dose_count + double_dose_count)

    # Calculate total vaccinated percentage
    vaccinated_percentage =
    (double_dose_count / total_people) * 100

    # Display results
    print("\nNot vaccinated people count: ",
          not_vaccinated_count)
    print("Total vaccinated percentage of people:
          {:.2f}".format(vaccinated_percentage))

    # Ask if the counselor wants to continue
    continue_option = input("\nDo you
                            want to continue (1
                                for yes, 0 for no): ")

    # Check if continue_option is not '0' or '1'
    if continue_option not in [0', '1']:
        print("Invalid Input")
        break

    # Check if continue_option is '0'
    if continue_option == '0':
        break
