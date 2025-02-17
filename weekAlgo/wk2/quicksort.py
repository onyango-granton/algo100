def quicksort(arr):
    #basecase
    if len(arr) <= 1:
        return arr
    
    # calls a pivot as the last elem of the array
    pivot = arr[len(arr)-1]

    # creates a list of elements less or equal to the pivot
    left = [x for x in arr[:-1] if x<=pivot]

    # creates a list of elements greater to the pivot
    right = [x for x in arr[:-1] if x > pivot]

    # return recursive ca;;s to sort the left and right arrays recursively
    # till the basecase
    return quicksort(left) + [pivot] + quicksort(right)

print(quicksort([3,6,8,3,2,5,7,9,3,6,3,0,4]))