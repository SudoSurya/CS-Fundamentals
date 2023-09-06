def max_consistent_rating(ratings):
  max_rating = 0
  current_rating = 0
  for rating in ratings:
    if current_rating < 0:
      current_rating = 0
    current_rating += rating
    max_rating = max(max_rating, current_rating)
  return max_rating

ratings = [8,8,-1,-4,4,-2,0,1,4,-5]
print(max_consistent_rating(ratings))

