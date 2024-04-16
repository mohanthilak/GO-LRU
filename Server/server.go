package server

import (
	"context"
	app "go-lru/App"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type HTTPServer struct {
	router *mux.Router
	port   int
	app    *app.App
}

func NewHTTPServer(port int, app *app.App) *HTTPServer {
	return &HTTPServer{router: mux.NewRouter(), port: port, app: app}
}

func (S *HTTPServer) initiateRoutes() {
	S.router.Use(loggingMiddleware)
	S.router.HandleFunc("/get/{key}", S.HandleGetValueFromKey).Methods("GET")
	S.router.HandleFunc("/set", S.HandleSetKeyValuePair).Methods("POST")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (S *HTTPServer) StartServer() {
	S.initiateRoutes()

	server := http.Server{
		Addr:         "0.0.0.0:" + strconv.Itoa(S.port),
		Handler:      handlers.CORS(handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}), handlers.AllowCredentials(), handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "X-Requested-With", "Authorization", "Content-Type"}), handlers.AllowedOrigins([]string{"http://localhost:5173"}))(S.router),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Println("server starting at port: ", S.port)
		err := server.ListenAndServe()
		if err != nil {
			log.Panicf("error while serving the server. %s", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c

	log.Printf("Gracefully Ending the Server after receiving signal %s", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
