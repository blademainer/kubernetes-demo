#!/usr/bin/env bash

# Copyright 2015 The Kubernetes Authors.
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

KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

# NOTE: All output from this script needs to be copied back to the calling
# source tree.  This is managed in kube::build::copy_output in build/common.sh.
# If the output set is changed update that function.

APIROOTS=${APIROOTS:-$(git grep --files-with-matches -e '// +k8s:protobuf-gen=package' pkg | \
	xargs -n 1 dirname | \
	sed 's,^,k8s.io/kubernetes/,;s,k8s.io/kubernetes/staging/src/,,' | \
	sort | uniq
)}

hack/update-generated-protobuf-dockerized.sh "${APIROOTS}" "$@"

# ex: ts=2 sw=2 et filetype=sh
