def findMissingNum(arr):
    # change name of variables as max and min are inbuilt in python
    max_num, min_num = max(arr), min(arr)

# Not needed thus
    # for x in arr:
    #     if x > max:
    #         max = x
    #     if x < min:
    #         min = x

    # create a num set, set in python have one elem rep 8,8 comes to 8
    # newArr = [None] * (max-min+1)
    num_set = set(arr)

    for x in range(min_num, max_num+1):
        if x not in num_set:
            return x


    return -1

print(findMissingNum([3, 7, 1, 2, 8, 4, 5]))