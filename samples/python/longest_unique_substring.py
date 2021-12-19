from collections import defaultdict


class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        letters_in_use = dict()
        current_last = 1
        max_distance = 0
        i = 0

        if s:
            letters_in_use[s[0]] = 0

        while i < len(s):
            char = s[i]
            while current_last < len(s):
                current_char = s[current_last]
                if current_char in letters_in_use:
                    max_distance = max(max_distance, current_last - i)
                    for j in range(i, letters_in_use[current_char]):
                        del letters_in_use[s[j]]
                    i = letters_in_use[current_char] + 1
                letters_in_use[current_char] = current_last
                current_last += 1
            max_distance = max(max_distance, current_last - i)
            i += 1
        return max_distance



print(Solution().lengthOfLongestSubstring("abcabcbb"))
