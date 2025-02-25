def findPairWithSum(arr, target):
    intMap = {}

    for index, value in enumerate(arr):
        complement = target - value
        if complement in intMap:
            return [intMap[complement], index]
        intMap[value] = index

    return []
