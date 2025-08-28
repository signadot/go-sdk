# Generate the client and model libraries with sane defaults.

default:
	# create backup
	rm -rf backup
	mkdir backup
	mv ../client backup
	mv ../models backup

	# unwrap allOf fields
	./unwrap_allof.sh

	# generate using go-swagger
	go run github.com/go-swagger/go-swagger/cmd/swagger@latest \
		generate client \
		--default-scheme=https \
		-f swagger.unwrapped.json \
		-t ../
	(cd .. && go build ./models/... ./client/...)

	# fixup go.mod
	go get -u ../...
	go mod tidy

	# cleanup
	rm swagger.unwrapped.json
	rm -rf backup

