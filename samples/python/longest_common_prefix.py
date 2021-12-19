class Solution(object):
    def check_all_strings(self, strs, i):
        for s in strs:
            if len(s) <= i:
                return False
            if s[i] != strs[0][i]:
                return False
        return True

    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        if not strs: return ""
        i = 0
        while self.check_all_strings(strs, i):
            i += 1
        return strs[0][:i]
