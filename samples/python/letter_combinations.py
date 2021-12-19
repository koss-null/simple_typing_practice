class Solution(object):
    def __init__(self):
        self.phone_keyboard = {
                "2": "abc",
                "3": "def",
                "4": "ghi",
                "5": "jkl",
                "6": "mno",
                "7": "pqrs",
                "8": "tuv",
                "9": "wxyz"
        }

    def get_combinations(self, digit, combinations):
        ret = list()
        for combination in combinations if combinations else [""]:
            for char in self.phone_keyboard[digit]:
                ret.append(combination + char)
        return ret

    def letterCombinations(self, digits):
        """
        :type digits: str
        :rtype: List[str]
        """
        combinations = list()
        for digit in digits:
            combinations = self.get_combinations(digit, combinations)
        return combinations


