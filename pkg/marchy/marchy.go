package marchy

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const tag = "[marchy]"

type protoMsg interface {
	ProtoReflect() protoreflect.Message
}

func Obj[T protoMsg](ctx context.Context, r io.ReadCloser) (o T, err error) {
	var b []byte

	b, err = io.ReadAll(r)
	if err != nil {
		logy.Log(ctx).Debugf(tag+" reading error: %q", err)

		return
	}
	if len(b) == 0 {
		logy.Log(ctx).Debugf(tag+" empty object: '%+v' of type '%T'", o, o)

		return o, nil
	}

	if err = json.Unmarshal(b, &o); err != nil {
		logy.Log(ctx).Errorf("unmarshalling error: ", err)

		return o, err
	}

	return
}

func Force(obj any) (r []byte) {
	r, _ = json.Marshal(obj)
	return
}

func ForceReader(obj any) (r io.Reader) {
	b, _ := json.Marshal(obj)
	return bytes.NewReader(b)
}
