def longest_common_prefix(words):
    min_word = min(words)
    end = 0
    for i in range(len(min_word)):
        if len(set([word[i] for word in words])) != 1:
            break
        end += 1
    return min_word[:end]
