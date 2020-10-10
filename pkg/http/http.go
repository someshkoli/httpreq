package http

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// =====================================
type Response struct {
	Data   string
	Body   string
	Status int
	Time   time.Duration
	Err    error
}

func NewResponse(body string, data string, status int, t time.Duration, err error) *Response {
	return &Response{
		Body:   body,
		Status: status,
		Err:    err,
		Time:   t,
	}
}

func makeErrResponse(err error) *Response {
	return NewResponse("", "", 0, 0, err)
}

// =====================================
func getRequestString(host string, path string) string {
	return fmt.Sprintf("GET %s HTTP/1.0\r\nHost: %s\r\nAccept: */*\r\n\r\n", path, host)
}

func parseResponse(response string) *Response {
	responseSplit := strings.Split(response, "\r\n\r\n")
	code, _ := strconv.Atoi(strings.Split(responseSplit[0], " ")[1])
	res := Response{}
	res.Body = strings.Join(responseSplit[1:], "\r\n\r\n")
	res.Status = code
	res.Data = response
	return &res
}

func makepathWithQuery(parsed *url.URL) (string, error) {
	path := parsed.Path
	query := parsed.Query().Encode()
	qPath := ""
	if path == "" {
		path = "/"
	}
	qPath += path
	if query != "" {
		qPath += "?"
		qPath += query
	}
	return qPath, nil
}

// Request represents request object and carries out all other tasks
type Request struct {
	URL *url.URL `json:"URL"`
}

func MakeRequest(urlString string) (Request, error) {
	parsed, err := url.Parse(urlString)
	if err != nil {
		return Request{}, err
	}
	return Request{URL: parsed}, nil
}

// Call returns body, status, time, error
func (req Request) Call() *Response {
	pathWithQuery, err := makepathWithQuery(req.URL)
	if err != nil {
		return makeErrResponse(err)
	}
	host := req.URL.Host
	var buf bytes.Buffer
	start := time.Now()
	conn, err := net.Dial("tcp", host+":80")
	if err != nil {
		return makeErrResponse(err)
	}
	defer conn.Close()
	reqstr := getRequestString(host, pathWithQuery)
	fmt.Fprintf(conn, reqstr)
	io.Copy(&buf, conn)
	end := time.Now()
	res := parseResponse(buf.String())
	res.Time = end.Sub(start)
	return res
}
