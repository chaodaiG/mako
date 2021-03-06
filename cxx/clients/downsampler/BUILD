#Google3 implementation of mako downsampler interface.
package(
    default_visibility = ["//visibility:public"],
)

licenses(["notice"])

cc_library(
    name = "standard_downsampler",
    srcs = ["standard_downsampler.cc"],
    hdrs = ["standard_downsampler.h"],
    deps = [
        ":metric_set",
        "//cxx/internal:proto_validation",
        "//cxx/internal/load/common:executor",
        "//cxx/spec:downsampler",
        "//cxx/spec:fileio",
        "//proto/internal:mako_internal_cc_proto",
        "//spec/proto:mako_cc_proto",
        "@com_google_absl//absl/algorithm:container",
        "@com_google_absl//absl/base:core_headers",
        "@com_google_absl//absl/memory",
        "@com_google_absl//absl/random",
        "@com_google_absl//absl/strings",
        "@com_google_absl//absl/synchronization",
        "@com_google_glog//:glog",
        "@com_google_protobuf//:protobuf",
    ],
)

cc_test(
    name = "standard_downsampler_test",
    size = "small",
    srcs = ["standard_downsampler_test.cc"],
    deps = [
        ":standard_downsampler",
        "//cxx/clients/fileio:memory_fileio",
        "//cxx/internal:filter_utils",
        "//cxx/internal:pgmath",
        "//spec/proto:mako_cc_proto",
        "@com_google_absl//absl/container:flat_hash_set",
        "@com_google_absl//absl/strings",
        "@com_google_glog//:glog",
        "@com_google_googletest//:gtest_main",
        "@com_google_protobuf//:protobuf",
    ],
)

cc_library(
    name = "metric_set",
    hdrs = ["metric_set.h"],
    visibility = ["//visibility:private"],
    deps = [
        "//spec/proto:mako_cc_proto",
        "@com_google_absl//absl/strings",
    ],
)

cc_test(
    name = "metric_set_test",
    size = "small",
    srcs = ["metric_set_test.cc"],
    deps = [
        ":metric_set",
        "@com_google_googletest//:gtest_main",
    ],
)
