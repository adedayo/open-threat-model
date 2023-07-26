
.PHONY: otm
otm: openapi_otm

.PHONY: openapi_otm
openapi_otm:
	@./scripts/model-gen.sh otm-openapi pkg otm
