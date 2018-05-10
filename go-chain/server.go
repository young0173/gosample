package go_chain

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

type message struct {
	BPM int
}

var BlockChain []Block = []Block{{0, "2018-04-08 14:43:52.798726491 +0800 CST m=+12.318602538", 0, "9a271f2a916b0b6ee6cecb2426f0b3206ef074578be55d9bc94f6f3fe3ab86aa", ""}}

func handleGetBlockChain(w http.ResponseWriter, r *http.Request) {

	bytes, err := json.MarshalIndent(BlockChain, "", " ")
	if err != nil {
		glog.Error("Get BlockChain failed:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(bytes))

}

func responseWithJson(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {

	res, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal server error."))
		return
	}
	w.WriteHeader(code)
	w.Write(res)

}

func handleWriteBlockChain(w http.ResponseWriter, r *http.Request) {

	var msg message
	var block Block

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&msg); err != nil {
		responseWithJson(w, r, http.StatusInternalServerError, r.Body)
		return
	}
	defer r.Body.Close()

	if len(BlockChain) >= 1 {
		block.New(BlockChain[len(BlockChain)-1], msg.BPM)
		if block.IsValid(BlockChain[len(BlockChain)-1]) {
			BlockChain = append(BlockChain, block)
		}
		responseWithJson(w, r, http.StatusOK, block)
	} else {
		glog.Error("Init block chain not exist.")
		responseWithJson(w, r, http.StatusInternalServerError, r.Body)
		return
	}

}

func NewMuxRouter() http.Handler {

	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockChain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlockChain).Methods("POST")

	return muxRouter

}

func RUN() error {

	router := NewMuxRouter()
	glog.Info("Server is starting, listening on :8080")

	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		glog.Fatal("Server exited:", err)
		return err
	}

	return nil

}
