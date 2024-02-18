package exceptions

import "errors"

var PARAM_NOT_PROVIDER = errors.New("error en los parametros. Debe enviar 'SecretName'")
