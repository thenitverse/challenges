def max_area(height):
    left = 0
    right = len(height) - 1
    max_water = 0

    while left < right:
        h = min(height[left], height[right])
        width = right - left
        max_water = max(max_water, h * width)

        if height[left] < height[right]:
            left += 1
        else:
            right -= 1

    return max_water


def run_tests():
    test_cases = [
        ([1, 8, 6, 2, 5, 4, 8, 3, 7], 49),
        ([1, 1], 1),
        ([0, 0, 0, 0], 0),
        ([4, 3, 2, 1, 4], 16),
        ([2, 3, 10, 5, 7, 8, 9], 36),
        ([1, 2, 3, 4, 6, 7, 8], 12),  # The trace we just walked through
    ]

    for idx, (heights, expected) in enumerate(test_cases, 1):
        result = max_area(heights)
        status = "PASS" if result == expected else "FAIL"
        print(f"Test {idx}: {status} | Heights: {heights} | Expected: {expected} | Got: {result}")


if __name__ == "__main__":
    run_tests()