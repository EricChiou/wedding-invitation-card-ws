package router

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var p = 8080
var h http.Handler
var get, post, put, delete, patch, copy, head, options *pathElement

type pathElement struct {
	path []string
	run  *func(*Context)
	prev *pathElement
	next *pathElement
}

// Context is use to pass variables between middleware
type Context struct {
	Req    *http.Request
	Res    http.ResponseWriter
	Params map[string]string
}

// INIT is use to set server port, handler
func INIT(port int, handler http.Handler) {
	p = port
	h = handler
}

// RUN server, pass server port you want to listen, default is 8080
func RUN() {
	router()
	fmt.Println("start server at port " + strconv.Itoa(p))
	if err := http.ListenAndServe(":"+strconv.Itoa(p), h); err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("start server error: ", err)
		}
	}()
}

// GET is use to build new GET API
func GET(path string, run func(*Context)) {
	if !checkPath("GET", path, get) {
		return
	}
	element := pathElement{path: strings.Split(path[1:], "/"), run: &run, prev: nil, next: get}
	if get != nil {
		get.prev = &element
	}
	get = &element
}

// POST is use to build new POST API
func POST(path string, run func(*Context)) {
	if !checkPath("POST", path, post) {
		return
	}
	element := pathElement{path: strings.Split(path[1:], "/"), run: &run, prev: nil, next: post}
	if post != nil {
		post.prev = &element
	}
	post = &element
}

// PUT is use to build new PUT API
func PUT(path string, run func(*Context)) {
	if !checkPath("PUT", path, put) {
		return
	}
	element := pathElement{path: strings.Split(path[1:], "/"), run: &run, prev: nil, next: put}
	if put != nil {
		put.prev = &element
	}
	put = &element
}

// DELETE is use to build new DELETE API
func DELETE(path string, run func(*Context)) {
	if !checkPath("DELETE", path, delete) {
		return
	}
	element := pathElement{path: strings.Split(path[1:], "/"), run: &run, prev: nil, next: delete}
	if delete != nil {
		delete.prev = &element
	}
	delete = &element
}

// PATCH is use to build new PATCH API
func PATCH(path string, run func(*Context)) {
	if !checkPath("PATCH", path, patch) {
		return
	}
	element := pathElement{path: strings.Split(path[1:], "/"), run: &run, prev: nil, next: patch}
	if patch != nil {
		patch.prev = &element
	}
	patch = &element
}

// COPY is use to build new COPY API
func COPY(path string, run func(*Context)) {
	if !checkPath("COPY", path, copy) {
		return
	}
	element := pathElement{path: strings.Split(path[1:], "/"), run: &run, prev: nil, next: copy}
	if copy != nil {
		copy.prev = &element
	}
	copy = &element
}

// HEAD is use to build new HEAD API
func HEAD(path string, run func(*Context)) {
	if !checkPath("HEAD", path, head) {
		return
	}
	element := pathElement{path: strings.Split(path[1:], "/"), run: &run, prev: nil, next: head}
	if head != nil {
		head.prev = &element
	}
	head = &element
}

// OPTIONS is use to build new OPTIONS API
func OPTIONS(path string, run func(*Context)) {
	if !checkPath("OPTIONS", path, options) {
		return
	}
	element := pathElement{path: strings.Split(path[1:], "/"), run: &run, prev: nil, next: options}
	if options != nil {
		options.prev = &element
	}
	options = &element
}

func portHandler(port []string) string {
	switch len(port) {
	case 1:
		return port[0]
	default:
		if port := os.Getenv("PORT"); port != "" {
			return port
		}
		return "8080"
	}
}

func checkPath(method string, path string, pathList *pathElement) bool {
	// check path is valid
	if len(path) == 0 || path[0:1] != "/" {
		fmt.Println("wrong path at " + method + ": '" + path + "', the first character must be the '/'.")
		return false
	}
	if strings.Contains(path, "?") || strings.Contains(path, "&") {
		fmt.Println("wrong path at " + method + ": '" + path + "', the path has invalid character.")
		return false
	}

	pathAry := strings.Split(path[1:], "/")
	// check path does not has wrong format
	for _, p := range pathAry {
		if p == "" {
			fmt.Println("wrong path at " + method + ": '" + path + "', the path has wrong format.")
			return false
		}
	}

	// check path if duplicate
	element := pathList
	for element != nil {
		if checkDuplicate(pathAry, element.path) {
			fmt.Println("path duplicate, " + method + ": " + path)
			return false
		}
		element = element.next
	}
	return true
}

func checkDuplicate(pathAry []string, targetPathAry []string) bool {
	if len(pathAry) == len(targetPathAry) {
		for i := 0; i < len(pathAry); i++ {
			if targetPathAry[i][0:1] != ":" {
				if pathAry[i] != targetPathAry[i] {
					return false
				}
			}
		}
		return true
	}
	return false
}

func router() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			pathHandler(res, req, get)
		case "POST":
			pathHandler(res, req, post)
		case "PUT":
			pathHandler(res, req, put)
		case "DELETE":
			pathHandler(res, req, delete)
		case "PATCH":
			pathHandler(res, req, patch)
		case "COPY":
			pathHandler(res, req, copy)
		case "HEAD":
			pathHandler(res, req, head)
		case "OPTIONS":
			pathHandler(res, req, options)
		default:
			res.Write([]byte("404 page not found"))
		}
	})
}

func pathHandler(res http.ResponseWriter, req *http.Request, pathList *pathElement) {
	element := pathList
	for element != nil {
		match, params := mapping(req.URL.Path, element.path)
		if match {
			run := *element.run
			run(&Context{Req: req, Res: res, Params: params})
			return
		}
		element = element.next
	}
	res.Write([]byte("404 page not found"))
}

func mapping(path string, targetPathAry []string) (bool, map[string]string) {
	pathAry := strings.Split(path[1:], "/")
	params := make(map[string]string)
	if len(pathAry) != len(targetPathAry) {
		return false, nil
	}
	for i := 0; i < len(pathAry); i++ {
		if targetPathAry[i][0:1] == ":" {
			params[targetPathAry[i][1:]] = pathAry[i]
		} else {
			if pathAry[i] != targetPathAry[i] {
				return false, nil
			}
		}
	}
	return true, params
}
