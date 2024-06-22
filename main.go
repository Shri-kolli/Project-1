package main

import (
	"fmt"
	"net/http"
	"time"
)

// handler function to handle HTTP requests
func handler(w http.ResponseWriter, r *http.Request) {
	response := `
		<html>
		<head>
			<title>India Time</title>
			<script type="text/javascript">
				function updateTime() {
					var xhr = new XMLHttpRequest();
					xhr.onreadystatechange = function() {
						if (xhr.readyState == 4 && xhr.status == 200) {
							document.getElementById("time").innerHTML = xhr.responseText;
						}
					}
					xhr.open("GET", "/time", true);
					xhr.send();
				}
				setInterval(updateTime, 1000);
				window.onload = updateTime;
			</script>
		</head>
		<body>
			<h1>India Time</h1>
			<p id="time"></p>
		</body>
		</html>
	`
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

// timeHandler function to handle time requests
func timeHandler(w http.ResponseWriter, r *http.Request) {
	location, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	currentTime := time.Now().In(location)
	w.Write([]byte(currentTime.Format("Monday, 02-Jan-06 15:04:05 MST")))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/time", timeHandler)
	fmt.Println("Starting server on :80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
