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
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Mako UTest Analyzer.
//
// For detailed information on usage, see http://go/mako-utest.
//
// For more information about the interface see
// mako/spec/py/analyzer.h
#ifndef CXX_CLIENTS_ANALYZERS_UTEST_ANALYZER_H_
#define CXX_CLIENTS_ANALYZERS_UTEST_ANALYZER_H_

#include <memory>
#include <string>
#include <vector>

#include "spec/proto/mako.pb.h"
#include "cxx/spec/analyzer.h"
#include "proto/clients/analyzers/utest_analyzer.pb.h"

namespace mako {
namespace utest_analyzer {

class StatsCalculator;

class Analyzer : public mako::Analyzer {
 public:
  // Constructor, takes configuration input.
  explicit Analyzer(const mako::utest_analyzer::UTestAnalyzerInput& config)
      : config_(config) {}

  ~Analyzer() override {}

  // Historical information the analyzer requires, see ifc for more info.
  bool ConstructHistoricQuery(
      const mako::AnalyzerHistoricQueryInput& input,
      mako::AnalyzerHistoricQueryOutput* output) override;

  // Return analyzer type, see ifc for more info
  std::string analyzer_type() override { return "UTest"; }

  // Return analyzer name, see ifc for more info.
  // If UTestAnalyzerInput.name is not set, empty std::string will be returned.
  std::string analyzer_name() override { return config_.name(); }

 private:
  // Run the analysis, see ifc for more info.
  bool DoAnalyze(const mako::AnalyzerInput& analyzer_input,
                 mako::AnalyzerOutput* analyzer_output) override;

  std::string ValidateUTestAnalyzerInput() const;

  std::string AnalyzeUTestConfig(const std::string& config_name,
                                 const StatsCalculator& s_calc,
                                 const UTestConfig& config,
                                 UTestConfigResult* result);

  bool SetAnalyzerOutputPassing(AnalyzerOutput* output,
                                UTestAnalyzerOutput* custom_output);

  bool SetAnalyzerOutputWithError(AnalyzerOutput* output,
                                  UTestAnalyzerOutput* custom_output,
                                  const std::string& err_msg);

  bool SetAnalyzerOutputWithRegression(
      AnalyzerOutput* output, UTestAnalyzerOutput* custom_output,
      const std::vector<std::string>& failed_config_names);

  UTestAnalyzerInput config_;
};

}  // namespace utest_analyzer
}  // namespace mako

#endif  // CXX_CLIENTS_ANALYZERS_UTEST_ANALYZER_H_
