def reverseArr(arr):
    leftPointer = 0
    rightPointer = len(arr) - 1

    while leftPointer < rightPointer:
        arr[leftPointer], arr[rightPointer] = arr[rightPointer], arr[leftPointer]
        leftPointer += 1
        rightPointer -= 1

    return arr

print(reverseArr([2,4,5,6,7,8]))