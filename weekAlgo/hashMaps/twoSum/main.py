def twoSum(arr, target):
    diffDictionary = {}
    for x in range(len(arr)):
        if arr[x] in diffDictionary:
            return [diffDictionary[arr[x]],x]
        diffDictionary[target - arr[x]] = x

    return

print(twoSum([2, 7, 11, 15], 9))