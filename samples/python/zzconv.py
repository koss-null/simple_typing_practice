class Solution(object):
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """

        used = set()
        converted = []
        step = max(2 * numRows - 2, 1)
        for i in range(0, numRows):
            # print(f"i = {i}")
            j = 0
            while i + step * j < len(s):
                # print(f"j = {j}")
                if i + step * j not in used:
                    # print("adding {}".format(i + step * j))
                    converted.append(s[i + step * j])
                    used.add(i + step * j)

                small_step = 2 * numRows - 2 * (i + 1)
                # print(f"small step: {small_step}")
                small_step_idx = i + step * j + small_step
                if small_step_idx >= 0 and small_step_idx < len(s) and small_step_idx not in used:
                    # print("adding1 {}".format(small_step_idx))
                    converted.append(s[small_step_idx])
                    used.add(small_step_idx)

                j += 1
        return "".join(converted)

def main():
    s = Solution().convert("P", 1)
    print(s)
    print('PAHNAPLSIIGYIR')
    print(len(s))


main()


