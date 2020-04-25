package Server

import "net/http"

func BadRequest(err error, writer http.ResponseWriter, request *http.Request) {
	Response(err, http.StatusBadRequest, writer, request)
	log.Error(err)
}

func Response(err error, status int, writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(status)
	_, _ = writer.Write([]byte(err.Error()))
	_ = request.Body.Close()
}

func GetHttpParam(param string, r *http.Request) string {
	params, ok := r.URL.Query()[param]
	if !ok {
		return ""
	}
	return params[0]
}
