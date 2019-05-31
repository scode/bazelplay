genrule(
    name = "hello",
    outs = ["hello_world.txt"],
    cmd = "echo Hello World > $@",
)


# Generated in the bazel repo with:
#
#   bazel run //src/tools/generate_workspace -- --artifact=junit:junit:4.12
#

# The following dependencies were calculated from:
# junit:junit:4.12

java_library(
    name = "org_hamcrest_hamcrest_core",
    visibility = ["//visibility:public"],
    exports = [
        "@org_hamcrest_hamcrest_core//jar",
    ],
)

java_library(
    name = "junit_junit",
    visibility = ["//visibility:public"],
    exports = [
        "@junit_junit//jar",
        "@org_hamcrest_hamcrest_core//jar",
    ],
)

# Per https://github.com/bazelbuild/bazel/issues/4815#issuecomment-460777113 until bug is fixed.
py_runtime(
    name = "pyruntime",
    visibility = ["//visibility:public"],
    interpreter_path = select({
        "@bazel_tools//tools/python:PY2": "/usr/bin/python2",
        "@bazel_tools//tools/python:PY3": "/usr/bin/python3",
    }),
    files = [],
)
