from collections import namedtuple


Interval = namedtuple("Interval", "start end")

class Solution(object):
    def findPlace(self, nums, val, saved_intervals):
        if val in saved_intervals:
            return saved_intervals[val]
        lf, rg = 0, len(nums) - 1
        while rg >= lf:
            md = (lf + rg) // 2
            if nums[md] > val:
                rg = md - 1
            elif nums[md] < val:
                lf = md + 1
            else:
                left_step, left_border = md, md
                right_step, right_border = len(nums) - md, md
                #left
                while left_step:
                    if left_border - left_step > -1 and nums[left_border - left_step] == nums[md]:
                        left_border -= left_step
                        continue
                    left_step = left_step // 2
                #right
                while right_step:
                    if right_border + right_step < len(nums) and nums[right_border + right_step] == nums[md]:
                        right_border += right_step
                        continue
                    right_step = right_step // 2
                return Interval(left_border, right_border + 1)

        return Interval(lf, lf)

    def findMedian(self, nums1, nums2, place):
        lf, rg = 0, len(nums1) - 1
        # start - how many values are less than md, end - how many values are greater than md
        saved_intervals1, saved_intervals2 = dict(), dict()
        while rg >= lf:
            md = (lf + rg) // 2
            nums1_res = self.findPlace(nums1, nums1[md], saved_intervals1)
            nums2_res = self.findPlace(nums2, nums1[md], saved_intervals2)
            saved_intervals1[nums1[md]] = nums1_res
            saved_intervals2[nums1[md]] = nums2_res

            if nums1_res.start + nums2_res.start <= place < nums1_res.end + nums2_res.end:
                return True, md

            if nums1_res.start + nums2_res.start < place:
                lf = md + 1
            elif nums1_res.end + nums2_res.end > place:
                rg = md - 1
            else:
                return True, md
        return False, -1


    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """

        if (len(nums1) + len(nums2)) % 2 != 0:
            found, p = self.findMedian(nums1, nums2, (len(nums1) + len(nums2)) // 2)
            if found:
                return nums1[p]
            return nums2[self.findMedian(nums2, nums1, (len(nums1) + len(nums2)) // 2)[1]]

        found, p = self.findMedian(nums1, nums2, (len(nums1) + len(nums2)) // 2 - 1)
        if found:
            first_num = nums1[p]
        else:
            first_num = nums2[self.findMedian(nums2, nums1, (len(nums1) + len(nums2)) // 2 - 1)[1]]

        found, p = self.findMedian(nums1, nums2, (len(nums1) + len(nums2)) // 2)
        if found:
            second_num = nums1[p]
        else:
            second_num = nums2[self.findMedian(nums2, nums1, (len(nums1) + len(nums2)) // 2)[1]]

        return (first_num + second_num) / 2.0

print(Solution().findMedianSortedArrays([1,3], [2, 4]))
