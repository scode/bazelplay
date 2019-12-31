load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

java_proto_library(
    name = "person_java_proto",
    visibility = ["//visibility:public"],
    deps = [":person_proto"],
)

proto_library(
    name = "person_proto",
    srcs = ["person.proto"],
    deps = [
    ],
)

proto_library(
    name = "protodemo_proto",
    srcs = ["person.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "protodemo_go_proto",
    importpath = "github.com/scode/bazelplay/src/proto",
    proto = ":protodemo_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    embed = [":protodemo_go_proto"],
    importpath = "github.com/scode/bazelplay/src/proto",
    visibility = ["//visibility:public"],
)
