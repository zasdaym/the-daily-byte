from valid_palindrome import valid_palindrome


def test_valid_palindrome():
    tests = [
        {
            "input": "",
            "want": True,
        },
        {
            "input": "tattarrattat",
            "want": True,
        },
        {
            "input": "coding",
            "want": False,
        },
    ]

    for test in tests:
        assert valid_palindrome(test["input"]) == test["want"]
