def findLargestSum(arr):
    sum = arr[0]
    maxSum = arr[0]

    for x in range(1,len(arr)):
        # sets the value of sum depending on value of current element
        sum = max(sum + arr[x], arr[x])
        # sets value of maxSum btwn sum and maxSum
        maxSum = max(maxSum, sum)
    
    return maxSum

print(findLargestSum([-2,1,-3,4,-1,2,1,-5,4]))