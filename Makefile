.PHONY: all
all:
	swagger-cli bundle api/aura.yaml --outfile bundle.yaml --type yaml && \
	openapi-generator generate -g go -o go -i bundle.yaml -p packageName=auraframes --git-user-id bp1222 --git-repo-id=auraframes-api/go

.PHONY: clean
clean:
	rm -fR go
