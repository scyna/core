package scyna_proto

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

func (r *Request) Build(req *http.Request) error {
	if req == nil {
		return errors.New("natsproxy: Request cannot be nil")
	}

	buf := bytes.NewBuffer(r.Body)
	buf.Reset()
	if req.Body != nil {
		if _, err := io.Copy(buf, req.Body); err != nil {
			return err
		}
		if err := req.Body.Close(); err != nil {
			return err
		}
	}

	r.Body = buf.Bytes()
	return nil
}
