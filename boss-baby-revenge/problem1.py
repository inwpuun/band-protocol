# time complexity: O(n)
# space complexity: O(1)
def boss_baby_revenge(s: str) -> str:
    # if boss baby initiate shots at the neighborhood kids first then he is bad boy
    if s[0] == "R":
        return "Bad boy"

    counter = 0
    for i in range(len(s)):
        if s[i] == "S":
            counter += 1
        else:
            # do not care if boss baby shot more than the kids
            if counter == 0 :
                continue
            counter -= 1

    # all the shots have been revenged
    if counter == 0:
        return "Good boy"

    return "Bad boy"
