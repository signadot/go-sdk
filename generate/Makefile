# Generate the client and model libraries with sane defaults.

default:
	# create backup
	rm -rf backup
	mkdir backup
	mv ../client backup
	mv ../models backup
	# generate
	go run github.com/go-swagger/go-swagger/cmd/swagger@latest \
		generate client \
		--default-scheme=https \
		-f swagger.json \
		-t ../
	# fixup go.mod
	go get -u ../...
	go mod tidy
	# cleanup
	rm -rf backup



