class Solution {
public:
    int romanToInt(string s) {
        unordered_map<char, int> rMap = {
            {'I', 1},   {'V', 5},   {'X', 10},   {'L', 50},
            {'C', 100}, {'D', 500}, {'M', 1000},
        };
        int result = 0;
        for (int i = 0; i < s.size() - 1; i++) {
            if (rMap[s[i]] >= rMap[s[i + 1]]) {
                result += rMap[s[i]];
            } else {
                result -= rMap[s[i]];
            }
        }
        return result + rMap[s[s.size() - 1]];
    }
};