class Solution {
public:
    bool isAnagram(string s, string t) {
        int count[26] = {0};

        // increment the count of chars from s
        for (char i : s) {
            count[i - 'a']++;
        }

        // decrement the count of chars from t
        for (char i : t) {
            count[i - 'a']--;
        }

        for (int i : count) {
            if (i != 0) {
                return false;
            }
        }
        return true;
    }
};