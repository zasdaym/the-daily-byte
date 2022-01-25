from reverse_string import reverse


def test_reverse():
    tests = [
        {
            "input": "",
            "want": "",
        },
        {
            "input": "coding",
            "want": "gnidoc",
        },
        {
            "input": "level",
            "want": "level",
        },
    ]

    for test in tests:
        assert reverse(test["input"]) == test["want"]
