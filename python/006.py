def moveToLast(arr, index, nonZeroPointer):
    while index >= len(arr):
        arr[nonZeroPointer] = 0
        nonZeroPointer+=1
        if nonZeroPointer == len(arr):
            return arr
        
    if arr[index] != 0:
        arr[nonZeroPointer] = arr[index]
        nonZeroPointer+=1

    return moveToLast(arr, index+1, nonZeroPointer)

print(moveToLast([3,4,5,10,0,0,3,4,4,0,4], 0,0))