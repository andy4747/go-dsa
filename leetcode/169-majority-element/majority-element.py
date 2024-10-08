class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        n = len(nums)
        nMap = {}
        for v in nums:
            nMap[v] = nMap.get(v, 0) + 1
        
        n = n // 2
        for element, count in nMap.items():
            if count > n:
                return element
        return  -1
