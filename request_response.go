package web

import (
	"net/http"
)

type Parameter struct {
	Query  Query
	Path   Path
	Header Header
	Cookie Cookie
	Body   Body
}

type Path map[string]PathParam

type PathParam struct {
	Description string
	Value       any
}

type Query map[string]QueryParam

type QueryParam struct {
	Optional    bool
	Description string
	Value       any
}

type Header map[string]HeaderField

type HeaderField struct {
	Optional    bool
	Description string
	Value       any
}

type Cookie map[string]CookieField

type CookieField struct {
	Optional    bool
	Description string
	Value       any
}

type Body struct {
	Description string
	Optional    bool
	Value       any
}

type Responses struct {
	Default any

	// All status codes can be found at http/status.go
	StatusContinue           any
	StatusSwitchingProtocols any
	StatusProcessing         any
	StatusEarlyHints         any

	StatusOK                   any
	StatusCreated              any
	StatusAccepted             any
	StatusNonAuthoritativeInfo any
	StatusNoContent            any
	StatusResetContent         any
	StatusPartialContent       any
	StatusMultiStatus          any
	StatusAlreadyReported      any
	StatusIMUsed               any

	StatusMultipleChoices   any
	StatusMovedPermanently  any
	StatusFound             any
	StatusSeeOther          any
	StatusNotModified       any
	StatusUseProxy          any
	StatusTemporaryRedirect any
	StatusPermanentRedirect any

	StatusBadRequest                   any
	StatusUnauthorized                 any
	StatusPaymentRequired              any
	StatusForbidden                    any
	StatusNotFound                     any
	StatusMethodNotAllowed             any
	StatusNotAcceptable                any
	StatusProxyAuthRequired            any
	StatusRequestTimeout               any
	StatusConflict                     any
	StatusGone                         any
	StatusLengthRequired               any
	StatusPreconditionFailed           any
	StatusRequestEntityTooLarge        any
	StatusRequestURITooLong            any
	StatusUnsupportedMediaType         any
	StatusRequestedRangeNotSatisfiable any
	StatusExpectationFailed            any
	StatusTeapot                       any
	StatusMisdirectedRequest           any
	StatusUnprocessableEntity          any
	StatusLocked                       any
	StatusFailedDependency             any
	StatusTooEarly                     any
	StatusUpgradeRequired              any
	StatusPreconditionRequired         any
	StatusTooManyRequests              any
	StatusRequestHeaderFieldsTooLarge  any
	StatusUnavailableForLegalReasons   any

	StatusInternalServerError           any
	StatusNotImplemented                any
	StatusBadGateway                    any
	StatusServiceUnavailable            any
	StatusGatewayTimeout                any
	StatusHTTPVersionNotSupported       any
	StatusVariantAlsoNegotiates         any
	StatusInsufficientStorage           any
	StatusLoopDetected                  any
	StatusNotExtended                   any
	StatusNetworkAuthenticationRequired any
}

func (r *Responses) Iterate() func(func(status int, value any) bool) {
	return func(yield func(status int, value any) bool) {
		if r.Default != nil {
			if !yield(0, r.Default) {
				return
			}
		}
		if r.StatusAccepted != nil {
			if !yield(http.StatusAccepted, r.StatusAccepted) {
				return
			}
		}
		if r.StatusAlreadyReported != nil {
			if !yield(http.StatusAlreadyReported, r.StatusAlreadyReported) {
				return
			}
		}
		if r.StatusBadGateway != nil {
			if !yield(http.StatusBadGateway, r.StatusBadGateway) {
				return
			}
		}
		if r.StatusConflict != nil {
			if !yield(http.StatusConflict, r.StatusConflict) {
				return
			}
		}
		if r.StatusContinue != nil {
			if !yield(http.StatusContinue, r.StatusContinue) {
				return
			}
		}
		if r.StatusCreated != nil {
			if !yield(http.StatusCreated, r.StatusCreated) {
				return
			}
		}
		if r.StatusEarlyHints != nil {
			if !yield(http.StatusEarlyHints, r.StatusEarlyHints) {
				return
			}
		}
		if r.StatusExpectationFailed != nil {
			if !yield(http.StatusExpectationFailed, r.StatusExpectationFailed) {
				return
			}
		}
		if r.StatusFailedDependency != nil {
			if !yield(http.StatusFailedDependency, r.StatusFailedDependency) {
				return
			}
		}
		if r.StatusForbidden != nil {
			if !yield(http.StatusForbidden, r.StatusForbidden) {
				return
			}
		}
		if r.StatusFound != nil {
			if !yield(http.StatusFound, r.StatusFound) {
				return
			}
		}
		if r.StatusGatewayTimeout != nil {
			if !yield(http.StatusGatewayTimeout, r.StatusGatewayTimeout) {
				return
			}
		}
		if r.StatusGone != nil {
			if !yield(http.StatusGone, r.StatusGone) {
				return
			}
		}
		if r.StatusHTTPVersionNotSupported != nil {
			if !yield(http.StatusHTTPVersionNotSupported, r.StatusHTTPVersionNotSupported) {
				return
			}
		}
		if r.StatusIMUsed != nil {
			if !yield(http.StatusIMUsed, r.StatusIMUsed) {
				return
			}
		}
		if r.StatusInsufficientStorage != nil {
			if !yield(http.StatusInsufficientStorage, r.StatusInsufficientStorage) {
				return
			}
		}
		if r.StatusInternalServerError != nil {
			if !yield(http.StatusInternalServerError, r.StatusInternalServerError) {
				return
			}
		}
		if r.StatusLengthRequired != nil {
			if !yield(http.StatusLengthRequired, r.StatusLengthRequired) {
				return
			}
		}
		if r.StatusLocked != nil {
			if !yield(http.StatusLocked, r.StatusLocked) {
				return
			}
		}
		if r.StatusLoopDetected != nil {
			if !yield(http.StatusLoopDetected, r.StatusLoopDetected) {
				return
			}
		}
		if r.StatusMethodNotAllowed != nil {
			if !yield(http.StatusMethodNotAllowed, r.StatusMethodNotAllowed) {
				return
			}
		}
		if r.StatusMisdirectedRequest != nil {
			if !yield(http.StatusMisdirectedRequest, r.StatusMisdirectedRequest) {
				return
			}
		}
		if r.StatusMovedPermanently != nil {
			if !yield(http.StatusMovedPermanently, r.StatusMovedPermanently) {
				return
			}
		}
		if r.StatusMultiStatus != nil {
			if !yield(http.StatusMultiStatus, r.StatusMultiStatus) {
				return
			}
		}
		if r.StatusMultipleChoices != nil {
			if !yield(http.StatusMultipleChoices, r.StatusMultipleChoices) {
				return
			}
		}
		if r.StatusNetworkAuthenticationRequired != nil {
			if !yield(http.StatusNetworkAuthenticationRequired, r.StatusNetworkAuthenticationRequired) {
				return
			}
		}
		if r.StatusNoContent != nil {
			if !yield(http.StatusNoContent, r.StatusNoContent) {
				return
			}
		}
		if r.StatusNonAuthoritativeInfo != nil {
			if !yield(http.StatusNonAuthoritativeInfo, r.StatusNonAuthoritativeInfo) {
				return
			}
		}
		if r.StatusNotAcceptable != nil {
			if !yield(http.StatusNotAcceptable, r.StatusNotAcceptable) {
				return
			}
		}
		if r.StatusNotExtended != nil {
			if !yield(http.StatusNotExtended, r.StatusNotExtended) {
				return
			}
		}
		if r.StatusNotFound != nil {
			if !yield(http.StatusNotFound, r.StatusNotFound) {
				return
			}
		}
		if r.StatusNotImplemented != nil {
			if !yield(http.StatusNotImplemented, r.StatusNotImplemented) {
				return
			}
		}
		if r.StatusNotModified != nil {
			if !yield(http.StatusNotModified, r.StatusNotModified) {
				return
			}
		}
		if r.StatusOK != nil {
			if !yield(http.StatusOK, r.StatusOK) {
				return
			}
		}
		if r.StatusPartialContent != nil {
			if !yield(http.StatusPartialContent, r.StatusPartialContent) {
				return
			}
		}
		if r.StatusPaymentRequired != nil {
			if !yield(http.StatusPaymentRequired, r.StatusPaymentRequired) {
				return
			}
		}
		if r.StatusPermanentRedirect != nil {
			if !yield(http.StatusPermanentRedirect, r.StatusPermanentRedirect) {
				return
			}
		}
		if r.StatusPreconditionFailed != nil {
			if !yield(http.StatusPreconditionFailed, r.StatusPreconditionFailed) {
				return
			}
		}
		if r.StatusPreconditionRequired != nil {
			if !yield(http.StatusPreconditionRequired, r.StatusPreconditionRequired) {
				return
			}
		}
		if r.StatusProcessing != nil {
			if !yield(http.StatusProcessing, r.StatusProcessing) {
				return
			}
		}
		if r.StatusProxyAuthRequired != nil {
			if !yield(http.StatusProxyAuthRequired, r.StatusProxyAuthRequired) {
				return
			}
		}
		if r.StatusRequestEntityTooLarge != nil {
			if !yield(http.StatusRequestEntityTooLarge, r.StatusRequestEntityTooLarge) {
				return
			}
		}
		if r.StatusRequestHeaderFieldsTooLarge != nil {
			if !yield(http.StatusRequestHeaderFieldsTooLarge, r.StatusRequestHeaderFieldsTooLarge) {
				return
			}
		}
		if r.StatusRequestTimeout != nil {
			if !yield(http.StatusRequestTimeout, r.StatusRequestTimeout) {
				return
			}
		}
		if r.StatusRequestURITooLong != nil {
			if !yield(http.StatusRequestURITooLong, r.StatusRequestURITooLong) {
				return
			}
		}
		if r.StatusRequestedRangeNotSatisfiable != nil {
			if !yield(http.StatusRequestedRangeNotSatisfiable, r.StatusRequestedRangeNotSatisfiable) {
				return
			}
		}
		if r.StatusResetContent != nil {
			if !yield(http.StatusResetContent, r.StatusResetContent) {
				return
			}
		}
		if r.StatusSeeOther != nil {
			if !yield(http.StatusSeeOther, r.StatusSeeOther) {
				return
			}
		}
		if r.StatusServiceUnavailable != nil {
			if !yield(http.StatusServiceUnavailable, r.StatusServiceUnavailable) {
				return
			}
		}
		if r.StatusServiceUnavailable != nil {
			if !yield(http.StatusServiceUnavailable, r.StatusServiceUnavailable) {
				return
			}
		}
		if r.StatusSwitchingProtocols != nil {
			if !yield(http.StatusSwitchingProtocols, r.StatusSwitchingProtocols) {
				return
			}
		}
		if r.StatusTeapot != nil {
			if !yield(http.StatusTeapot, r.StatusTeapot) {
				return
			}
		}
		if r.StatusTemporaryRedirect != nil {
			if !yield(http.StatusTemporaryRedirect, r.StatusTemporaryRedirect) {
				return
			}
		}
		if r.StatusTooEarly != nil {
			if !yield(http.StatusTooEarly, r.StatusTooEarly) {
				return
			}
		}
		if r.StatusTooManyRequests != nil {
			if !yield(http.StatusTooManyRequests, r.StatusTooManyRequests) {
				return
			}
		}
		if r.StatusUnauthorized != nil {
			if !yield(http.StatusUnauthorized, r.StatusUnauthorized) {
				return
			}
		}
		if r.StatusUnavailableForLegalReasons != nil {
			if !yield(http.StatusUnavailableForLegalReasons, r.StatusUnavailableForLegalReasons) {
				return
			}
		}
		if r.StatusUnprocessableEntity != nil {
			if !yield(http.StatusUnprocessableEntity, r.StatusUnprocessableEntity) {
				return
			}
		}
		if r.StatusUnsupportedMediaType != nil {
			if !yield(http.StatusUnsupportedMediaType, r.StatusUnsupportedMediaType) {
				return
			}
		}
		if r.StatusUpgradeRequired != nil {
			if !yield(http.StatusUpgradeRequired, r.StatusUpgradeRequired) {
				return
			}
		}
		if r.StatusUseProxy != nil {
			if !yield(http.StatusUseProxy, r.StatusUseProxy) {
				return
			}
		}
		if r.StatusVariantAlsoNegotiates != nil {
			if !yield(http.StatusVariantAlsoNegotiates, r.StatusVariantAlsoNegotiates) {
				return
			}
		}
	}
}

func (r *Responses) Set(status int, value any) {
	switch status {
	case 0:
		r.Default = value
	case http.StatusContinue:
		r.StatusContinue = value
	case http.StatusSwitchingProtocols:
		r.StatusSwitchingProtocols = value
	case http.StatusProcessing:
		r.StatusProcessing = value
	case http.StatusEarlyHints:
		r.StatusEarlyHints = value
	case http.StatusOK:
		r.StatusOK = value
	case http.StatusCreated:
		r.StatusCreated = value
	case http.StatusAccepted:
		r.StatusAccepted = value
	case http.StatusNonAuthoritativeInfo:
		r.StatusNonAuthoritativeInfo = value
	case http.StatusNoContent:
		r.StatusNoContent = value
	case http.StatusResetContent:
		r.StatusResetContent = value
	case http.StatusPartialContent:
		r.StatusPartialContent = value
	case http.StatusMultiStatus:
		r.StatusMultiStatus = value
	case http.StatusAlreadyReported:
		r.StatusAlreadyReported = value
	case http.StatusIMUsed:
		r.StatusIMUsed = value
	case http.StatusMultipleChoices:
		r.StatusMultipleChoices = value
	case http.StatusMovedPermanently:
		r.StatusMovedPermanently = value
	case http.StatusFound:
		r.StatusFound = value
	case http.StatusSeeOther:
		r.StatusSeeOther = value
	case http.StatusNotModified:
		r.StatusNotModified = value
	case http.StatusUseProxy:
		r.StatusUseProxy = value
	case http.StatusTemporaryRedirect:
		r.StatusTemporaryRedirect = value
	case http.StatusPermanentRedirect:
		r.StatusPermanentRedirect = value
	case http.StatusBadRequest:
		r.StatusBadRequest = value
	case http.StatusUnauthorized:
		r.StatusUnauthorized = value
	case http.StatusPaymentRequired:
		r.StatusPaymentRequired = value
	case http.StatusForbidden:
		r.StatusForbidden = value
	case http.StatusNotFound:
		r.StatusNotFound = value
	case http.StatusMethodNotAllowed:
		r.StatusMethodNotAllowed = value
	case http.StatusNotAcceptable:
		r.StatusNotAcceptable = value
	case http.StatusProxyAuthRequired:
		r.StatusProxyAuthRequired = value
	case http.StatusRequestTimeout:
		r.StatusRequestTimeout = value
	case http.StatusConflict:
		r.StatusConflict = value
	case http.StatusGone:
		r.StatusGone = value
	case http.StatusLengthRequired:
		r.StatusLengthRequired = value
	case http.StatusPreconditionFailed:
		r.StatusPreconditionFailed = value
	case http.StatusRequestEntityTooLarge:
		r.StatusRequestEntityTooLarge = value
	case http.StatusRequestURITooLong:
		r.StatusRequestURITooLong = value
	case http.StatusUnsupportedMediaType:
		r.StatusUnsupportedMediaType = value
	case http.StatusRequestedRangeNotSatisfiable:
		r.StatusRequestedRangeNotSatisfiable = value
	case http.StatusExpectationFailed:
		r.StatusExpectationFailed = value
	case http.StatusTeapot:
		r.StatusTeapot = value
	case http.StatusMisdirectedRequest:
		r.StatusMisdirectedRequest = value
	case http.StatusUnprocessableEntity:
		r.StatusUnprocessableEntity = value
	case http.StatusLocked:
		r.StatusLocked = value
	case http.StatusFailedDependency:
		r.StatusFailedDependency = value
	case http.StatusTooEarly:
		r.StatusTooEarly = value
	case http.StatusUpgradeRequired:
		r.StatusUpgradeRequired = value
	case http.StatusPreconditionRequired:
		r.StatusPreconditionRequired = value
	case http.StatusTooManyRequests:
		r.StatusTooManyRequests = value
	case http.StatusRequestHeaderFieldsTooLarge:
		r.StatusRequestHeaderFieldsTooLarge = value
	case http.StatusUnavailableForLegalReasons:
		r.StatusUnavailableForLegalReasons = value
	case http.StatusInternalServerError:
		r.StatusInternalServerError = value
	case http.StatusNotImplemented:
		r.StatusNotImplemented = value
	case http.StatusBadGateway:
		r.StatusBadGateway = value
	case http.StatusServiceUnavailable:
		r.StatusServiceUnavailable = value
	case http.StatusGatewayTimeout:
		r.StatusGatewayTimeout = value
	case http.StatusHTTPVersionNotSupported:
		r.StatusHTTPVersionNotSupported = value
	case http.StatusVariantAlsoNegotiates:
		r.StatusVariantAlsoNegotiates = value
	case http.StatusInsufficientStorage:
		r.StatusInsufficientStorage = value
	case http.StatusLoopDetected:
		r.StatusLoopDetected = value
	case http.StatusNotExtended:
		r.StatusNotExtended = value
	case http.StatusNetworkAuthenticationRequired:
		r.StatusNetworkAuthenticationRequired = value
	}
}
