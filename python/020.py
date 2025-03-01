def zeroSumSubArray(arr):
    subArr = [[]]
    for x in range(len(arr)):
       start = x
       stop = x
       sum = 0
       for i in range(x, len(arr)): 
           sum += arr[i]
           if sum == 0:
               stop = i
               subArr.append([start, stop+1])
               start = i
               sum = 0

    return subArr
        


print(zeroSumSubArray([6, -1, -3, 4, -2, 2, 4, 6, -12, -7]))