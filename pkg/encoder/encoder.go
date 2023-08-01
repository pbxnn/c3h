package encoder

import (
	"encoding/json"
	"fmt"
	nethttp "net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseEncoder encodes the object to the HTTP response.
func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if v == nil {
		return nil
	}
	if rd, ok := v.(http.Redirector); ok {
		url, code := rd.Redirect()
		nethttp.Redirect(w, r, url, code)
		return nil
	}
	codec, _ := http.CodecForRequest(r, "Accept")

	v1, e := codec.Marshal(v)
	fmt.Println(v1, e)

	resp := response{
		Code:    200,
		Message: "",
		Data:    v,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", strings.Join([]string{"application", codec.Name()}, "/"))
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// ErrorEncoder encodes the error to the HTTP response.
func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	resp := response{
		Code:    se.Code,
		Message: se.Reason + ", " + se.Message,
		Data:    se.Metadata,
	}

	body, err := codec.Marshal(resp)
	if err != nil {
		w.WriteHeader(nethttp.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", strings.Join([]string{"application", codec.Name()}, "/"))
	w.WriteHeader(int(se.Code))
	_, _ = w.Write(body)
}
