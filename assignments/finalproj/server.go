package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// ContactInfo is a struct that contains the information recorded for each contact
type WatchInfo struct {
	Brand string
	Model string
	DialSize string
	Price string
	Collected bool
}

type UserInfo struct {
	Username string
	Password string
	IsLoggedIn bool

	mu sync.Mutex
	Watches []WatchInfo
}

type UserDatabase struct {
	mu 	sync.Mutex
	users []UserInfo
}

// main initializes an empty database: contacts that uses the struct ContactInfo
func main() {
	u := &UserDatabase{users: []UserInfo{}} // Users start with an empty collection of watches 
	http.ListenAndServe(":8080", u.handler())
}

// handler provides the response writer and requests for the process function if the URL path is correct.
func (u *UserDatabase) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/signup" {
			u.signup(w, r)
		} else if r.URL.Path == "/user" {
			u.login(w, r)
		} else if r.URL.Path == "/logout" {
			u.logout(w, r)
		} else if r.URL.Path == "/delete" {
			u.delete(w, r)

		} else if r.URL.Path == "/add" {
			u.addWatch(w, r)
		} else if r.URL.Path == "/remove" {
			u.removeWatch(w, r)
		} else if r.URL.Path == "/mark" {
			u.markWatch(w, r)
		} else if r.URL.Path == "/list" {
			u.watchList(w, r)
		} else {
			message := "Bad Request: Error 404. url " + r.URL.Path + " was not found on this server."
			http.Error(w, message, http.StatusBadRequest)
		}
	}
}

// process processes the switch cases for POST and GET when URL path is /contacts
func (u *UserDatabase) signup(w http.ResponseWriter, r *http.Request) {
	var ui UserInfo
	switch r.Method {
	case "POST":
		w.Header().Set("Content-Type", "application/json")
		
		if err := json.NewDecoder(r.Body).Decode(&ui); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		isDuplicate := false

		for _, user := range u.users {
			if ui.Username == user.Username {
				isDuplicate = true
				http.Error(w, "Bad Request: Error 409. User already exists in Database", http.StatusConflict)
			} 
		}

		if !isDuplicate {
			u.mu.Lock()
			u.users = append(u.users, ui)
			u.mu.Unlock()
			http.Error(w, "##########\nUser successfully added to database.\n##########\n", http.StatusCreated)
		}
		
	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "Bad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
	}
}

func (u *UserDatabase) login(w http.ResponseWriter, r *http.Request) {
	var ui UserInfo
	isLogged := false
	amtLogged := 0
	userCorrect := false
	passCorrect := false
	index := -1
	alreadyLogged := false

	switch r.Method {
	case "GET": 
		w.Header().Set("Content-Type", "application/json")
		
		if len(u.users) == 0 {
			http.Error(w, "Bad Request: Error 404. Database not found.", http.StatusNotFound)
		} else {
			if err := json.NewEncoder(w).Encode(u.users); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	case "POST":
		w.Header().Set("Content-Type", "application/json")
		
		if err := json.NewDecoder(r.Body).Decode(&ui); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		for _, user := range u.users {
			if user.IsLoggedIn {
				amtLogged++
			}
		}

		u.mu.Lock()
		for j, user := range u.users {
			if ui.Username == user.Username && ui.Password == user.Password && amtLogged == 0 {
				u.users[j].IsLoggedIn = true
				isLogged = true
				index = j
				userCorrect = true
				passCorrect = true
				amtLogged++
				//fmt.Fprintln(w, "Successfully logged in ran")
			} else if ui.Username == user.Username && ui.Password == user.Password && user.IsLoggedIn && amtLogged == 1 {
				isLogged = true
				index = j
				userCorrect = true
				passCorrect = true
				alreadyLogged = true
			} else if ui.Username == user.Username && ui.Password != user.Password {
				userCorrect = true
				//fmt.Fprintln(w, "Correct user, but wrong pass ran")
			} else if ui.Username == user.Username && ui.Password == user.Password && amtLogged == 1 {
				amtLogged++
			} 
		}
		u.mu.Unlock()
		
		if alreadyLogged {
			fmt.Fprintf(w, "####################\nYou are already logged in, %v\n####################\n",  u.users[index].Username)
		} else if isLogged && amtLogged == 1 && userCorrect && passCorrect {
			fmt.Fprintf(w, "####################\nLogged in successfully.\nWelcome to your Watch Collection, %v.\n####################\n", u.users[index].Username)
		} else if !isLogged && amtLogged == 0 && userCorrect {
				fmt.Fprintln(w, "####################\nThe password you entered is incorrect. Please try again.\n####################")
		} else if !isLogged && amtLogged == 0 && !userCorrect && !passCorrect {
			fmt.Fprintln(w, "####################\nAccount does not exist.\n####################")
		} else {
			fmt.Fprintln(w, "####################\nPlease log out of your current account before attempting to log in again.\n####################")
		}
		//fmt.Fprintf(w, "userCorrect: %v, passCorrect: %v, amtLogged: %v\n", userCorrect, passCorrect, amtLogged)
		
	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "Bad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
	}
}

func (u *UserDatabase) logout(w http.ResponseWriter, r *http.Request) {
	var ui UserInfo
	
	switch r.Method {
	case "GET": 
		w.Header().Set("Content-Type", "application/json")
		
		if len(u.users) == 0 {
			http.Error(w, "Bad Request: Error 404. Database not found.", http.StatusNotFound)
		} else {
			if err := json.NewEncoder(w).Encode(u.users); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}	

	case "POST":
		w.Header().Set("Content-Type", "application/json")
		
		allOut := false
		if err := json.NewDecoder(r.Body).Decode(&ui); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for j, _ := range u.users {
			if u.users[j].IsLoggedIn {
				u.users[j].IsLoggedIn = false
				fmt.Fprintln(w, "####################\nSuccessfully logged out.\nThank you for using WatchCollection.\n####################")
				allOut = true
				break
			} 
		}

		if !allOut {
			fmt.Fprintln(w, "####################\nYou are not logged in.\n####################")
		}
		
	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "Bad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
	}
}

func (u *UserDatabase) delete(w http.ResponseWriter, r *http.Request) {
	var ui UserInfo
	switch r.Method {
	case "GET": 
		w.Header().Set("Content-Type", "application/json")
		
		if len(u.users) == 0 {
			http.Error(w, "Bad Request: Error 404. Database not found.", http.StatusNotFound)
		} else {
			if err := json.NewEncoder(w).Encode(u.users); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}	

	case "POST":
		w.Header().Set("Content-Type", "application/json")
		
		if err := json.NewDecoder(r.Body).Decode(&ui); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for j, user := range u.users {
			if ui.Username == user.Username && ui.Password == user.Password {
				fmt.Fprintf(w, "####################\nThank you for trying our app, %v.\nYour account has been successfully deleted.\n####################\n", u.users[j].Username)
				u.users = append(u.users[:j], u.users[j+1:]...)
				break
			} 
		}
		
	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "Bad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
	}
}

func (u *UserDatabase) watchList (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")

		if len(u.users) == 0 {
			http.Error(w, "Bad Request: Error 404. Database not found.", http.StatusNotFound)
		} else {

			index := -1

			for i, user := range u.users {
				if user.IsLoggedIn {
					index = i
				}
			}

			for j, user := range u.users {
				if user.IsLoggedIn {
					if err := json.NewEncoder(w).Encode(u.users[j].Watches); err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
				}
			}
	
			if len(u.users[index].Watches) > 0 {
				fmt.Fprintln(w, "")
				fmt.Fprintln(w, "~~~~~ YOUR CURRENT WATCH COLLECTION ~~~~~~")
				for _, watch := range u.users[index].Watches {
					fmt.Fprintln(w, "")
					fmt.Fprintln(w, "----------")
					fmt.Fprintln(w, "")
					fmt.Fprintf(w, "-| %v %v |-\n", watch.Brand, watch.Model )
					fmt.Fprintf(w, "     Dial Size: %v     \n", watch.DialSize )
					fmt.Fprintf(w, "     %v     \n", watch.Price )
					fmt.Fprintln(w, "")
					fmt.Fprintln(w, "----------")
					fmt.Fprintln(w, "")
				}
			} else {
				fmt.Fprintln(w, "Your collection is currently empty.")
			}
		}
		
	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "Bad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
	}
}

func (u *UserDatabase) addWatch(w http.ResponseWriter, r *http.Request) {
	var wi WatchInfo
	switch r.Method {
	case "POST":
		w.Header().Set("Content-Type", "application/json")
		
		index := -1

		if err := json.NewDecoder(r.Body).Decode(&wi); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		for j, user := range u.users {
			if user.IsLoggedIn {
				index = j
				u.users[j].Watches = append(u.users[j].Watches, wi)
				fmt.Fprintln(w, "Sucessfully added watch to the collection.")
			}
		}

		if index == -1 {
			fmt.Fprintln(w, "Please log in before trying to add a watch from your list.")
		}

		if len(u.users[index].Watches) > 0 {
			fmt.Fprintln(w, "")
			fmt.Fprintln(w, "~~~~~ YOUR CURRENT WATCH COLLECTION ~~~~~~")
			for _, watch := range u.users[index].Watches {
				fmt.Fprintln(w, "")
				fmt.Fprintln(w, "----------")
				fmt.Fprintln(w, "")
				fmt.Fprintf(w, "-| %v %v |-\n", watch.Brand, watch.Model )
				fmt.Fprintf(w, "     Dial Size: %v     \n", watch.DialSize )
				fmt.Fprintf(w, "     %v     \n", watch.Price )
				fmt.Fprintln(w, "")
				fmt.Fprintln(w, "----------")
				fmt.Fprintln(w, "")
			}
		} else {
			fmt.Fprintln(w, "Your collection is currently empty.")
		}
		
	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "Bad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
	}
}

func (u *UserDatabase) removeWatch(w http.ResponseWriter, r *http.Request) {
	var wi WatchInfo
	switch r.Method {
	case "POST":
		w.Header().Set("Content-Type", "application/json")
		
		if err := json.NewDecoder(r.Body).Decode(&wi); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		index := -1
		for i, user := range u.users {
			if user.IsLoggedIn {
				index = i
			}
		}

		for j, watch := range u.users[index].Watches {
			if wi.Brand == u.users[index].Watches[j].Brand && wi.Model == u.users[index].Watches[j].Model {
				u.users[index].Watches = append(u.users[index].Watches[:j], u.users[index].Watches[j+1:]...)
				fmt.Fprintf(w, "%v %v successfully removed from your list.\n", watch.Model, watch.Brand)			
				break
			}
		}

		if index == -1 {
			fmt.Fprintln(w, "Please log in before trying to remove a watch from your list.\n")
		}

		if len(u.users[index].Watches) > 0 {
			fmt.Fprintln(w, "")
			fmt.Fprintln(w, "~~~~~ YOUR CURRENT WATCH COLLECTION ~~~~~~")
			for _, watch := range u.users[index].Watches {
				fmt.Fprintln(w, "")
				fmt.Fprintln(w, "----------")
				fmt.Fprintln(w, "")
				fmt.Fprintf(w, "-| %v %v |-\n", watch.Brand, watch.Model )
				fmt.Fprintf(w, "     Dial Size: %v     \n", watch.DialSize )
				fmt.Fprintf(w, "     %v     \n", watch.Price )
				fmt.Fprintln(w, "")
				fmt.Fprintln(w, "----------")
				fmt.Fprintln(w, "")
			}
		} else {
			fmt.Fprintln(w, "Your collection is currently empty.")
		}

	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "Bad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
	}
}

func (u *UserDatabase) markWatch(w http.ResponseWriter, r *http.Request) {
	var wi WatchInfo
	switch r.Method {
	case "POST":
		w.Header().Set("Content-Type", "application/json")
		
		if err := json.NewDecoder(r.Body).Decode(&wi); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		index := -1
		for i, user := range u.users {
			if user.IsLoggedIn {
				index = i
			}
		}

		for j, watch := range u.users[index].Watches {
			if wi.Brand == u.users[index].Watches[j].Brand && wi.Model == u.users[index].Watches[j].Model {
				if !u.users[index].Watches[j].Collected {
					u.users[index].Watches[j].Collected = true
					fmt.Fprintf(w, "%v %v successfully marked as COLLECTED in your list.\n", watch.Model, watch.Brand)
					break
				} else if u.users[index].Watches[j].Collected {
					u.users[index].Watches[j].Collected = false
					fmt.Fprintf(w, "%v %v successfully marked as NOT YET COLLECTED in your list.\n", watch.Model, watch.Brand)
					break
				}				
			}
		}

		if index == -1 {
			fmt.Fprintln(w, "Please log in before trying to mark a watch from your list.")
		}

		if len(u.users[index].Watches) > 0 {
			fmt.Fprintln(w, "")
			fmt.Fprintln(w, "~~~~~ YOUR CURRENT WATCH COLLECTION ~~~~~~")
			for _, watch := range u.users[index].Watches {
				fmt.Fprintln(w, "")
				fmt.Fprintln(w, "----------")
				fmt.Fprintln(w, "")
				fmt.Fprintf(w, "-| %v %v |-\n", watch.Brand, watch.Model )
				fmt.Fprintf(w, "     Dial Size: %v     \n", watch.DialSize )
				fmt.Fprintf(w, "     %v     \n", watch.Price )
				fmt.Fprintln(w, "")
				fmt.Fprintln(w, "----------")
				fmt.Fprintln(w, "")
			}
		} else {
			fmt.Fprintln(w, "Your collection is currently empty.")
		}

	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "Bad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
	}
}