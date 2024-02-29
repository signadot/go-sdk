# Signadot Go SDK

This repository houses the Signadot Go SDK, a Signadot API library
generated with swagger.

NOTE: This SDK is still in development and may be changed in ways that break consumers.

## Generation

The Go-SDK is consumed by the Signadot CLI and as such supports local workloads
in sandboxes.   However, local workloads in sandboxes are not really usable 
as is at the HTTP API level and require a lot of supporting code, root access,
etc.  As a result, the Go-SDK _must_ include the local workload specs in the 
swagger.json, whereas other SDKs should not include it.

The source of truth for the swagger.json in the `signadot` repository.  To
include the local workloads, one should remove the `swaggerignore` annotations
related to local and run `make swagger` in the signadot repository.  This will
create a swagger.json that can be used here in the Go-SDK (located at `generate/swagger.json`).

With this in mind, the directory `generate` contains all sources for generating the code.
See the Makefile there.

## Issues

To file an issue, please use our [community issue tracker](https://github.com/signadot/community/issues).


