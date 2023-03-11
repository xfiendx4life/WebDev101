package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	logInfo  *log.Logger
	logError *log.Logger
)

func init() {
	logInfo = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logError = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
}

type queue struct {
	name string
	data []string
	mx   sync.Mutex
}

type httpHandler struct {
	storage sync.Map
	context context.Context
}

func (h *httpHandler) handlePut(w http.ResponseWriter, r *http.Request, qName string) {
	v := r.URL.Query().Get("v")
	if v == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if q, ok := h.storage.Load(qName); ok {
		if qu, ok := q.(*queue); ok {
			if err := qu.push(h.context, v); err != nil {
				logError.Printf("can't add to queue %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	} else {
		logInfo.Printf("new queue %s", qName)
		q := newQueue(qName)
		logInfo.Printf("new queue %v created", q)
		q.push(h.context, v)
		h.storage.Store(qName, q)
	}
	logInfo.Printf("message %s added to queue %s", v, qName)
	w.WriteHeader(http.StatusOK)
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	qName := strings.TrimPrefix(r.URL.Path, "/")
	if qName == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.Method == "PUT" {
		h.handlePut(w, r, qName)
	} else if r.Method == "GET" {
		logInfo.Printf("Query to queue %s", qName)
		logInfo.Printf("Storage %#v", h.storage)
		if res, ok := h.storage.Load(qName); ok {
			w.Write([]byte(fmt.Sprintf("Get queue %v", res)))
			logInfo.Println(res)
		}
	}
}

func newQueue(name string) *queue {
	return &queue{
		name: name,
		data: make([]string, 0, 10),
	}
}

func (q *queue) push(ctx context.Context, message string) error {
	select {
	case <-ctx.Done():
		logError.Println("push stopped by context")
		return fmt.Errorf("context done")
	default:
		q.mx.Lock()
		defer q.mx.Unlock()
		q.data = append(q.data, message)
		logInfo.Printf("message: %s, put to queue: %s", message, q.name)
		return nil
	}
}

func (q *queue) pop(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		logError.Println("pop stopped by context")
		return "", fmt.Errorf("context done")
	default:
		q.mx.Lock()
		defer q.mx.Unlock()
		if len(q.data) == 0 {
			logError.Printf("queue %s is empty", q.name)
			return "", fmt.Errorf("queue %s is empty", q.name)
		}
		res := q.data[0]
		q.data = q.data[1:]
		logInfo.Printf("get message %s from queue %s", res, q.name)
		return res, nil
	}
}

func main() {
	port := os.Getenv("PORT")
	storage := sync.Map{}
	handler := httpHandler{
		storage: storage,
		context: context.TODO(),
	}
	logInfo.Printf("starting server at %s", port)
	http.Handle("/", &handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%s", port), nil))
}
