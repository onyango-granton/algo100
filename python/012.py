def mergeTwoArr(arr1, arr2):
    mainArr = list(range(len(arr1)+len(arr2)))

    mainP, arr1P, arr2P = 0,0,0

    while arr1P < len(arr1) and arr2P < len(arr2):
        print("here")
        if arr1[arr1P] < arr2[arr2P]:
            mainArr[mainP] = arr1[arr1P]
            arr1P+=1
        else:
            mainArr[mainP] = arr2[arr2P]
            arr2P+=1
        mainP+=1

    while arr1P < len(arr1):
        mainArr[mainP] = arr1[arr1P]
        arr1P+=1
        mainP+=1

    while arr2P < len(arr2):
        mainArr[mainP] = arr2[arr2P]
        arr2P+=1
        mainP+=1

    return mainArr

print(mergeTwoArr([2,34,5,5,3],[3,4,42,2,3]))

arr1 = [1, 3, 5]
arr2 = [2, 4, 6]
print(mergeTwoArr(arr1, arr2))