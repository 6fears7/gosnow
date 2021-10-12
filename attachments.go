package gosnow

import (
	"os"
	"path/filepath"
)

var (
	HASMAGIC = false
)

type Attachment struct {
	resource  Resource
	tableName string
}

func NewAttachment(resource Resource, tableName string) (A Attachment) {
	A.resource = resource
	A.tableName = tableName
	return
}

func (A Attachment) Get(sys_id string, limit int) (Response, error) {
	if sys_id == "" {
		query := map[string]interface{}{"table_name": A.tableName}
		return A.resource.Get(query, 1, 0, true, nil)
	}
	return A.resource.Get(map[string]interface{}{"table_sys_id": sys_id, "table_name": A.tableName}, limit, 0, true, nil)
}

func (A Attachment) Upload(sys_id, file_path string, multipart bool) (Response, error) {
	resource := A.resource
	name := filepath.Base(file_path)
	resource.Parameters.AddCustom(map[string]interface{}{"table_name": A.tableName, "table_sys_id": sys_id, "file_name": name})
	data, _ := os.ReadFile(file_path)

	headers := map[string]string{}
	var path_append string

	if multipart {
		headers["Content-Type"] = "multipart/form-data"
		path_append = "/upload"
	} else {
		//TODO Add ability to read magic from files

		if HASMAGIC {
			//magic.from_file(file_path, mime=True)
		} else {
			headers["Content-Type"] = ("text/plain")
		}
		path_append = "/file"
	}
	return resource.request("POST", data, path_append, headers, map[string]string{})
}

func (A Attachment) Delete(sys_id string) (map[string]interface{}, error) {
	query := map[string]interface{}{"sys_id": sys_id}
	return A.resource.Delete(query)
}
