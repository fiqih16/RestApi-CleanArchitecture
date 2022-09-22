package helper

import "strings"

// Respons digunakan untuk pengembalian json bentuk statis
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// Objek EmptyObj digunakan ketika data tidak ingin menjadi nol di json
type EmptyObj struct {

}

// Metode BuildResponse adalah menyuntikkan nilai data ke respons sukses yang dinamis
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

//Metode BuildErrorResponse adalah menyuntikkan nilai data ke respons gagal dinamis
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}