# Generate the client and model libraries with sane defaults.

default:
	# create backup
	rm -rf backup
	mkdir backup
	mv ../client backup
	mv ../models backup

	# unwrap allOf fields
	./unwrap_allof.sh

	# generate using go-swagger. Pinned: @latest (v0.34.x) generates a
	# client.New that takes runtime.ContextualTransport, which breaks
	# consumers (e.g. the CLI) whose transport returns runtime.ClientTransport.
	# v0.33.0 is the last release that generates the ClientTransport signature.
	go run github.com/go-swagger/go-swagger/cmd/swagger@v0.33.0 \
		generate client \
		--default-scheme=https \
		-f swagger.unwrapped.json \
		-t ../
	(cd .. && go build ./models/... ./client/...)

	# fixup go.mod. NOTE: intentionally no `go get -u ../...` here: it bumps
	# go-openapi/runtime to v0.32.3, which also forces consumers' module graphs
	# up. `go mod tidy` alone keeps runtime at v0.29.x so the client is a
	# drop-in and the CLI compiles with zero dependency changes.
	go mod tidy

	# cleanup
	rm swagger.unwrapped.json
	rm -rf backup
