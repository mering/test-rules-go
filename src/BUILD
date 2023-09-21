load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "bar_proto",
    srcs = ["bar.proto"],
)

go_proto_library(
    name = "bar_go_proto",
    protos = [":bar_proto"],
    importpath = "src/bar_go_proto",
)

proto_library(
    name = "foo_proto",
    srcs = ["foo.proto"],
    deps = [":bar_proto"],
)

go_proto_library(
    name = "foo_go_proto",
    protos = [":foo_proto"],
    deps = [":bar_go_proto"],
    importpath = "src/foo_go_proto",
)

go_binary(
    name = "main",
    srcs = ["main.go"],
    deps = [
        ":bar_go_proto",
        ":foo_go_proto",
    ],
    importpath = "src/main",
)
