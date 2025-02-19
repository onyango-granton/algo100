def quicksort(arr):
    if len(arr) <= 1:
        return arr
    pivot = arr[len(arr)-1]

    lessArr = [x for x in arr[:len(arr)-1] if x <= pivot]
    moreArr = [x for x in arr[:len(arr)-1] if x > pivot]

    return quicksort(lessArr) + [pivot] + quicksort(moreArr)

def findLongestConsecutiveSequence(arr):
    sortArr = quicksort(arr)

    indexArr = [sortArr[0]]
    listArr = [[]]

    for x in range(len(sortArr)):
        if sortArr[x] - indexArr[len(indexArr)-1] == 1:
            indexArr.append(sortArr[x])
        else:
            listArr.append(indexArr)
            indexArr = [sortArr[x]]

        if x == len(sortArr)-1:
            listArr.append(indexArr)

    return max(len(x) for x in listArr)
        

print(findLongestConsecutiveSequence([0,3,7,2,5,8,4,6,0,1]))