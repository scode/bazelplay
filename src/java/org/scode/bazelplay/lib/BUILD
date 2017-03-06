java_library(
        name = "greeter",
        srcs = ["Greeter.java"],
        visibility = ["//src/java/org/scode/bazelplay:__pkg__"]
)

java_library(
        name = "greeter-test",
        srcs = ["GreeterTest.java"],
        deps = [
                ":greeter",
                "//:junit_junit",
        ]
)

java_test(
        name = "GreeterTest",
        size = "small",
        runtime_deps = [
                ":greeter-test",
        ]
)