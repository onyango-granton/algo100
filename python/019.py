def findPairWithSum(arr, target):
    intMap = {}

    for index, value in enumerate(arr):
        # calculating the complement for every n element
        complement = target - value
        #  checking for the precense of complement elem in dictionary
        if complement in intMap:
            return [intMap[complement], index]
        # adding each value and index to dictionary
        intMap[value] = index

    return []
