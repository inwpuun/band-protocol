# time complexity: O(n)
# space complexity: O(1)
def max_chickens_rescued(n, k, positions):
    result = 0
    start_point = 0

    for end_point in range(n):
        # While the roof's range is exceeded, move the start pointer
        while positions[end_point] - positions[start_point] >= k:
            start_point += 1

        # Update the result
        result = max(result, end_point - start_point + 1)

    return result