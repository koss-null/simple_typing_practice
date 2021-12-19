class Solution(object):
    def find_closest_sum(self, nums, target, i, j):
        lf, rg = 0, len(nums) - 1
        while lf < rg:
            md = (lf + rg) // 2
            md_diff = (target - nums[i] - nums[j] - nums[md])
            if md_diff > 0:
                lf = md + 1
            else:
                rg = md - 1

        lf = max(0, md - 1)
        rg = min(len(nums) - 1, md + 1)

        min_sum = 1000 * 20
        sm = 0
        # print(nums, i, j, md)
        if md != i and md != j and abs(md_diff) < abs(min_sum):
            min_sum = md_diff
            sm = nums[i] + nums[j] + nums[md]
            # print("choose md {} ms is {}".format(md, sm))
        if lf != i and lf != j and abs(target - nums[i] - nums[j] - nums[lf]) < abs(min_sum):
            min_sum = (target - nums[i] - nums[j] - nums[lf])
            sm = nums[i] + nums[j] + nums[lf]
            # print("choose lf {} ms is {}".format(lf, sm))
        if rg != i and rg != j and abs(target - nums[i] - nums[j] - nums[rg]) < abs(min_sum):
            min_sum = (target - nums[i] - nums[j] - nums[rg])
            sm = nums[i] + nums[j] + nums[rg]
            # print("choose rg {} ms is {}".format(rg, sm))
        return sm

    def threeSumClosest(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        min_sum = 1000 * 20
        nums = sorted(nums)
        for i in range(0, len(nums)):
            for j in range(i + 1, len(nums)):
                # print("i {}, j {}".format(i, j))
                closest_sum = self.find_closest_sum(nums, target, i, j)
                if abs(min_sum) > abs(target - closest_sum):
                    # print("sum have been chousen as the best")
                    min_sum = target - closest_sum
        return target - min_sum


print(Solution().threeSumClosest([-1, 2, 1, -4], 1))
