def quicksort(arr):
    if len(arr) <= 1:
        return arr

    pivot = arr[len(arr)-1]
    lessArr = [x for x in arr[:len(arr)-1] if x <= pivot]
    moreArr = [x for x in arr[:len(arr)-1] if x > pivot]

    return quicksort(lessArr)+[pivot]+quicksort(moreArr)


def twoSum(arr,target):
    arr = quicksort(arr)

    l,r = 0,len(arr)-1

    intDict = {}

    sumList = [[]]

    for x in range(len(arr)):
        intDict[x] = False

    while l < r:
        if arr[l] + arr[r] == target:
            if not intDict[l] and not intDict[r]:
                sumList.append([l,r])
                intDict[l] = True
                intDict[r] = True
            l += 1
            r -= 1
        elif arr[l] + arr[r] > target:
            if arr[l] > target and l - 1 >= 0:
                l -= 1
            else:
                r -= 1
        elif arr[l] + arr[r] < target:
            if arr[r] < target and r + 1 < len(arr):
                r += 1
            else:
                l += 1
            
    return sumList



print(twoSum([2, 3, 2, 5, 7, 8, 5, 9, 6],9))