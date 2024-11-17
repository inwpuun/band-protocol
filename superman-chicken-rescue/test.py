from problem2 import max_chickens_rescued

def test_max_chickens_rescued():
    n, k = 5, 5
    positions = [2, 5, 10, 12, 15]
    assert max_chickens_rescued(n, k, positions) == 2

    n, k = 6, 10
    positions = [1, 11, 30, 34, 35, 37]
    assert max_chickens_rescued(n, k, positions) == 4

    # additional test case
    n, k = 1, 1
    positions = [1]
    assert max_chickens_rescued(n, k, positions) == 1

    n, k = 5, 3
    positions = [1, 2, 4, 7, 8]
    assert max_chickens_rescued(n, k, positions) == 2

    n, k = 6, 100
    positions = [5, 15, 25, 35, 45, 55]
    assert max_chickens_rescued(n, k, positions) == 6

    n, k = 10, 1
    positions = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    assert max_chickens_rescued(n, k, positions) == 1

    n, k = 7, 5
    positions = [2, 5, 6, 7, 12, 15, 20]
    assert max_chickens_rescued(n, k, positions) == 3

    n, k = 6, 4
    positions = [1, 3, 4, 8, 10, 12]
    assert max_chickens_rescued(n, k, positions) == 3

    n, k = 50, 10
    positions = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
                11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
                21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
                31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
                41, 42, 43, 44, 45, 46, 47, 48, 49, 50]
    assert max_chickens_rescued(n, k, positions) == 10

    print("All tests passed.")

test_max_chickens_rescued()