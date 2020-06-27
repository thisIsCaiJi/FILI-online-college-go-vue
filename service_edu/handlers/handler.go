package handlers

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func HandlerError(err error,message string,a ...interface{}) {
	logrus.Errorf("ERROR: message: %s, err: %v\n",fmt.Sprintf(message,a),err)
}