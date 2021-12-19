class Solution(object):
    def myAtoi(self, s):
        """
        :type s: str
        :rtype: int
        """
        splitted = s.split()
        if not len(splitted):
            return 0
        num = s.split()[0]
        print(num)
        int_num = 0
        sign = 1
        is_start = True
        for digit in num:
            if is_start and (digit == "-" or digit == "+"):
                if digit == '-':
                    sign = -1
                is_start = False
                continue
            is_start = False
            if digit in ['1', '2', '3', '4', '5', '6', '7', '8', '9', '0']:
                int_num *= 10
                int_num += int(digit)
                continue
            break
        int_num = int_num * sign
        if int_num > (1 << 31) - 1:
            return (1 << 31) - 1
        elif int_num < -1 << 31:
            return -1 << 31
        return int_num

