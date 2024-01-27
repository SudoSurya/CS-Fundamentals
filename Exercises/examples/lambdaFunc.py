subject1 = float(input("Enter marks for subject1: "))
subject2 = float(input("Enter marks for subject2: "))
subject3 = float(input("Enter marks for subject3: "))

cal = lambda s1, s2, s3: (s1 + s2 + s3) / 3

percentage = cal(subject1, subject2, subject3)
print(f"Percentage is {percentage}")
