# Generated in the bazel repo with:
#
#   bazel run //src/tools/generate_workspace -- --artifact=junit:junit:4.12
#

# proto_library, cc_proto_library, and java_proto_library rules implicitly
# depend on @com_google_protobuf for protoc and proto runtimes.
# This statement defines the @com_google_protobuf repo.
http_archive(
    name = "com_google_protobuf",
    sha256 = "cef7f1b5a7c5fba672bec2a319246e8feba471f04dcebfe362d55930ee7c1c30",
    strip_prefix = "protobuf-3.5.0",
    urls = ["https://github.com/google/protobuf/archive/v3.5.0.zip"],
)

# The following dependencies were calculated from:
# junit:junit:4.12


# junit:junit:jar:4.12
maven_jar(
    name = "org_hamcrest_hamcrest_core",
    artifact = "org.hamcrest:hamcrest-core:1.3",
    sha1 = "42a25dc3219429f0e5d060061f71acb49bf010a0",
)

maven_jar(
    name = "junit_junit",
    artifact = "junit:junit:4.12",
)
