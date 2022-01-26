from vacuum_cleaner_route import is_complete_route


def test_is_complete_route():
    tests = [
        {
            "moves": "",
            "want": True,
        },
        {
            "moves": "LR",
            "want": True,
        },
        {
            "moves": "RUULLDRD",
            "want": True,
        },
        {
            "moves": "URURD",
            "want": False,
        },
    ]

    for test in tests:
        assert is_complete_route(test["moves"]) == test["want"]
