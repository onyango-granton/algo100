def print_subsets(arr):
    n = len(arr)
    total_subsets = 2 ** n  # Number of subsets

    for i in range(total_subsets):
        subset = []
        for j in range(n):
            if (i & (1 << j)) > 0:  # Check if jth element is included
                subset.append(arr[j])
        print(subset)

# Example usage
arr = [1, 2, 3]
print_subsets(arr)