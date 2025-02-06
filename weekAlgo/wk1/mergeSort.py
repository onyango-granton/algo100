def divide(arr):
    if len(arr) < 2:
        return arr
    midPoint = len(arr) / 2

    leftArr = arr[:midPoint]
    rightArr = arr[midPoint:]

    return conquer(leftArr, rightArr)

def conquer(leftArr, rightArr):
    mainArr = []
    mainP, rightP, leftP = 0,0,0
    while rightP < len(rightArr) and leftP < len(leftArr):
        if rightArr[rightP] <= leftArr[leftP]:
            mainArr[mainP] = rightArr[rightP]
            mainP += 1
            rightP += 1
        else:
            mainArr[mainP] = leftArr[leftP]
            mainP += 1
            leftP += 1
    while rightP < len(rightArr):
        mainArr[mainP] = rightArr[rightP]
        mainP += 1
        rightP += 1
    while leftP < len(leftArr):
        mainArr[mainP] = leftArr[leftP]
        mainP += 1
        leftP += 1
    return mainArr

print(divide([3,6,7,3,2,5,6,57,4,34]))