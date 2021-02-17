load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "hello_lib",
    srcs = ["hello.go"],
    importpath = "github.com/jamie-digital/rules_go_tests",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "hello_broken",
    embed = [":hello_lib"],
    visibility = ["//visibility:public"],
)

go_binary(
    name = "hello_working",
    embed = [":hello_lib"],
    gotags = ["gotags"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "hello_broken_test",
    size = "small",
    srcs = ["hello_test.go"],
    embed = [":hello_lib"],
)

go_test(
    name = "hello_working_test",
    size = "small",
    srcs = ["hello_test.go"],
    embed = [":hello_lib"],
    gotags = ["gotags"],
)
