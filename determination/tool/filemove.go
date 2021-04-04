package tool

import (
	"os"
	"io"
	"net/http"
	"io/ioutil"
)
// @formkey form表单的input file name
// @position 具体保存地址
func Filemove(r *http.Request,formkey string,position string) []interface{} {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile(formkey)
	if err != nil {
		return []interface{}{false,err}
	}
	defer file.Close()

	Filename:= "./upload/"+position
	_,err = ioutil.ReadDir(Filename)
	if err != nil {
		err = os.MkdirAll(Filename, 0766)
		if err != nil {
		   panic("根目录找不到upload目录或该目录没有写权限")
		}
	}

	Filename = Filename+"/"+ handler.Filename
	f, err := os.OpenFile(Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return []interface{}{false,err}
	}
	defer f.Close()

	io.Copy(f, file)

	return []interface{}{true,Filename}
}