import array

balance = array.array("i", [1, 2, 3])
balance.insert(2, 150)

balance.insert(2, 150)
balance.insert(2, 150)
balance.insert(2, 150)
print(balance)
for x in balance:
    print(x)
