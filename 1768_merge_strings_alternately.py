def mergeAlternately(word1, word2):
    """
    :type word1: str
    :type word2: str
    :rtype: str

    Example 1:
    Input: word1 = "abc", word2 = "pqr"
    Output: "apbqcr"
    Explanation: The merged string will be merged as so:
    word1:  a   b   c
    word2:    p   q   r
    merged: a p b q c r
    """
    merged, idx = "", 0
    while idx < min(len(word1), len(word2)):
        merged = merged + word1[idx] + word2[idx]
        idx = idx + 1
    return merged + word1[idx:] + word2[idx:]

    # for idx in range(minlen):
    #     merged += word1[idx] + word2[idx]
    # if len(word1) < len(word2):
    #     merged += word2[idx+1:]
    # else:
    #     merged += word1[idx+1:]
    # return merged

print(mergeAlternately("abc", "pqr"))
print(mergeAlternately("ab", "pqrs"))
print(mergeAlternately("abcd", "pq"))
