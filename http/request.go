package http

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/goccy/go-json"
)

type Request struct {
	url    string
	method string

	headers map[string]string

	bodyRaw  any
	bodyType bodyType
}

func New() *Request {
	return &Request{}
}

func (r *Request) Get(url string) *Request {
	r.url = url
	r.method = GET
	return r
}

func (r *Request) Post(url string) *Request {
	r.url = url
	r.method = POST
	return r
}

func (r *Request) Header(key, value string) *Request {
	if r.headers == nil {
		r.headers = make(map[string]string)
	}
	r.headers[key] = value
	return r
}

func (r *Request) Headers(headers map[string]string) *Request {
	if r.headers == nil {
		r.headers = make(map[string]string)
	}
	for k, v := range headers {
		r.headers[k] = v
	}
	return r
}

func (r *Request) Query(params any) *Request {
	switch v := params.(type) {
	case map[string]string:
		values := url.Values{}
		for k := range v {
			values.Set(k, v[k])
		}
		if !strings.Contains(r.url, "?") {
			r.url += "?" + values.Encode()
		} else {
			r.url += "&" + values.Encode()
		}
	case string:
		r.url += "?" + v
	}
	return r
}

func (r *Request) FormBody(body any) *Request {
	// 保存请求体
	r.bodyRaw = body
	r.bodyType = Form

	// 设置 Content-Type 为 application/x-www-form-urlencoded
	r.Header("Content-Type", ContentTypeForm)
	return r
}

func (r *Request) JsonBody(body any) *Request {
	// 保存请求体
	r.bodyRaw = body
	r.bodyType = Json

	// 设置 Content-Type 为 application/json
	r.Header("Content-Type", ContentTypeJson)
	return r
}

func (r *Request) FileBody(file File, field FormField) *Request {
	// 保存请求体
	r.bodyRaw = &fileBody{
		file:  file,
		field: field,
	}
	r.bodyType = FileForm
	return r
}

func (r *Request) Execute() (*Response, error) {
	var (
		err error

		req    *http.Request
		resp   *http.Response
		client = http.Client{}
	)

	req, err = http.NewRequest(r.method, r.url, nil)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for k, v := range r.headers {
		req.Header.Set(k, v)
	}

	// 设置请求体
	if r.bodyRaw != nil {
		switch r.bodyType {
		case Json:
			switch v := r.bodyRaw.(type) {
			case struct{}, map[string]any, []any:
				var val []byte
				val, err = json.Marshal(v)
				if err != nil {
					return nil, err
				}

				buf := bytes.NewBuffer(val)
				req.Body = io.NopCloser(buf)
				req.ContentLength = int64(buf.Len())
			case string:
				buf := strings.NewReader(v)
				req.Body = io.NopCloser(buf)
				req.ContentLength = int64(buf.Len())
			}
		case Form:
			switch v := r.bodyRaw.(type) {
			case url.Values:
				buf := strings.NewReader(v.Encode())
				req.Body = io.NopCloser(buf)
				req.ContentLength = int64(buf.Len())
			case string:
				buf := strings.NewReader(v)
				req.Body = io.NopCloser(buf)
				req.ContentLength = int64(buf.Len())
			case map[string]string:
				val := url.Values{}
				for k := range v {
					val.Set(k, v[k])
				}

				buf := strings.NewReader(val.Encode())
				req.Body = io.NopCloser(buf)
				req.ContentLength = int64(buf.Len())
			}
		case FileForm:
			v := r.bodyRaw.(*fileBody)
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)

			// 设置参数
			for k := range v.field {
				err = writer.WriteField(k, v.field[k])
				if err != nil {
					return nil, err
				}
			}

			// 设置文件参数
			var part io.Writer
			if v.file.Content != nil {
				part, err = writer.CreateFormFile(v.file.FieldName, v.file.Name)
				if err != nil {
					return nil, err
				}
				_, err = part.Write(v.file.Content)
				if err != nil {
					return nil, err
				}
			} else if v.file.Path != "" {
				var file *os.File
				file, err = os.Open(v.file.Path)
				if err != nil {
					return nil, err
				}

				part, err = writer.CreateFormFile(v.file.FieldName, v.file.Name)
				if err != nil {
					return nil, err
				}

				_, err = io.Copy(part, file)
				if err != nil {
					return nil, err
				}
				_ = file.Close()
			}
			err = writer.Close()
			if err != nil {
				return nil, err
			}

			req.Body = io.NopCloser(body)
			req.ContentLength = int64(body.Len())

			// 设置 Content-Type
			req.Header.Set("Content-Type", writer.FormDataContentType())
		}
	}

	// 发送请求
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	// 构造响应值
	return &Response{raw: resp}, nil
}
