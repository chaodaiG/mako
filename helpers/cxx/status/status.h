// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#ifndef HELPERS_CXX_STATUS_H_
#define HELPERS_CXX_STATUS_H_

#include "google/cloud/status.h"
#include "absl/strings/string_view.h"

namespace mako {
namespace helpers {

using ::google::cloud::StatusCode;
using ::google::cloud::StatusCodeToString;
using ::google::cloud::Status;
// Returns an OK status, equivalent to a default constructed instance.
Status OkStatus();

std::string StatusToString(const Status& status);

bool HasErrorCode(const Status& status, StatusCode code);

Status Annotate(const Status& status, absl::string_view message);

}  // namespace helpers
}  // namespace mako

#endif  // HELPERS_CXX_STATUS_H_
