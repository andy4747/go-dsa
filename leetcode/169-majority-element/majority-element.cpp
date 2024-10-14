class Solution {
public:
    int majorityElement(vector<int>& nums) {
        unordered_map<int, int> count;
        int n = nums.size() - 1;
        int check = n / 2;
        for (int i : nums) {
            count[i]++;
            if (count[i] > check) {
                return i;
            }
        }
        return -1;
    }
};