def bubblesort(arr):
    for x in range(0,len(arr)):
        for x in range(0,len(arr)):
            if x + 1 < len(arr) and arr[x] > arr[x+1]:
                arr[x], arr[x+1] = arr[x+1], arr[x]
    return arr

print(bubblesort([2,3,5,6,7,2,3,4,2,4,453,5,4]))