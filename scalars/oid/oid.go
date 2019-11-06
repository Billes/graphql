package oid

import (
	"errors"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/globalsign/mgo/bson"
)

const (
	typeError = "Not a valid ObjectId"
)

func UnmarshalOID(v interface{}) (bson.ObjectId, error) {
	id, ok := v.(string)
	if !ok {
		return bson.NewObjectId(), errors.New(typeError)
	}

	if !bson.IsObjectIdHex(id) {
		return bson.NewObjectId(), errors.New(typeError)
	}

	return bson.ObjectIdHex(id), nil
}

func MarshalOID(id bson.ObjectId) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(id.Hex()))
	})
}

