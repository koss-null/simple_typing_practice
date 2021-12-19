class Solution(object):
    def intToRoman(self, num):
        """
        :type num: int
        :rtype: str
        """
        g2r = {
            1: "I",
            5: "V",
            10: "X",
            50: "L",
            100: "C",
            500: "D",
            1000: "M",
            4: "IV",
            9: "IX",
            40: "XL",
            90: "XC",
            400: "CD",
            900: "CM",
        }
        available_nums = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
        available_nums_i = 0
        result = list()
        for subst_num in available_nums:
            while num - subst_num >= 0:
                result.append(g2r[subst_num])
                num -= subst_num
        return "".join(result)
