package request

import "testing"

func TestRequest_Get(t *testing.T) {
	resp, err := New().
		Get("https://baidu.com").
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())
}

func TestRequest_Post(t *testing.T) {

}

func TestRequest_PostFile(t *testing.T) {
	resp, err := New().
		Post("http://localhost:8080/group1/upload").
		FileBody(
			File{
				Name:      "黑钻.png",
				Path:      "C:\\Users\\19679\\Documents\\黑钻.png",
				FieldName: "file",
			},
			FormField{
				"output": "json",
			},
		).Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())
}
