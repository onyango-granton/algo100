def divide(arr):
    if len(arr) < 2:
        return arr
    mid = len(arr) // 2
    leftArr = divide(arr[:mid])
    rightArr = divide(arr[mid:])
    return conquer(leftArr, rightArr)

def conquer(left,right):
    mainArr = list(range(len(left)+len(right)))
    mPointer,lPointer, rPointer = 0,0,0
    while lPointer < len(left) and rPointer < len(right):
        if left[lPointer] < right[rPointer]:
            mainArr[mPointer] = left[lPointer]
            mPointer += 1
            lPointer += 1
        else:
            mainArr[mPointer] = right[rPointer]
            mPointer += 1
            rPointer += 1
    while lPointer < len(left):
        mainArr[mPointer] = left[lPointer]
        mPointer += 1
        lPointer += 1
    while rPointer < len(right):
        mainArr[mPointer] = right[rPointer]
        mPointer += 1
        rPointer += 1
    return mainArr

def findCommonElems(arrOne, arrTwo):
    sortArrOne, sortArrTwo = divide(arrOne), divide(arrTwo)
    onePointer, twoPointer = 0,0
    resultArr = []

    print(len(resultArr))

    if len(sortArrTwo) < len(sortArrOne):
        copyArr = sortArrTwo
        sortArrTwo = sortArrOne
        sortArrOne = copyArr

    while onePointer < len(sortArrOne):
        if sortArrOne[onePointer] == sortArrTwo[twoPointer]:
            if len(resultArr) == 0:
                resultArr.append(sortArrOne[onePointer])
                onePointer += 1
                twoPointer += 1
            elif sortArrOne[onePointer] != resultArr[len(resultArr)-1]:
                resultArr = resultArr.append(sortArrOne[onePointer])
                onePointer += 1
                twoPointer += 1
            else:
                onePointer += 1
                twoPointer += 1
        elif sortArrOne[onePointer] < sortArrTwo[twoPointer]:
            onePointer += 1
        elif sortArrTwo[twoPointer] < sortArrOne[onePointer]:
            twoPointer += 1

    return resultArr

print(divide([6,3,2,8,9,2,5]))
print(findCommonElems([4,5,6,7,3,23,5],[2,5,6,2,4,4,2]))