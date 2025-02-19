def findFirstNonRepeatingChar(word):
    charDic = {}
    index = len(word)
    for x in range(len(word)):
        charDic.setdefault(word[x],[]).append(x)

    for key in charDic:
        if len(charDic[key]) == 1:
            if charDic[key][0] < index:
                index = charDic[key][0]

    if index >= 0 and index < len(word):
        return index
    else:
        return -1

print(findFirstNonRepeatingChar("wordw"))