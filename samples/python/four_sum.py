class Solution(object):
    def unique(self, a, b, c, d):
        return (
            a != b and a != c and a != d and
            b != c and b != d and
            c != d
        )

    def fourSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        nums = sorted(nums)
        quadrupels = set()

        for i in range(0, len(nums)):
            for j in range(i + 1, len(nums)):
                lf, rg = j + 1, len(nums) - 1
                tg = target - nums[i] - nums[j]
                while(lf < rg):
                    if (tg - nums[lf] - nums[rg]) > 0:
                        lf += 1
                        continue
                    if (tg - nums[lf] - nums[rg]) < 0:
                        rg -= 1
                        continue
                    lf_tmp, rg_tmp = nums[lf], nums[rg]
                    quadrupels.add((nums[i], nums[j], nums[lf], nums[rg]))
                    while (lf < rg and (nums[lf] == lf_tmp or nums[rg] == rg_tmp)):
                        lf += 1 if nums[lf] == lf_tmp else 0
                        rg -= 1 if nums[rg] == rg_tmp else 0
                    continue

        return list(quadrupels)

print(Solution().fourSum([2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2],8))
