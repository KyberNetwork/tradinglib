package finder

import "errors"

var (
	ErrRouteNotFound    = errors.New("route not found")
	ErrTokenInNotFound  = errors.New("token in not found")
	ErrTokenOutNotFound = errors.New("token out not found")
	ErrGasTokenRequired = errors.New("gas token required")
	ErrGasPriceRequired = errors.New("gas price required")
	ErrGasTokenNotFound = errors.New("gas token not found")
)
