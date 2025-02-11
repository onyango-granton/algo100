def rotateArr(arr, positions):
    realPointer = 0
    copyArr =  [x for x in arr]


    for x in range (positions+1, len(arr)):
        arr[realPointer] = arr[x]
        realPointer+=1
    
    for x in range (0, positions+1):
        arr[realPointer] = copyArr[x]
        realPointer+=1

    return arr

print(rotateArr([3,4,5,6,7,8,9,10],3))