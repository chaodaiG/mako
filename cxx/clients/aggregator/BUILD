package(default_visibility = ["//visibility:public"])

licenses(["notice"])

cc_library(
    name = "standard_aggregator",
    srcs = ["standard_aggregator.cc"],
    hdrs = [
        "standard_aggregator.h",
        "threadsafe_running_stats.h",
    ],
    deps = [
        "//cxx/internal:pgmath",
        "//cxx/internal:proto_validation",
        "//cxx/internal/load/common:executor",
        "//cxx/spec:aggregator",
        "//cxx/spec:fileio",
        "//spec/proto:mako_cc_proto",
        "@com_google_absl//absl/memory",
        "@com_google_absl//absl/strings",
        "@com_google_absl//absl/synchronization",
        "@com_google_absl//absl/time",
        "@com_google_glog//:glog",
        "@com_google_protobuf//:protobuf",
    ],
)

cc_test(
    name = "standard_aggregator_test",
    size = "small",
    srcs = [
        "standard_aggregator_test.cc",
    ],
    deps = [
        ":standard_aggregator",
        "//cxx/clients/fileio:memory_fileio",
        "//spec/proto:mako_cc_proto",
        "@com_google_absl//absl/memory",
        "@com_google_absl//absl/strings",
        "@com_google_absl//absl/strings:str_format",
        "@com_google_googletest//:gtest_main",
    ],
)
