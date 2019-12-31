# Generated in the bazel repo with:
#
#   bazel run //src/tools/generate_workspace -- --artifact=junit:junit:4.12
#

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:jvm.bzl", "jvm_maven_import_external")

http_archive(
    name = "bazel_skylib",
    type = "tar.gz",
    url = "https://github.com/bazelbuild/bazel-skylib/releases/download/0.9.0/bazel_skylib-0.9.0.tar.gz",
    sha256 = "1dde365491125a3db70731e25658dfdd3bc5dbdfd11b840b3e987ecf043c7ca0",
)

# proto_library, cc_proto_library, and java_proto_library rules implicitly
# depend on @com_google_protobuf for protoc and proto runtimes.
# This statement defines the @com_google_protobuf repo.
http_archive(
    name = "com_google_protobuf",
    sha256 = "e4f8bedb19a93d0dccc359a126f51158282e0b24d92e0cad9c76a9699698268d",
    strip_prefix = "protobuf-3.11.2",
    urls = ["https://github.com/google/protobuf/archive/v3.11.2.zip"],
)


http_archive(
   name = "com_google_protobuf",
   sha256 = "e4f8bedb19a93d0dccc359a126f51158282e0b24d92e0cad9c76a9699698268d",
   strip_prefix = "protobuf-3.11.2",
   urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.11.2.zip"],
)


load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
protobuf_deps()

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
