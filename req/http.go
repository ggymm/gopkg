package req

type bodyType int

const (
	Json bodyType = iota
	Form
	FileForm
)

const (
	GET  = "GET"
	POST = "POST"
)

const (
	ContentTypeJson = "application/json"
	ContentTypeForm = "application/x-www-form-urlencoded"
)

type fileBody struct {
	file  File
	field FormField
}

type File struct {
	Name      string
	Path      string
	Content   []byte
	FieldName string
}

type FormField map[string]string
