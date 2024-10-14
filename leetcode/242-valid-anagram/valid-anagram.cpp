class Solution {
public:
    bool isAnagram(string s, string t) {
        if (s.size() != t.size()) {
            return false;
        }
        unordered_map<char, int> sMap;
        unordered_map<char, int> tMap;
        for (int i = 0; i < s.size(); i++) {
            sMap[s[i]]++;
            tMap[t[i]]++;
        }
        for (const auto& [chr, count] : sMap) {
            if (tMap[chr] != count) {
                return false;
            }
        }
        return true;
    }
};