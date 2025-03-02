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

Python
🔹 Rating: 8.5/10
Your implementation of finding the first non-repeating character in a string is clear, efficient, and well-structured. It effectively tracks both first occurrences and total occurrences of each character in two separate dictionaries.

✅ Strengths:
Correct Logic:
It correctly finds the first non-repeating character based on its occurrence count and position.
Uses enumerate():
This is a clean and efficient way to iterate through the string while keeping track of indices.
Efficient Dictionary Lookups:
Uses two dictionaries (firstCharOccurence and totalCharOccurence), which makes character lookups O(1).
❌ Areas for Improvement:
Unnecessary Initialization of totalCharOccurence[char]

Instead of manually checking if char not in totalCharOccurence, a defaultdict(int) from collections can simplify this.
Looping Over firstCharOccurence Instead of word

firstCharOccurence already stores unique characters, but iterating over word preserves order while still checking non-repeating chars.
Returns an Empty String Instead of None (Optional Improvement)

In some use cases, returning None instead of "" could be more semantically correct.