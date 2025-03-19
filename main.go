package main

import (
	_ "github.com/aws/aws-sdk-go-v2"
	_ "github.com/aws/aws-sdk-go-v2/service/kms"
	_ "github.com/danielgtaylor/huma/v2"
	_ "github.com/decred/dcrd/dcrec/secp256k1/v4"
	_ "github.com/evanphx/json-patch/v5"
	_ "github.com/go-resty/resty/v2"
	_ "github.com/gofiber/fiber/v2"
	_ "github.com/gojuno/minimock/v3"
	_ "github.com/google/uuid"
	_ "github.com/invopop/jsonschema"
	_ "github.com/lestrrat-go/jwx/v2"
	_ "github.com/mr-tron/base58"
	_ "github.com/multiformats/go-multibase"
	_ "github.com/multiformats/go-multicodec"
	_ "github.com/multiformats/go-multihash"
	_ "github.com/santhosh-tekuri/jsonschema/v5"
	_ "github.com/stretchr/testify"
	_ "github.com/tidwall/gjson"
)

func main() {

}
