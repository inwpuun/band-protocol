from problem1 import boss_baby_revenge

def test_boss_baby_revenge():
    assert boss_baby_revenge("SRSSRRR") == "Good boy"
    assert boss_baby_revenge("RSSRR") == "Bad boy"
    assert boss_baby_revenge("SSSRRRRS") == "Bad boy"
    assert boss_baby_revenge("SRRSSR") == "Bad boy"
    assert boss_baby_revenge("SSRSRR") == "Good boy"

    # additional test case
    assert boss_baby_revenge("SSRR") == "Good boy"
    assert boss_baby_revenge("SSSS") == "Bad boy"
    assert boss_baby_revenge("SR") == "Good boy"
    assert boss_baby_revenge("SSSSSRRRRR") == "Good boy"
    assert boss_baby_revenge("SRRRRR") == "Good boy"
    assert boss_baby_revenge("S") == "Bad boy"
    assert boss_baby_revenge("R") == "Bad boy"
    assert boss_baby_revenge("SSRSRSRR") == "Good boy"
    assert boss_baby_revenge("SSRSRSRS") == "Bad boy"

    print("All tests passed.")

test_boss_baby_revenge()