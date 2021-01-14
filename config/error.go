package config

import "github.com/pkg/errors"

var(
	ERROR_SURPPORT_INTERFACE_NULL_165 = errors.New("SupportInterface return null")
	ERROR_GET_BASE_URI_NULL = errors.New("GetBaseURI return null")
	ERROR_GET_TOKENID_URI_NULL = errors.New("GetTokenIDURI return null")
)
