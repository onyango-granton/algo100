def firstNonRepeatingChar(word):
    firstCharOccurence = {}
    totalCharOccurence = {}
    firstOccurence = len(word)

    for index ,char in enumerate(word):
        totalCharOccurence[char] += 1
        if char not in firstCharOccurence:
            firstCharOccurence[char] = index

    for char, occurence in enumerate(firstCharOccurence):
        if totalCharOccurence[char] == 1 and occurence < firstOccurence:
            firstOccurence = occurence

    if firstOccurence == len(word):
        return ""
    else:
        return word[firstOccurence]
    
print(firstNonRepeatingChar("leetcode"))