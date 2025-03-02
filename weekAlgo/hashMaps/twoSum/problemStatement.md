Problem: Given an array of integers nums and an integer target, return the indices of the two numbers that add up to target.
Example:

plaintext
Copy
Edit
Input: nums = [2, 7, 11, 15], target = 9  
Output: [0, 1]  // Because nums[0] + nums[1] = 2 + 7 = 9
✅ Concept: Use a hashmap to store seen numbers and their indices for O(1) lookups.