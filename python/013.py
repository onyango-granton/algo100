def findMissingNum(arr):
    max, min = arr[0],arr[0]

    for x in arr:
        if x > max:
            max = x
        if x < min:
            min = x

    newArr = [None] * (max-min+1)

    for x in arr:
        newArr[x-min] = 1

    for x in range (len(newArr)):
        if newArr[x] == None:
            return x+1
        
    return -1

print(findMissingNum([3, 7, 1, 2, 8, 4, 5]))