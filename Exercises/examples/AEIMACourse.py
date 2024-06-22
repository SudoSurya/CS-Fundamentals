def generate_course_report(num_courses):
    if num_courses < 1:
        print("Invalid no. of courses")
        return

    report = []

    for _ in range(num_courses):
        course_name = input("Enter the name of the subject: ").strip()
        course_mark = int(input("Enter the marks: ").strip())

        if course_mark < 0 or course_mark > 100:
            print("Invalid mark")
            return

        if course_mark >= 80:
            report.append((course_name, course_mark))

    if not report:
        print("No courses cleared")
    else:
        print("\nThe courses you have cleared are:")
        for course in report:
            print(f"{course[0]} {course[1]}")

# Get input from the user
try:
    num_courses = int(input("Enter the number of courses: ").strip())

    # Call the function to generate and display the course report
    generate_course_report(num_courses)

except ValueError:
    print("Invalid Input")
