import collections as cl

class Solution(object):
    def find_nice_idx(self, available_idxs, i, j):
        for idx in available_idxs:
            if idx != i and idx != j:
                return idx
        return None

    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        num2idx = cl.defaultdict(list)
        for i, num in enumerate(nums):
            num2idx[num].append(i)

        triples = set()
        all_triples = set()

        for i in range(0, len(nums)):
            for j in range(i + 1, len(nums)):
                available_idxs = num2idx[0 - nums[i] - nums[j]]
                idx = self.find_nice_idx(available_idxs, i, j)
                if idx is not None:
                    if (nums[i], nums[j], nums[idx]) not in all_triples:
                        triples.add((nums[i], nums[j], nums[idx]))

                        all_triples.add((nums[i], nums[j], nums[idx]))
                        all_triples.add((nums[j], nums[i], nums[idx]))
                        all_triples.add((nums[idx], nums[j], nums[i]))
                        all_triples.add((nums[idx], nums[i], nums[j]))
                        all_triples.add((nums[i], nums[idx], nums[j]))
                        all_triples.add((nums[j], nums[idx], nums[i]))
        return triples
