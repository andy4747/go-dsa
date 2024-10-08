class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        numsDict: dict[int, bool] = {}
        for i, v in enumerate(nums):
            if v in numsDict:
                return True
            numsDict[v] = True
        return False