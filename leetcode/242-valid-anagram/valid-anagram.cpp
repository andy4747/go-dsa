class Solution {
public:
    bool isAnagram(string s, string t) {
        if(s.size()!=t.size()){
            return false;
        }
        int count[26] = {0};
        for(char i : s) {
            count[i - 'a']++;
        }
        for (char i : t) {
            count[i - 'a']--;
        }
        for (int i : count) {
            if(i!=0) {
                return false;
            }
        }
        return true;
    }
};