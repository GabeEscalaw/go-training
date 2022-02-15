package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// ContactInfo is a struct that contains the information recorded for each contact
type ContactInfo struct {
	ID   int			`json:"id"`
	Last string			`json:"last"`
	First string		`json:"first"`
	Company string		`json:"company"`
	Address string		`json:"address"`
	Country string		`json:"country"`
	Position string		`json:"position"`
}

// Database is a struct that takes care of appending ID as well as the mutex and slice of ContactInfos
type Database struct {
	NextID int        `json:"nextid"`
	mu     sync.Mutex // initialized in lock state
	Contacts   []ContactInfo   `json:"contacts"`
}

func readFile(fileName string) Database {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open json file: %s", err)
	}
	defer file.Close()
	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read json file: %s", err)
	}
	var result Database
	json.Unmarshal([]byte(byteData), &result)
	
	return result
}

func writeFile(result Database) {
	data := result
	byteData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("failed to Marshal data: %s", err)
	}
	ioutil.WriteFile("output.json", byteData, 0644)
}

func main() {
	writeFile(readFile("users.json"))
}
// func readDatabase () int {
// 	var fileRead map[string] interface{}
// 	_ = json.Unmarshal([]byte(byteData), &result)

// 	id := result["nextID"].(string)
// 	contacts := result["conts"].([]interface{})
// 	var ct []Database
// 	for _, item := range contacts {
// 		ctemp := item.(map[string]interface{})
// 		contact := Contact{ID: ctemp["ID"].(string), Last: ctemp["Last"].(string), First: ctemp["First"].(string), Company: ctemp["Company"].(string), Address: ctemp["Address"].(string), Country: ctemp["Country"].(string), Position: ctemp["Position"].(string)}
// 		ct = append(ct, contact)
// 	}

// 	idInt, _ := strconv.Atoi(id)
// 	fmt.Println(id)
// 	fmt.Println(contacts)

// 	return idInt, ct
// }

// main initializes an empty database: contacts that uses the struct ContactInfo
// func main() {
	
// 	db := &Database{contacts: []ContactInfo{}} // Database starts wtih an empty slice of records
// 	http.ListenAndServe(":8080", db.handler())
// }

// handler provides the response writer and requests for the process function if the URL path is correct.
// func (db *Database) handler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var id int
// 		if r.URL.Path == "/contacts" {
// 			db.process(w, r)
// 		} else if n, _ := fmt.Sscanf(r.URL.Path, "/contacts/%d", &id); n == 1 {
// 			db.processID(id, w, r)
// 		} else {
// 			message := "404. That's an error.\nThe requested URL " + r.URL.Path + " was not found on this server. That's all we know."
// 			http.Error(w, message, http.StatusBadRequest)
// 		}
// 	}
// }

// // process processes the switch cases for POST and GET when URL path is /contacts
// func (db *Database) process(w http.ResponseWriter, r *http.Request) {
// 	var ci ContactInfo
// 	switch r.Method {
// 	case "POST":
// 		w.Header().Set("Content-Type", "application/json")
		
// 		if err := json.NewDecoder(r.Body).Decode(&ci); err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		isDuplicate := false

// 		for _, item := range db.contacts {
// 			if ci.First == item.First && ci.Last == item.Last {
// 				isDuplicate = true
// 				http.Error(w, "Bad Request: Error 409. Contact already exists in Database", http.StatusConflict)
// 			}
// 		}

// 		if isDuplicate == false {
// 			http.Error(w, "Contact successfully added to database", http.StatusCreated)
// 			db.mu.Lock()
// 			ci.ID = db.nextID
// 			db.nextID++
// 			db.contacts = append(db.contacts, ci)
// 			db.mu.Unlock()
// 		}
		
// 	case "GET":
// 		w.Header().Set("Content-Type", "application/json")
		
// 		if len(db.contacts) == 0 {
// 			http.Error(w, "Bad Request: Error 404. Database not found.", http.StatusNotFound)
// 		} else {
// 			if err := json.NewEncoder(w).Encode(db.contacts); err != nil {
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 				return
// 			}
// 		}
		
// 	default:
// 		w.Header().Set("Content-Type", "application/json")

// 		http.Error(w, "Bad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
// 	}
// }

// // processID processes the switch cases for GET, DELETE, and PUT when URL path is /contacts/{id}
// func (db *Database) processID(id int, w http.ResponseWriter, r *http.Request) {
// 	var ci ContactInfo
// 	exists := false
// 	switch r.Method {
// 	case "GET":
// 		w.Header().Set("Content-Type", "application/json")
		
// 		for j, item := range db.contacts {
// 			if id == item.ID {
// 				if err := json.NewEncoder(w).Encode(db.contacts[j]); err != nil {
// 					http.Error(w, err.Error(), http.StatusInternalServerError)
// 					return
// 				}
// 				exists = true
// 				break
// 			}
// 		}

// 		if len(db.contacts) == 0 {
// 			http.Error(w, "Bad Request: Error 404. Database not found.", http.StatusNotFound)
// 		} else if !exists {
// 			http.Error(w, "Bad Request: Error 404. Item was not found.", http.StatusNotFound)
// 		}
		
// 	case "DELETE":
// 		db.mu.Lock()
// 		for j, item := range db.contacts {
// 			if id == item.ID {
// 				fmt.Fprintf(w, "Sucessfully Deleted Contact# %v\n", id)
// 				db.contacts = append(db.contacts[:j], db.contacts[j+1:]...)
// 				exists = true
// 				break
// 			}
// 		}
// 		db.mu.Unlock()
// 		w.Header().Set("Content-Type", "application/json")
// 		if !exists {
// 			http.Error(w, "Bad Request: Error 404. Item not found.", http.StatusNotFound)
// 		}
// 	case "PUT":
	
// 		if err := json.NewDecoder(r.Body).Decode(&ci); err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		db.mu.Lock()
// 		for j, item := range db.contacts {
// 			if id == item.ID {
// 				fmt.Fprintf(w, "Sucessfully Replaced Contact# %v\n", id)
// 				tempID := id
// 				db.contacts[j] = ci
// 				db.contacts[j].ID = tempID
// 				exists = true
// 				break
// 			}
// 		}
// 		db.mu.Unlock()
// 		w.Header().Set("Content-Type", "application/json")
// 		if !exists {
// 			http.Error(w, "Bad Request: Error 404. Item to replace was not found.", http.StatusNotFound)
// 		}

// 	default:
// 		w.Header().Set("Content-Type", "application/json")

// 		http.Error(w, "Bad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
// 	}
// } 
