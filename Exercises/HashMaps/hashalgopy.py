def closureSum(inputNum):
    num_str = str(inputNum)
    length = len(num_str)
    sums = []
    R = 0
    L = length - 1
    while(L > R):
        sums.append(num_str[R] + num_str[L])
        R += 1
        L -= 1

    if(length % 2 != 0):
        sums.append(num_str[R])

    sums = [int(i) for i in sums]

    if(len(sums) == 1):
        print(sums[0])
    else:
       print(closureSum(sum(sums)))



# Example usage:
inputNum1 = 1039
inputNum2 = 2224555

hashValue1 = closureSum(inputNum1)
hashValue2 = closureSum(inputNum2)

print("Hash value for inputNum1:", hashValue1)
print("Hash value for inputNum2:", hashValue2)
