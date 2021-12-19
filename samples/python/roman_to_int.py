class Solution(object):
    def romanToInt(self, s):
        g2r = {
                "I": 1,
                "V": 5,
                "X": 10,
                "L": 50,
                "C": 100,
                "D": 500,
                "M": 1000,
                "IV": 4,
                "IX": 9,
                "XL": 40,
                "XC": 90,
                "CD": 400,
                "CM": 900,
        }
        available_nums = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"]
        s_i = 0
        result = 0
        for subst_num in available_nums:
            while (
                (len(subst_num) == 1 and s_i < len(s) and s[s_i] == subst_num) or
                (len(subst_num) == 2 and s_i + 1 < len(s) and s[s_i: s_i + 2] == subst_num)
            ):
                result += g2r[subst_num]
                s_i += len(subst_num)
        return result
