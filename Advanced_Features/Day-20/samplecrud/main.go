package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items []Item

func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	items = append(items, newItem)
	w.WriteHeader(http.StatusCreated)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	id := 1 // Assuming ID to update is 1, you can change this as needed
	var updatedItem Item
	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	items[id-1] = updatedItem
	w.WriteHeader(http.StatusOK)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	id := 1 // Assuming ID to delete is 1, you can change this as needed
	items = append(items[:id-1], items[id:]...)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/items", getItems)
	http.HandleFunc("/items/create", createItem)
	http.HandleFunc("/items/update", updateItem)
	http.HandleFunc("/items/delete", deleteItem)

	fmt.Println("Server is starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
