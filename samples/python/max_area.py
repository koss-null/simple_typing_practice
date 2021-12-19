class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        lf, rg = 0, len(height) - 1
        max_vol = 0
        while lf < rg:
            vol = min(height[lf], height[rg]) * (rg - lf)
            if vol > max_vol:
                max_vol = vol
            if height[lf] > height[rg]:
                last_height = height[rg]
                rg -= 1
                while rg > lf and height[rg] < last_height:
                    rg -= 1
            else:
                last_height = height[lf]
                lf += 1
                while rg > lf and height[lf] < last_height:
                    lf += 1
        return max_vol
