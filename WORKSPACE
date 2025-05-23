load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:jvm.bzl", "jvm_maven_import_external")

http_archive(
    name = "bazel_skylib",
    sha256 = "bc283cdfcd526a52c3201279cda4bc298652efa898b10b4db0837dc51652756f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.7.1/bazel-skylib-1.7.1.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.7.1/bazel-skylib-1.7.1.tar.gz",
    ],
)
load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")
bazel_skylib_workspace()

# proto_library, cc_proto_library, and java_proto_library rules implicitly
# depend on @com_google_protobuf for protoc and proto runtimes.
# This statement defines the @com_google_protobuf repo.
http_archive(
    name = "com_google_protobuf",
    sha256 = "3da9705d4659f0f03f09901926c821654197f7a9db84375c0a3049133f39350c", # Correct SHA256 for v31.0
    strip_prefix = "protobuf-31.0",
    urls = ["https://github.com/protocolbuffers/protobuf/releases/download/v31.0/protobuf-all-31.0.tar.gz"],
)


load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
protobuf_deps()

RULES_JVM_EXTERNAL_TAG = "6.7"
RULES_JVM_EXTERNAL_SHA = "a1e351607f04fed296ba33c4977d3fe2a615ed50df7896676b67aac993c53c18"

http_archive(
    name = "rules_jvm_external",
    strip_prefix = "rules_jvm_external-%s" % RULES_JVM_EXTERNAL_TAG,
    sha256 = RULES_JVM_EXTERNAL_SHA,
    url = "https://github.com/bazel-contrib/rules_jvm_external/releases/download/%s/rules_jvm_external-%s.tar.gz" % (RULES_JVM_EXTERNAL_TAG, RULES_JVM_EXTERNAL_TAG)
)

load("@rules_jvm_external//:repositories.bzl", "rules_jvm_external_deps")

rules_jvm_external_deps()

load("@rules_jvm_external//:setup.bzl", "rules_jvm_external_setup")

rules_jvm_external_setup()

maven_install(
    artifacts = [
        "junit:junit:4.13.2",
    ],
    repositories = [
        "https://repo1.maven.org/maven2",
    ],
    maven_install_json = "//:maven_install.json",
)

load("@maven//:defs.bzl", "pinned_maven_install")
pinned_maven_install()

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "130739704540caa14e77c54810b9f01d6d9ae897d53eedceb40fd6b75efc3c23",
    urls = [
        "https://mirror.bazel.build/github.com/bazel-contrib/rules_go/releases/download/v0.54.1/rules_go-v0.54.1.zip",
        "https://github.com/bazel-contrib/rules_go/releases/download/v0.54.1/rules_go-v0.54.1.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "01192f07755644762c89925c322ef70b9eba094a690919f548003786a1eb1575",
    urls = [
        "https://mirror.bazel.build/github.com/bazel-contrib/bazel-gazelle/releases/download/v0.43.0/bazel-gazelle-v0.43.0.tar.gz",
        "https://github.com/bazel-contrib/bazel-gazelle/releases/download/v0.43.0/bazel-gazelle-v0.43.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.24.3")

gazelle_dependencies()