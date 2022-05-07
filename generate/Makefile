# Generate the client and model libraries with sane defaults.

default:
	rm -rf backup
	mkdir backup
	mv ../client backup
	mv ../models backup
	go run github.com/go-swagger/go-swagger/cmd/swagger@latest \
		generate client \
		--default-scheme=https \
		-f swagger.json \
		-t ../
	go get -u ../...
	go mod tidy
	rm -rf backup



