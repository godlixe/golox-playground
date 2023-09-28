package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/rs/cors"
)

type Request struct {
	Code string
}

func main() {
	def := http.NewServeMux()

	def.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			var rq Request
			err := json.NewDecoder(r.Body).Decode(&rq)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
			cmd := exec.CommandContext(ctx, "./golox")
			var outb, errb bytes.Buffer

			cmd.Stdout = &outb
			cmd.Stderr = &errb
			cmd.Stdin = strings.NewReader(rq.Code)

			go func() {
				cmd.Run()
				cancel()
			}()

			// var timeout string = ctx.Err().Error() + "\n"
			// for {
			w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5501")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, hx-current-url, hx-request")
			w.Header().Set("Access-Control-Allow-Methods", "POST`")
			w.Header().Set("Content-Type", "text/html")
			<-ctx.Done()
			switch ctx.Err() {
			case context.DeadlineExceeded:
				fmt.Fprint(w, "Time limit exceeded.\n"+outb.String())
			case context.Canceled:
				fmt.Fprint(w, outb.String())
			}
		case "OPTIONS":
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, hx-current-url, hx-request")
			w.Header().Set("Access-Control-Allow-Methods", "POST`")
			w.Header().Set("Access-Control-Max-Age", "86400")
			fmt.Fprint(w, "hi")
		}

		w.Header().Set("Content-Type", "application/json")
		// w.Write([]byte("{\"hello\": \"world\"}"))
	})

	handler := cors.Default().Handler(def)
	http.ListenAndServe(":8080", handler)
}
