# Go SDK

## generation

docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger generate client --default-scheme=https

(I couldn't figure out how to get swagger-codegen-cli to work reasonably, goswagger seems easier)

