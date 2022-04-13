package marchy

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/artemmarkaryan/fisha/facade/pkg/logy"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const tag = "[marchy]"

type protoMsg interface {
	ProtoReflect() protoreflect.Message
}

func Obj[T protoMsg](ctx context.Context, r *http.Request) (o T, err error) {
	var b []byte

	b, err = io.ReadAll(r.Body)
	if err != nil {
		logy.Log(ctx).Debugf(tag+" reading error: %q", err)

		return
	}
	if len(b) == 0 {
		logy.Log(ctx).Debugf(tag+" empty object: '%+v' of type '%T'", o, o)

		return o, nil
	}

	err = json.Unmarshal(b, &o)

	return
}
