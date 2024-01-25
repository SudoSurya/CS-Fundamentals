from itertools import permutations

def solve(N,A):
    count = int(0)
    dic = {k: v for k, v in enumerate(A)}
    perms = list(permutations(dic.keys(), 4))
    tempmap = {}
    for i in perms:
        temp = list(i)
        if str(temp) not in tempmap:
            tempmap[str(temp)] = 1
            if A[temp[0]] > A[temp[1]] and A[temp[2]] < A[temp[3]] and temp[0] < temp[1] < temp[2] < temp[3]:
                count += 1
    return count % (10**9 + 7)

print(solve(4,[3,2,1,2,3]))
