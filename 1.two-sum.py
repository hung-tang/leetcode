#
# @lc app=leetcode id=1 lang=python3
#
# [1] Two Sum
#
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        lookup = {}
        for i in range(len(nums)):
            ans = target - nums[i]
            if ans in lookup:
                return [lookup[ans], i]
            lookup[nums[i]] = i
        

