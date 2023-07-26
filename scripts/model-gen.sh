#!/bin/bash
set -e

readonly service="$1"
readonly output_dir="$2"
readonly package="$3"


version=$(sed -n 's/.*version:\s*\(.*\).*/\1/p' "openapi/$service.yml")
echo -n $version > "$output_dir/api_version.txt"
# oapi-codegen -generate types -o "$output_dir/model.gen.go" -package "$package" "openapi/$service.yml"
# oapi-codegen -generate server -o "$output_dir/api.gen.go" -package "$package" "openapi/$service.yml"


docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i /local/openapi/otm-openapi.yml \
  -g go --type-mappings integer=int --additional-properties withGoAdditionalProperties=true \
  -o /local/out/go --package-name otm

  mkdir -p pkg/model
  mv out/go/model_*.go out/go/utils.go pkg/model

# add yaml struct tags too
#   find pkg/model -name "*.go" -exec sed -i 's/json:"\([^"]*\)"/json:"\1" yaml:"\1"/g' {} +
# find pkg/model -name "*.go" -exec sed -E -i 's/(json:"[^"]*")/\1 yaml:"\1"/g' {} +
find pkg/model -name "*.go" -exec sed -i.bak -E 's/json:"([^"]*)"/json:"\1" yaml:"\1"/g' {} \; -exec rm {}.bak \;


 rm -rf out

  go mod tidy