"""Counting sort, as in https://en.wikipedia.org/wiki/Counting_sort."""


def sort(input, key=int):
    """Returns a sorted list of elements from input using counting sort.

    The input should be a sequence of integers.
    The key should be a function that maps an integer from the input into its
    sorting order.
    """
    # variables:
    #    input -- the array of items to be sorted; key(x) returns the key for item x
    #    n -- the length of the input
    n = len(input)
    #    k -- a number such that all keys are in the range 0..k-1
    k = max(input, default=0) + 1
    #    count -- an array of numbers, with indexes 0..k-1, initially all zero
    count = [0] * k
    #    output -- an array of items, with indexes 0..n-1
    output = [None] * n
    #    x -- an individual input item, used within the algorithm
    #    total, oldCount, i -- numbers used within the algorithm

    # calculate the histogram of key frequencies:
    for x in input:
        count[key(x)] += 1

    # calculate the starting index for each key:
    total = 0
    for i in range(k):   # i = 0, 1, ... k-1
        oldCount = count[i]
        count[i] = total
        total += oldCount

    # copy to output array, preserving order of inputs with equal keys:
    for x in input:
        output[count[key(x)]] = x
        count[key(x)] += 1

    return output
