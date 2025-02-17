// Copyright 2021-2023 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package client

import requestv1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"

// RequestOpt defines per-request options.
type RequestOpt func(*reqOpt)

type reqOpt struct {
	auxData     *requestv1.AuxData
	includeMeta bool
}

// AuxDataJWT sets the JWT to be used as auxiliary data for the request.
func AuxDataJWT(token, keySetID string) RequestOpt {
	return func(opts *reqOpt) {
		if opts.auxData == nil {
			opts.auxData = &requestv1.AuxData{}
		}

		if opts.auxData.Jwt == nil {
			opts.auxData.Jwt = &requestv1.AuxData_JWT{}
		}

		opts.auxData.Jwt.Token = token
		opts.auxData.Jwt.KeySetId = keySetID
	}
}

// IncludeMeta sets the flag on requests that support it to signal that evaluation metadata should be sent back with the response.
func IncludeMeta(f bool) RequestOpt {
	return func(opt *reqOpt) {
		opt.includeMeta = f
	}
}
