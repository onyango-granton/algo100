def removeDuplicates(arr):
    return list(dict.fromkeys(arr))

print(removeDuplicates([3,5,6,7,8,6,4,4]))