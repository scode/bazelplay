genrule(
    name = "hello",
    outs = ["hello_world.txt"],
    cmd = "echo Hello World > $@",
)

java_library(
    name = "junit_junit",
    visibility = ["//visibility:public"],
    exports = [
        "@maven//:junit_junit",
    ],
)
