from longest_common_prefix import longest_common_prefix


def test_longest_common_prefix():
    tests = [
        {
            "name": "single word",
            "words": ["Pneumonoultramicroscopicsilicovolcanoconiosis"],
            "want": "Pneumonoultramicroscopicsilicovolcanoconiosis"
        },
        {
            "name": "common prefix exists",
            "words": ["colorado", "color", "cold"],
            "want": "col"
        },
        {
            "name": "duplicate words",
            "words": ["color", "color", "color"],
            "want": "color"
        },
        {
            "name": "no common prefix",
            "words": ["a", "b", "c"],
            "want": ""
        },
    ]

    for test in tests:
        assert longest_common_prefix(test["words"]) == test["want"]
