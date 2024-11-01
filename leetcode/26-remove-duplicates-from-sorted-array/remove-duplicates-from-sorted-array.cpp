class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int unique = 1;
        int n = nums.size();
        if(n==1) {
            return 1;
        }
        for(int i=1;i<n;i++) {
            if(nums[i] != nums[i-1]) {
                nums[unique] = nums[i];
                unique++;
            }
        }
        return unique;
    }
};
