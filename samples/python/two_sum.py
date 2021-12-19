class Solution(object):
    def search(self, nums, target, current_idx):
        lf, rg = 0, len(nums) - 1
        while rg >= lf:
            md = int((lf + rg) / 2)
            if nums[md][0] < target:
                lf = md + 1
            elif nums[md][0] > target:
                rg = md - 1
            else:
                if current_idx == nums[md][1]:
                    if md + 1 < len(nums) and nums[md+1][0] == target:
                        return True, nums[md+1][1]
                    else:
                        return False, -1
                return True, nums[md][1]
        return False, -1

    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        nums = [(num, i) for i, num in enumerate(nums)]
        sorted_nums = sorted(nums, key=lambda x: x[0])
        for i, num in enumerate(sorted_nums):
            opposite = target - num[0]
            print(num[0], opposite)
            found, idx = self.search(sorted_nums, opposite, num[1])
            if found:
                return [num[1], idx]


def main():
    print(Solution().twoSum([3,2,4], 6))

main()
