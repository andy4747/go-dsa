class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        numsDict: dict[int, int] = {}
        for i, v in enumerate(nums):
            more = target - v
            if more in numsDict:
                return [numsDict[more], i]
            numsDict[v] = i
        return []
