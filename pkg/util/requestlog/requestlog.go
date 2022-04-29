package requestlog

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	frpLog "github.com/salessync/frp/pkg/util/log"
	"github.com/salessync/frp/pkg/util/sqlite"
)

type RequestLog struct {
	method      string
	originalUrl string
	url         string
	queryParams string
	body        []byte
	headers     string
}

func logToDb(requestLog RequestLog) {
	insertQuery := fmt.Sprintf("INSERT INTO request(original_host, url, method, query_params, body, headers) VALUES ('%s', '%s', '%s', '%s', '%s', '%s')", requestLog.originalUrl, requestLog.url, requestLog.method, requestLog.queryParams, requestLog.body, requestLog.headers)
	sqlite.RunQuery(insertQuery)
}

func LogRequest(req *http.Request, url string, oldHost string, baseUrl string) {
	serializedQueryParamsStr := serializeQueryParams(req)
	serializedBodyStr := serializeRequestBody(req)
	serializedHeadersStr := serializeRequestHeaders(req)

	protocol := "http://"
	if req.TLS != nil {
		protocol = "https://"
	}

	fullUrl := protocol + baseUrl + req.URL.Path

	requestLog := RequestLog{
		req.Method,
		oldHost,
		fullUrl,
		serializedQueryParamsStr,
		serializedBodyStr,
		serializedHeadersStr,
	}

	logToDb(requestLog)
}

func serializeQueryParams(req *http.Request) string {
	req.ParseForm()

	serializedParams := ""
	for k, v := range req.Form {
		serializedParams = serializedParams + fmt.Sprintf("%s=%s&", k, v[0])
	}

	return strings.TrimSuffix(serializedParams, "&")
}

func serializeRequestBody(req *http.Request) []byte {
	if req.Body == nil {
		return []byte{}
	}

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		frpLog.Warn(err.Error())

		return []byte{}
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
	return reqBody
}

func serializeRequestHeaders(req *http.Request) string {
	serializedHeaders := ""

	for k, v := range req.Header {
		serializedHeaders = serializedHeaders + fmt.Sprintf("%q: %q|::|", k, v[0])
	}

	return strings.TrimSuffix(serializedHeaders, "|::|")
}
