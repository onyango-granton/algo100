def findPairWithSum(arr, target):
    intMap = {}
    intList = list()
    indexList = list()

    for x in range(len(arr)):
        intMap[arr[x]] = target - arr[x]

    for x in intMap:
        if intMap[intMap[x]] != 0:
            intList.append([x,intMap[x]])
            break

    while x < 2:
        for y in range(len(arr)):
            if intList[x] == arr[y]:
                indexList.append(y)
                x += 1
            if x == 2:
                break

    return indexList
