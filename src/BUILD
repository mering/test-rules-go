load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library", "go_grpc_library")

proto_library(
    name = "msg_proto",
    srcs = ["msg.proto"],
)

go_proto_library(
    name = "msg_go_proto",
    deps = [":msg_proto"],
)

proto_library(
    name = "foo_proto",
    srcs = ["foo.proto"],
    deps = [":msg_proto"],
)

go_proto_library(
    name = "foo_go_proto",
    deps = [":foo_proto"],
)

proto_library(
    name = "svc_proto",
    srcs = ["svc.proto"],
    deps = [":msg_proto"],
)

go_grpc_library(
    name = "svc_go_grpc",
    protos = [":svc_proto"],
    deps = [":msg_go_proto"],
)

go_library(
    name = "svc_go",
    srcs = ["svc.go"],
    deps = [
        ":msg_go_proto",
        ":svc_go_grpc",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//credentials/insecure:go_default_library",
    ],
)

go_binary(
    name = "main",
    srcs = ["main.go"],
    deps = [
        ":msg_go_proto",
        ":foo_go_proto",
        ":svc_go",
    ],
)
