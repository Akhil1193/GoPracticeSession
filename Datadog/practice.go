package main

import (
	"github.com/go-chi/chi"
	chitracer "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-chi/chi"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
)

func main() {
	tracer.Start(
		tracer.WithServiceName("go-agent-service"),
		tracer.WithEnv("staging"),
	)

	defer tracer.Stop()

	// build chi router
	router := chi.NewRouter()
	router.Use(chitracer.Middleware(chitracer.WithServiceName("go-agent-service")))
	router.Get("/", MyRequestHandler)
	http.ListenAndServe(":8080", router)
}
func MyRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Chi's all good"))
}