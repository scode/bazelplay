# Generated in the bazel repo with:
#
#   bazel run //src/tools/generate_workspace -- --artifact=junit:junit:4.12
#

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:jvm.bzl", "jvm_maven_import_external")

# bazel-skylb 0.8.0 released 2019.03.20 (https://github.com/bazelbuild/bazel-skylib/releases/tag/0.8.0)
skylib_version = "0.8.0"
http_archive(
    name = "bazel_skylib",
    type = "tar.gz",
    url = "https://github.com/bazelbuild/bazel-skylib/releases/download/{}/bazel-skylib.{}.tar.gz".format (skylib_version, skylib_version),
    sha256 = "2ef429f5d7ce7111263289644d233707dba35e39696377ebab8b0bc701f7818e",
)

# proto_library, cc_proto_library, and java_proto_library rules implicitly
# depend on @com_google_protobuf for protoc and proto runtimes.
# This statement defines the @com_google_protobuf repo.
http_archive(
    name = "com_google_protobuf",
    sha256 = "f976a4cd3f1699b6d20c1e944ca1de6754777918320c719742e1674fcf247b7e",
    strip_prefix = "protobuf-3.7.1",
    urls = ["https://github.com/google/protobuf/archive/v3.7.1.zip"],
)

#### begin protobuf hack
# https://github.com/protocolbuffers/protobuf/blob/master/protobuf_deps.bzl
# https://github.com/bazelbuild/rules_cc/commit/e86b282e6fb72e6c508937f396c4e3b50cb9ad96
# https://github.com/protocolbuffers/protobuf/issues/5918)
#
# load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
http_archive(
    name = "net_zlib",
    build_file = "@com_google_protobuf//examples:third_party/zlib.BUILD",
    sha256 = "c3e5e9fdd5004dcb542feda5ee4f0ff0744628baf8ed2dd5d66f8ca1197cb1a1",
    strip_prefix = "zlib-1.2.11",
    urls = ["https://zlib.net/zlib-1.2.11.tar.gz"],
)

bind(
    name = "zlib",
    actual = "@net_zlib//:zlib"
)
#### end protobuf hack


# The following dependencies were calculated from:
# junit:junit:4.12


jvm_maven_import_external(
    name = "org_hamcrest_hamcrest_core",
    artifact = "org.hamcrest:hamcrest-core:1.3",
    artifact_sha256 = "66fdef91e9739348df7a096aa384a5685f4e875584cce89386a7a47251c4d8e9",
    server_urls = ["http://central.maven.org/maven2"],
    licenses = ["notice"],
)

jvm_maven_import_external(
    name = "junit_junit",
    artifact = "junit:junit:4.12",
    artifact_sha256 = "59721f0805e223d84b90677887d9ff567dc534d7c502ca903c0c2b17f05c116a",
    server_urls = ["http://central.maven.org/maven2"],
    licenses = ["notice"],
)
