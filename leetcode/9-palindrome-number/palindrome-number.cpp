class Solution {
public:
    bool isPalindrome(int x) {
        if(x<0) {
            return false;
        }
        int old = x;
        long long result = 0;
        while (x!=0) {
            int temp = x % 10;
            result = result * 10 + temp;
            x = x / 10;
        }
        if (old == int(result)) {
            return true;
        }
        return false;
    }
};