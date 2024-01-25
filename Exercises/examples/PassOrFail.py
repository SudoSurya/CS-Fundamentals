def count_pass_fail_subjects(num_subjects, marks):
    if num_subjects <= 0:
        print("Invalid no. of subjects")
        return

    pass_count = 0
    fail_count = 0

    for mark in marks:
        if mark < 0 or mark > 100:
            print("Invalid mark")
            return

        if mark >= 50:
            pass_count += 1
        else:
            fail_count += 1

    print("No. of subjects passed:", pass_count)
    print("No. of subjects failed:", fail_count)

# Get input from the user
try:
    num_subjects = int(input("Enter the no. of subjects: "))
    if num_subjects <= 0:
        print("Invalid no. of subjects")
    else:
        marks = []
        print("Enter the marks:")
        for i in range(num_subjects):
            mark = int(input())
            if mark < 0 or mark > 100:
                print("Invalid mark")
                exit()
            marks.append(mark)

        count_pass_fail_subjects(num_subjects, marks)

except ValueError:
    print("Invalid input. Please enter a valid number.")
