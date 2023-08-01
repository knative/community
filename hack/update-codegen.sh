#!/usr/bin/env bash

# Copyright 2018 The Knative Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

# We put this in a function just in case it's usable elsewhere
function make_aliases() {
  # We don't have a top-level `go.mod` because each tool should be small and
  # self-contained and we don't release them.
  ( cd mechanics/tools/gen-aliases; go build . )
  # Auto-generate OWNERS_ALIASES from peribolos configs
  ./mechanics/tools/gen-aliases/gen-aliases knative peribolos/knative.yaml peribolos/knative-OWNERS_ALIASES
  ./mechanics/tools/gen-aliases/gen-aliases knative-extensions peribolos/knative-extensions.yaml peribolos/knative-extensions-OWNERS_ALIASES
  # Clean up the tool so we don't try to check it in.
  rm ./mechanics/tools/gen-aliases/gen-aliases
}

make_aliases
