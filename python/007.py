def rotateArr(arr, positions):
#    positions = positions % len(arr)
   return arr[positions+2:]+arr[:len(arr)-positions]

print(rotateArr([3,4,5,6,7,8,9,10],3))