def bubblesort(arr):
    for x in range(0,len(arr)):
        swapped =  False
        for x in range(0,len(arr)):
            if x + 1 < len(arr) and arr[x] > arr[x+1]:
                arr[x], arr[x+1] = arr[x+1], arr[x]
                swapped = True
        if not swapped:
            break
    return arr

def recurseBubbleSort(arr, iteration, pointer):
    if pointer == len(arr):
        pointer = 0
        iteration+=1

    if iteration == len(arr):
        return arr
    
    if pointer + 1 < len(arr) and arr[pointer] > arr[pointer + 1]:
        arr[pointer], arr[pointer + 1] = arr[pointer + 1], arr[pointer]
    return recurseBubbleSort(arr, iteration, pointer+1)

print(bubblesort([2,3,5,6,7,2,3,4,2,4,453,5,4]))