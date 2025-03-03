# from functools import lru_cache

# @lru_cache(None)
def fibonacci(position, memo):
    #base case
    if position < 2:
        return position
    #memo case
    if position in memo:
        return memo[position]
    
    # add to memo
    memo[position] = fibonacci(position-1, memo) + fibonacci(position-2,memo)

    return memo[position]

print(fibonacci(6,{}))