class Solution(object):
    def get_digit_list(self, x):
        x = abs(x)
        digits = list()
        while x:
            digits.append(x % 10)
            x //= 10
        return digits

    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        if not x: return 0
        sign = x / abs(x)
        x_digits = self.get_digit_list(x)
        if sign > 0:
            ethalon_digits = list(reversed(self.get_digit_list((1 << 31) - 1)))
        else:
            ethalon_digits = list(reversed(self.get_digit_list(-1 << 31)))
        print(ethalon_digits)
        print(x_digits)
        if len(ethalon_digits) < len(x_digits):
            print(1)
            return 0
        if len(ethalon_digits) == len(x_digits):
            print(2)
            for i in range(0, len(x_digits) - 1):
                print(f"checking {ethalon_digits[i]} < {x_digits[i]}")
                if ethalon_digits[i] < x_digits[i]:
                    return 0
                elif ethalon_digits[i] > x_digits[i]:
                    break

        num = 0
        for x in x_digits:
            num *= 10
            num += x
        return int(num * sign)

def main():
    print(Solution().reverse((-1<<31) - 2))


main()

