package routes

import (
	"fmt"
	"net/http"

	"bey.com/models"
)

// HomeHandler home page handler
func HomeHandler(rw http.ResponseWriter, rq *http.Request) {
	c := models.Circle{Radius: 5.12}
	rw.Header().Set("Content-Type", "text/html")
	fmt.Println(c.Area())
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
