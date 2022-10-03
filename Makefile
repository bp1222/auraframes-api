.PHONY: all
all:
	npx --yes @apidevtools/swagger-cli bundle api/aura.yaml --outfile bundle.yaml --type yaml
	npx --yes @openapitools/openapi-generator-cli generate --generator-key auraframes-client

.PHONY: clean
clean:
	rm -fR go
