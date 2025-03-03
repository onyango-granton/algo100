def fibonacci(position):
    memo = {}
    #base case
    if position < 2:
        return position
    #memo case
    if position in memo:
        return memo[position]
    
    # add to memo
    memo[position] = fibonacci(position-1) + fibonacci(position-2)

    return memo[position]

print(fibonacci(6))