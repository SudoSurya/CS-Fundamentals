def maxOnes(N, A):
  maxOnesWithoutFlip = 0
  countOnes = 0
  for i in range(N):
    if A[i] == 1:
      countOnes += 1
      maxOnesWithoutFlip = max(maxOnesWithoutFlip, countOnes)
    else:
      countOnes = 0
  maxOnesWithFlip = 0
  countZeros = 0
  for i in range(N):
    if A[i] == 0:
      countZeros += 1
      maxOnesWithFlip = max(maxOnesWithFlip, N - countZeros)
    else:
      countZeros = 0

  # Return the maximum of the two values.
  return max(maxOnesWithoutFlip, maxOnesWithFlip)

N = 7
AD = [1, 0, 0, 1, 0,0,1]
result = maxOnes(N, AD)
print(result)  # Output: 4

