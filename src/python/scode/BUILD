py_binary(
    name = "hello",
    srcs = ["hello.py"],
    python_version = "PY3",
)

py_test(
    name = "hello_test",
    srcs = ["hello_test.py"],
    python_version = "PY3",
    deps = [
        ":hello",
    ]
)
