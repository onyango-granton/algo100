def insertionSort(arr):
    for x in range (0, len(arr)):
        pointer = x
        while pointer - 1 >= 0 and arr[pointer - 1] > arr[pointer]:
            arr[pointer - 1], arr[pointer] = arr[pointer], arr[pointer-1]
            pointer-=1
    return arr

print(insertionSort([4,2,5,6,6,23,23,3,42,6]))