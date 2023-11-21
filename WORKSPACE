workspace(name = "test_rules_go")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

local_repository(
    name = "io_bazel_rules_go",
    path = "./rules_go/",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
go_rules_dependencies()
go_register_toolchains(version = "1.21.0")

http_archive(
    name = "com_google_protobuf",
    strip_prefix = "protobuf-25.1",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v25.1.zip"],
)
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
protobuf_deps()

# CC toolchain
BAZEL_TOOLCHAIN_TAG = "0.10.3"
BAZEL_TOOLCHAIN_SHA = "c2a58fdd8420cca3645843b1be7b18a2d5df388d192d50c238ae3edd9b693011"
http_archive(
    name = "com_grail_bazel_toolchain",
    sha256 = BAZEL_TOOLCHAIN_SHA,
    strip_prefix = "bazel-toolchain-{tag}".format(tag = BAZEL_TOOLCHAIN_TAG),
    canonical_id = BAZEL_TOOLCHAIN_TAG,
    url = "https://github.com/grailbio/bazel-toolchain/archive/{tag}.tar.gz".format(tag = BAZEL_TOOLCHAIN_TAG),
)

load("@com_grail_bazel_toolchain//toolchain:deps.bzl", "bazel_toolchain_dependencies")

bazel_toolchain_dependencies()

load("@com_grail_bazel_toolchain//toolchain:rules.bzl", "llvm_toolchain")

llvm_toolchain(
    name = "llvm_toolchain",
    llvm_version = "16.0.4",
)

load("@llvm_toolchain//:toolchains.bzl", "llvm_register_toolchains")

llvm_register_toolchains()

http_archive(
    name = "bazel_gazelle",
    sha256 = "d3fa66a39028e97d76f9e2db8f1b0c11c099e8e01bf363a923074784e451f809",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.33.0/bazel-gazelle-v0.33.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.33.0/bazel-gazelle-v0.33.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:SXUpjxeVF3FKrTYQI4f4KvbGD5u2xccdYdurwowix5I=",
    version = "v1.58.2",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:ugBLEUaxABaB5AJqW9enI0ACdci2RUd4eP51NTBvuJ8=",
    version = "v0.15.0",
)

go_repository(
    name = "org_golang_google_genproto_googleapis_rpc",
    importpath = "google.golang.org/genproto/googleapis/rpc",
    sum = "h1:N3bU/SQDCDyD6R528GJ/PwW9KjYcJA3dgyH+MovAkIM=",
    version = "v0.0.0-20230920204549-e6e6cdab5c13",
)
