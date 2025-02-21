def quicksort(arr):
    if len(arr) <= 1:
        return arr

    pivot = arr[len(arr)-1]
    lessArr = [x for x in arr[:len(arr)-1] if x <= pivot]
    moreArr = [x for x in arr[:len(arr)-1] if x > pivot]

    return quicksort(lessArr)+[pivot]+quicksort(moreArr)


def twoSum(arr,target):
    arr = quicksort(arr)

    l,r = 0,0

    while l < r:
        if arr[l] + arr[r] == target:
            


print(quicksort([2, 3, 2, 5, 7, 8, 5, 9, 6]))