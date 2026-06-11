# Generate the client and model libraries with sane defaults.

default:
	# create backup
	rm -rf backup
	mkdir backup
	mv ../client backup
	mv ../models backup

	# unwrap allOf fields
	./unwrap_allof.sh

	# generate using go-swagger (pinned: @latest drifts and produces
	# non-reproducible diffs / toolchain mismatches)
	go run github.com/go-swagger/go-swagger/cmd/swagger@v0.34.1 \
		generate client \
		--default-scheme=https \
		-f swagger.unwrapped.json \
		-t ../

	# fixup go.mod BEFORE building: the freshly generated client can reference
	# symbols from a newer go-openapi/runtime than go.mod currently pins (e.g.
	# runtime.ContextualTransport), so bump deps first or `go build` fails.
	go get -u ../...
	go mod tidy

	(cd .. && go build ./models/... ./client/...)

	# cleanup
	rm swagger.unwrapped.json
	rm -rf backup

