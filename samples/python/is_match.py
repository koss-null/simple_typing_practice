class RegItem(object):
    def __init__(self, char, prev):
        self.char = char
        self.prev = prev

    def match(self, s, i):
        if self.char == ".":
            return True, (i+1,)

        if self.char == "*":
            match_func = self.prev.match
            next_matches = set([i])
            for j in range(i, len(s)):
                # print("checking {} {} {}".format(s[j], next_matches, match_func(s, j)))
                if not match_func(s, j)[0]:
                    # print(f"return True, {next_matches}")
                    return True, next_matches
                next_matches.add(j+1)
            return True, list(next_matches)

        if i >= len(s):
            return False, ()
        return self.char == s[i], (i+1,)

class Solution(object):
    def match(self, s, s_i, reg, reg_i):
        if reg_i == len(reg):
            return s_i == len(s)
        is_match, next_positions = reg[reg_i].match(s, s_i)
        if is_match:
            for pos in next_positions:
                if self.match(s, pos, reg, reg_i+1):
                    return True
        return False

    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        last_reg_item = None
        reg = list()
        for i_p_c, p_c in enumerate(p):
            if p_c == "*":
                continue
            if i_p_c + 1 < len(p) and p[i_p_c + 1] == "*":
                reg.append(RegItem("*", RegItem(p_c, None)))
                continue
            reg.append(RegItem(p_c, last_reg_item))

        return self.match(s, 0, reg, 0)


print(Solution().isMatch('b', 'b.bc'))

