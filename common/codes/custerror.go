package codes

import (
	"github.com/hugplus/go-walker/core/errs"
)

func ErrInvalidParameter(reqId string, cause error) errs.IError {
	return errs.Err(InvalidParameter, reqId, cause)
}

func ErrNotFound(id, kind, reqId string, cause error) errs.IError {
	data := map[string]interface{}{"kind": kind, "id": id}
	return errs.ErrWithData(NotFound, reqId, cause, data)
}
