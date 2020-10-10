package main

import (
	"encoding/json"
	"net/http"
)

func homeHandler(rw http.ResponseWriter, rq *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.Write([]byte(`
<!DOCTYPE html>
<html lang="en">
	<head>
		<style>
			html {
				font-size: 16px;
				font-weight: 500;
				font-family: 'Microsoft Yahei', sans-serif, monospace;
				color: #2c2c2c;
			}
		</style>
	</head>
	<body>
		<h1>WELLCOME!</h1>
		<ul>
			<li><a href="http://localhost:3000/list">list</a></li>
			<li><a href="http://xhdev.docimaxvip.com">xhdev</a></li>
		</ul>
	</body>
</html>
	`))
}

func listHandler(rw http.ResponseWriter, rq *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	data := make(map[string]string, 0)
	data["id"] = "ID00001"
	data["name"] = "bey"
	data["gender"] = "male"

	res := response{
		Code: 200,
		Msg:  "success!",
		Data: data,
	}

	json, err := json.Marshal(res)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Write(json)
}
