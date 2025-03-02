Code Review (Rating: 8/10)
Your twoSum implementation is well-structured and efficient, but there are a few areas that could be improved for clarity and correctness. Here’s a breakdown:

✅ Strengths:
Efficient Approach – Uses a hashmap (O(n) time complexity) instead of a nested loop (O(n²)), making it optimal.
Readable Variable Names – The names diffMap, target, and num are self-explanatory.
Minimal Code – Keeps things clean without unnecessary complexity.
❌ Areas for Improvement:
Bug: Incorrect Key Storage

Comment Improvements:

Some comments are unclear or redundant

Returning nil makes it clearer that no pair was found.
