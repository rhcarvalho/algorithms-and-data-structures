import pytest

from counting import sort


@pytest.mark.parametrize("input,want", [
    ([], []),
    ([1], [1]),
    ([1, 2], [1, 2]),
    ([3, 1, 2], [1, 2, 3]),
    ([4, 3, 1, 2], [1, 2, 3, 4]),
])
def test_counting_sort(input, want):
    assert sort(input) == want
