package main

import(
	"net/http"
	"io"
)

func hello(res http.ResponseWriter, req*http.Request){
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<!Doctype html>
		<html lang="en" dir="ltr">
			<head>
				<title>Hello World</title>
			</head>
			<body>
				Hello World!
			</body>
		</html>`,
	)
}

func main(){
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)
}