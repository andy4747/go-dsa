class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        sDict: dict[str, int] = {}
        tDict: dict[str, int] = {}

        for i in range(len(s)):
            sDict[s[i]] = sDict.get(s[i], 0) + 1
            tDict[t[i]] = tDict.get(t[i], 0) + 1
        
        for char, count in sDict.items():
            if tDict.get(char) != count:
                return False
        return True