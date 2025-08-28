#!/usr/bin/env bash
set -euo pipefail

# When using swaggo/swag to generate OpenAPI (Swagger) from Go code, if a field
# has a GoDoc comment or an extension (e.g. `extensions:"x-nullable"`), swag will
# emit the property as an `allOf` wrapper around the `$ref` in order to attach the
# extra metadata. Example:
#
#   "ttl": {
#     "allOf": [
#       { "$ref": "#/definitions/RouteGroupTTL" }
#     ],
#     "description": "TTL define when/how the route group should be deleted",
#     "x-nullable": true
#   }
#
# This is technically spec-compliant, but go-swagger codegen interprets it as a
# *composed type*, and generates an **inline struct embedding RouteGroupTTL**:
#
#   TTL struct {
#       RouteGroupTTL
#   } `json:"ttl,omitempty"`
#
# That’s not what we want — we want a plain pointer:
#
#   TTL *RouteGroupTTL `json:"ttl,omitempty"`
#
# The trick: if the property is just a direct `$ref` without siblings, go-swagger
# generates a pointer. So we post-process the spec to unwrap these single-ref allOfs.
#
# After running this script, the above becomes:
#
#   "ttl": {
#     "$ref": "#/definitions/RouteGroupTTL"
#   }
#
# which then produces the correct Go type.


IN="swagger.json"
OUT="swagger.unwrapped.json"

cp "$IN" "${IN}.bak"

jq '
  def unwrap:
    if type=="object" and has("allOf")
       and (.allOf | type=="array" and length==1)
       and (.allOf[0] | type=="object" and has("$ref"))
    then { "$ref": .allOf[0]["$ref"] }
    else .
    end;
  walk(unwrap)
' "$IN" > "$OUT"

echo "✔ Unwrapped single-ref allOfs."
rm "${IN}.bak"
