def moveToLast(arr, index, nonZeroPointer):
    while index >= len(arr):
        # replaces all elems that are zero past index into 0s
        arr[nonZeroPointer] = 0
        nonZeroPointer+=1
        # base instance
        if nonZeroPointer == len(arr):
            return arr

    # moves all non-zero instances before 0s   
    if arr[index] != 0:
        arr[nonZeroPointer] = arr[index]
        nonZeroPointer+=1

    return moveToLast(arr, index+1, nonZeroPointer)

print(moveToLast([3,4,5,10,0,0,3,4,4,0,4], 0,0))