// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// see the license for the specific language governing permissions and
// limitations under the license.

// Perform validation on REQUIRED fields in mako protocol buffers.
//
// Functions below perform validation on the supplied protobuffers. An error
// string is returned, empty means successful.
//
// NOTE: Some validation is still left to the mako server (eg. valid
// characters in labels, etc). The purpose of this validation library is to
// ensure that protobufs can safely be processed.
#ifndef CXX_INTERNAL_PROTO_VALIDATION_H_
#define CXX_INTERNAL_PROTO_VALIDATION_H_

#include <string>

#include "spec/proto/mako.pb.h"

// TODO(b/149356416) Remove the optnone attribute and delete this macro
#ifndef MAKO_INTERNAL_OPTNONE
// SWIG parser cannot understand the attribute, so #define it away.
// This only impacts the SWIG parser, not the code that is built into the C++
// library, so it does not reintroduce the bug we're working around here.
#ifdef SWIG
#define MAKO_INTERNAL_OPTNONE
#else
#define MAKO_INTERNAL_OPTNONE __attribute__((optnone))
#endif  // SWIG
#endif  // MAKO_INTERNAL_OPTNONE

namespace mako {
namespace internal {

// Validate that all REQUIRED AggregatorInput fields are set. Returns error
// message as string or empty if successful.
std::string ValidateAggregatorInput(const mako::AggregatorInput& input)
    MAKO_INTERNAL_OPTNONE;

// Validate that all REQUIRED RunInfo fields are set. Returns error
// message as string or empty if successful.
std::string ValidateRunInfo(const mako::RunInfo& input)
    MAKO_INTERNAL_OPTNONE;

// Validate that all REQUIRED BenchmarkInfo fields are set. Returns error
// message as string or empty if successful.
std::string ValidateBenchmarkInfo(const mako::BenchmarkInfo& input)
    MAKO_INTERNAL_OPTNONE;

// Validate that all REQUIRED SampleFile fields are set. Returns error
// message as string or empty if successful.
std::string ValidateSampleFile(const mako::SampleFile& input)
    MAKO_INTERNAL_OPTNONE;

// Validate that all REQUIRED DownsamplerInput fields are set. Returns error
// message as string or empty if successful.
std::string ValidateDownsamplerInput(const mako::DownsamplerInput& input)
    MAKO_INTERNAL_OPTNONE;

// Clears the aux_data field so it will not take up space on the
// server.
void StripAuxData(mako::SamplePoint* point);
// Overload included to allow template functions to compile.
void StripAuxData(mako::SampleError* error);

// Validate that all REQUIRED BenchmarkInfo fields are set for Benchmark
// creation. Returns error message as string or empty if successful.
std::string ValidateBenchmarkInfoCreationRequest(
    const mako::BenchmarkInfo& input) MAKO_INTERNAL_OPTNONE;

// Validate that all REQUIRED RunInfo fields are set for RunInfo creation.
// Returns error message as string or empty if successful.
std::string ValidateRunInfoCreationRequest(const mako::RunInfo& input)
    MAKO_INTERNAL_OPTNONE;

// Validate that all REQUIRED SampleBatch fields are set for SampleBatch
// creation.  Returns error message as string or empty if successful.
std::string ValidateSampleBatchCreationRequest(
    const mako::SampleBatch& input) MAKO_INTERNAL_OPTNONE;

// Validate that all REQUIRED ProjectInfo fields are set. Returns error message
// as string or empty if successful.
std::string ValidateProjectInfo(const mako::ProjectInfo& input)
    MAKO_INTERNAL_OPTNONE;

}  // namespace internal
}  // namespace mako

#ifdef MAKO_INTERNAL_OPTNONE
#undef MAKO_INTERNAL_OPTNONE
#endif

#endif  // CXX_INTERNAL_PROTO_VALIDATION_H_
