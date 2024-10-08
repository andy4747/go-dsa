class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        nMap = {}
        for i, v in enumerate(nums):
            nMap[v] = nMap.get(v, 0) + 1
        
        majorityItemCount = -1
        majorityItem = 0
        for element, count in nMap.items():
            if count > majorityItemCount:
                majorityItemCount = count
                majorityItem = element
        return majorityItem