class Solution(object):
    def reverse(self, c):
        if c == "]":
            return "["
        if c == ")":
            return "("
        if c == "}":
            return "{"

    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        stack = list()
        for c in s:
            if c in "({[":
                stack.append(c)
            else:
                last_char = stack[-1]
                stack = stack[:-1]
                if last_char != self.reverse(c):
                    return False
        return not len(stack)

