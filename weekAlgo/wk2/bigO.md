The quicksort algorithm has a time complexity that varies based on the choice of the pivot and the distribution of the input data.

1. **Best and Average Case Time Complexity**: The best and average case time complexity of quicksort is O(n log n). This occurs when the pivot divides the array into two roughly equal halves at each recursive step, leading to a logarithmic number of levels of recursion, with each level processing all n elements.

2. **Worst Case Time Complexity**: The worst case time complexity is O(n^2). This happens when the pivot is consistently the smallest or largest element, leading to highly unbalanced partitions. For example, if the input array is already sorted or reverse sorted, and the last element is always chosen as the pivot, the algorithm will make n recursive calls, each processing n-1, n-2, ..., down to 1 elements.

3. **Space Complexity**: The space complexity of quicksort is O(log n) for the recursive stack space in the best and average cases, as the depth of recursion is logarithmic. However, in the worst case, the space complexity can degrade to O(n) due to the depth of recursion when the partitions are highly unbalanced.

Overall, quicksort is efficient for large datasets and is often faster in practice than other O(n log n) algorithms like mergesort or heapsort, due to its cache-friendly nature and lower constant factors. However, care must be taken with pivot selection to avoid the worst-case scenario.

source : [Big O Calculator](https://www.bigocalc.com/)