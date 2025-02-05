def getMaxMin(arr):
    max, min = arr[0],arr[0]
    for x in arr[1:]:
        if x > max:
            max = x
        if x < min:
            min = x
    return [max, min]

print(getMaxMin([3,4,6,7,3,2,4,6]))