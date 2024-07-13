package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

type kvStore struct {
	m  map[string]string
	mu sync.Mutex
}

func main() {

	var router = mux.NewRouter()
	var kv = kvStore{
		m: make(map[string]string),
	}

	router.HandleFunc("/kv/{param1}", kv.getKV).Methods("GET")
	router.HandleFunc("/kv/{param1}/{param2}", kv.putKV).Methods("PUT")
	router.HandleFunc("/kv/{param1}/{param2}", kv.postKV).Methods("POST")
	router.HandleFunc("/kv/{param1}", kv.deleteKV).Methods("DELETE")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}

func (kv *kvStore) deleteKV(w http.ResponseWriter, r *http.Request) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	params := mux.Vars(r)
	key := params["param1"]
	fmt.Println("this is the deleted key and value = ", key, kv.m[key])
	delete(kv.m, key)

}

func (kv *kvStore) postKV(w http.ResponseWriter, r *http.Request) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	params := mux.Vars(r)
	key := params["param1"]
	value := params["param2"]
	fmt.Println("this is the request = ", key, value)
	kv.m[key] = value
	fmt.Println("this is the updated key and value = ", key, kv.m[key])
}

func (kv *kvStore) putKV(w http.ResponseWriter, r *http.Request) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	params := mux.Vars(r)
	key := params["param1"]
	value := params["param2"]
	kv.m[key] = value
	fmt.Println("this is the new key and value = ", key, kv.m[key])
}

func (kv *kvStore) getKV(w http.ResponseWriter, r *http.Request) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	params := mux.Vars(r)
	key := params["param1"]
	fmt.Println("this is the key and value = ", key, kv.m[key])
}
