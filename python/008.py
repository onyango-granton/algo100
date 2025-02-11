def binarySort(arr, target, start, stop):
    if start > stop:
        return -1
    
    midIndex = start + (stop-start)//2

    if target == arr[midIndex]:
        return midIndex

    if target > arr[midIndex]:
        return binarySort(arr, target, midIndex + 1, stop)
    else :
        return binarySort(arr, target, start, midIndex - 1)
    
print(binarySort([2,4,5,6,7,8,10],10,0,6))