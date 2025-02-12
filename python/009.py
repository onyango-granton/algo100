def findFirstAndLast(arr, target):
    position = binarySort(arr,target, 0, len(arr)- 1)
    if position == -1:
        return [-1,-1]
    positionCopy = position
    first, last = position
    while arr[position] == target:
        first = position
        position-=1
    while arr[positionCopy] == target:
        last = positionCopy
        positionCopy+=1
    return [first, last]

def binarySort(arr, target, start,stop):
    if start > stop:
        return -1
    
    midVal = start + (stop-start) // 2

    if target == arr[midVal]:
        return midVal

    elif target > arr[midVal]:
        return binarySort(arr, target, midVal + 1, stop)

    else:
        return binarySort(arr, target, start, midVal - 1)
    

print(binarySort([2,3,4,5,6,7,8],8,0,6))
print(findFirstAndLast([3,4,5,5,5,6,7,8],5))
