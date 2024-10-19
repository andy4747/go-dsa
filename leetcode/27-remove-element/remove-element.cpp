class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        int left = 0;
        int right = nums.size()-1;
        int len = nums.size();
        while(left<=right){
            if(nums[right]==val){
                nums[right] = 0;
                len--;
                right--;
            }

            if(nums[left]==val) {
                nums[left] = 0;
                std::swap(nums[left], nums[right]);
                right--;
                len--;
            }else {
                left++;
            }
        }
        return len;
    }
};