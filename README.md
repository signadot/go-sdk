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

## Note
At this time, we have marked two properties related to Local with swaggerignore in Sandbox spec
so that these properties don't show up in the public API Reference.

```
Local []Local `json:"local,omitempty" swaggerignore:"true"`
LocalMachineID *string `json:"localMachineID,omitempty" validate:"optional" swaggerignore:"true"`
```

However, they are needed by the CLI. Hence, here's the step to get the swagger accounting for that.

```
1. remove the swaggerignore for the 2 local related properties mentioned above in the Signadot repo
2. generate swagger: `make swagger-control`
3. copy the generated swagger to the go-sdk swagger.json
4. run `make` in the generate directory in the go-sdk
```
