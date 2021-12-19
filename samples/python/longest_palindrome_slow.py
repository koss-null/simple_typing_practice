class Solution(object):
    def findPalindromeLength(self, i):
        between = (float(int(i)) == i)
        if between:
            shift = 1
        else:
            shift = 0.5
        while i - shift >= 0 and i + shift < len(self.s):
            if self.s[int(i - shift)] == self.s[int(i + shift)]:
                shift += 1
            else: break
        if between: shift -= 1
        else: shift -= 0.5
        return 2 * shift + (1 if between else 0)


    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        self.s = s

        max_palindrome_length, max_palindrome_idx = -1, -1
        i = 0.
        while i < len(s):
            palindrome_length = self.findPalindromeLength(i)
            if palindrome_length > max_palindrome_length:
                max_palindrome_idx = i
                max_palindrome_length = palindrome_length[-1]
            i += 0.5
        # print(palindrome_length)

        if float(int(max_palindrome_idx)) == max_palindrome_idx:
            return s[int(max_palindrome_idx) - int(max_palindrome_length / 2) : int(max_palindrome_idx) +
                    int(max_palindrome_length / 2) + 1]
        else:
            return s[int(max_palindrome_idx) - int(max_palindrome_length / 2) + 1: int(max_palindrome_idx) +
                    int(max_palindrome_length / 2 + 1)]


def main():
    print(Solution().longestPalindrome(""))


if __name__ == "__main__":
    main()


