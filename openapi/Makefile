.PHONY: gen
gen:
	@echo "Generating API code..."
		@for config in openapi/*.yaml; do \
			if [ -f $$config ]; then \
				echo "Generating from $$config..."; \
				oapi-codegen -config $$config init.yaml; \
			fi \
	done
	@echo "API generation completed"
