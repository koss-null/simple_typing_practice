class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """

        num_cnt = 0
        places = [-1] * len(nums)
        for i in range(len(nums)):
            if nums[i]:
                places[i] = num_cnt
                num_cnt += 1

        for i in range(len(nums)):
            if places[i] != -1:
                nums[places[i]] = nums[i]

        for i in range(num_cnt, len(nums)):
            nums[i] = 0


def main():
    nums = [0,0,0,0,0,1,2,3,0,4,5,6,0,0,0,0]
    Solution().moveZeroes(nums)
    print(nums)


main()
