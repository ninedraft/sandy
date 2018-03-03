package sanderr

import (
	"bytes"
	"fmt"
	"text/template"
)

const (
	errorTemplate = "{{.Message}}{{if .Details}}: {{range $detail := .Details}}{{$detail}}{{end}}{{end}}"
)

type Err struct {
	Message string   `json: "message"`
	Details []string `json:"details"`
}

func (err *Err) AddError(er error) *Err {
	err.Details = append(err.Details, er.Error())
	return err
}

func (err *Err) AddStr(str string) *Err {
	err.Details = append(err.Details, str)
	return err
}

func (err *Err) AddStrF(f string, vals ...interface{}) *Err {
	return err.AddStr(fmt.Sprintf(f, vals...))
}

func (err *Err) Error() string {
	templ, er := template.New("").Parse(errorTemplate)
	if er != nil {
		panic(er)
	}
	buf := bytes.NewBufferString(errorTemplate)
	buf.Reset()
	er = templ.Execute(buf, err)
	if er != nil {
		panic(er)
	}
	return buf.String()
}
