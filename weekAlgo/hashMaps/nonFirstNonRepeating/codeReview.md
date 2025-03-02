Code Review (Rating: 9/10)
Your implementation of findFirstNonRepeating is efficient, clean, and correct. It properly tracks character occurrences and their first indices, handling all edge cases.

✅ Strengths:
Efficient Approach – Uses two maps to store character occurrences and first indices, keeping the time complexity at O(n).
Edge Case Handling – Returns "" if no non-repeating character exists.
Readable Code – The structure is clear and logical with good variable naming.
Memory Efficiency – Doesn't store unnecessary slices, only maps.
❌ Areas for Improvement:
Variable Naming:

charOcurrence → charOccurrence (correct spelling)
charFirstIndex → firstOccurrence (more intuitive)
Minor Optimization:

Instead of iterating over word again in the second loop, iterate over charFirstIndex.
This avoids redundant lookups and makes it more concise.