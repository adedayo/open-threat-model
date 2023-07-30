#!/bin/bash
set -e

readonly service="$1"
readonly output_dir="$2"
readonly package="$3"


version=$(sed -n 's/.*version:\s*\(.*\).*/\1/p' "openapi/$service.yml")
echo -n $version > "$output_dir/api_version.txt"

docker run --rm \
  -v ${PWD}:/local \
  openapitools/openapi-generator-cli:latest-release generate \
  -i /local/openapi/otm-openapi.yml \
  -g go --package-name otm --type-mappings integer=int \
  --additional-properties withGoAdditionalProperties=true \
  -o /local/out/go

# copy generated go files to the model directory
  mkdir -p pkg/model
  mv out/go/model_*.go out/go/utils.go pkg/model

# add yaml struct tags too so we can control yaml (de)serialisation
find pkg/model -name "*.go" -exec sed -i.bak -E 's/json:"([^"]*)"/json:"\1" yaml:"\1"/g' {} \; -exec rm {}.bak \;

#cleanup
rm -rf out
go mod tidy