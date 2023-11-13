package common

import (
	"errors"
	"log"
)

type RestError interface {
	Status() int
	Error() string
	Causes() interface{}
}

type ErrResp struct {
	ErrStatus int         `json:"status,omitempty"`
	ErrError  string      `json:"error,omitempty"`
	ErrCauses interface{} `json:"-"`
}

func (e ErrResp) Status() int {
	return e.ErrStatus
}

func (e ErrResp) Error() string {
	return e.Error()
}

func (e ErrResp) Causes() interface{} {
	return e.ErrCauses
}

func NewRestErr(status int, err string, causes interface{}) RestError {
	log.Println(causes)
	return ErrResp{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: err,
	}
}

var (
	BadRequest            = errors.New("Bad request")
	WrongCredentials      = errors.New("Wrong Credentials")
	NotFound              = errors.New("Not Found")
	Unauthorized          = errors.New("Unauthorized")
	Forbidden             = errors.New("Forbidden")
	PermissionDenied      = errors.New("Permission Denied")
	ExpiredCSRFError      = errors.New("Expired CSRF token")
	WrongCSRFToken        = errors.New("Wrong CSRF token")
	CSRFNotPresented      = errors.New("CSRF not presented")
	NotRequiredFields     = errors.New("No such required fields")
	BadQueryParams        = errors.New("Invalid query params")
	InternalServerError   = errors.New("Internal Server Error")
	RequestTimeoutError   = errors.New("Request Timeout")
	ExistsEmailError      = errors.New("User with given email already exists")
	InvalidJWTToken       = errors.New("Invalid JWT token")
	InvalidJWTClaims      = errors.New("Invalid JWT claims")
	InvalidAuthHeader     = errors.New("Invalid authenrization header")
	ErrJWTExpired         = errors.New("JWT token expired")
	NotAllowedImageHeader = errors.New("Not allowed image header")
	NoCookie              = errors.New("not found cookie header")
	NotExistAccount       = errors.New("Account doesnt exists")
	WrongPassword         = errors.New("Wrong password")
)
