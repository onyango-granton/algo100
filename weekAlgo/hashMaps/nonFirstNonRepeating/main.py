def firstNonRepeatingChar(word):
    firstCharOccurence = {}
    totalCharOccurence = {}
    firstOccurence = len(word)

    for index ,char in enumerate(word):
        if char not in totalCharOccurence:
            totalCharOccurence[char] = 0
        totalCharOccurence[char] += 1
        if char not in firstCharOccurence:
            firstCharOccurence[char] = index

    for char in firstCharOccurence:
        if totalCharOccurence[char] == 1 and firstCharOccurence[char] < firstOccurence:
            firstOccurence = firstCharOccurence[char]

    if firstOccurence == len(word):
        return ""
    else:
        return word[firstOccurence]
    
print(firstNonRepeatingChar("leetcode"))