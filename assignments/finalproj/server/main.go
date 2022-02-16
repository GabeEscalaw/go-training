package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

// WatchInfo sets-up the details that each watch contains.
type WatchInfo struct {
	Brand 		string		`json:"brand"`
	Model 		string		`json:"model"`
	Width 	string			`json:"width"`
	Price 		string		`json:"price"`
	Collected 	bool		`json:"collected"`
}

// UserInfo sets-up the details of each user registered in the database including their own watch collection.
type UserInfo struct {
	Username 	string		`json:"username"`
	Password 	string		`json:"password"`
	IsLoggedIn 	bool		`json:"isLoggedIn"`

	Watches 	[]WatchInfo	`json:"watches"`
}

// UserDatabase contains all of the users registered in the database.
type UserDatabase struct {
	mu 			sync.Mutex  
	Users 		[]UserInfo	`json:"users"`
}

// var sets-up the variables used for init.
var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

// databaseReader reads the users.json file for the database of the program.
func (u *UserDatabase) databaseReader () {
	file, err := os.Open("users.json")
	if err != nil {
		log.Fatalf("%v", err)
		ErrorLogger.Println("databaseReader | Fatal error with opening users.json file.")
	}
	defer file.Close()
	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read json file: %s", err)
		ErrorLogger.Println("databaseReader | Fatal error with reading users.json file.")
	}

	var result UserDatabase
	json.Unmarshal([]byte(byteData), &result)
	*u = result

	//json.Unmarshal([]byte(byteData), &u.Users) 
	InfoLogger.Println("databaseReader | Database was read")
}

// databaseWriter outputs the json file that was read from users.json.
func (u *UserDatabase) databaseWriter() {
	byteData, err := json.Marshal(u)
	if err != nil {
		log.Fatalf("failed to Marshal data: %s", err)
		ErrorLogger.Println("databaseWriter | Fatal error with marshalling result from databaseReader.")
	}
	ioutil.WriteFile("users.json", byteData, 0644)
	InfoLogger.Println("databaseWriter | Database was updated")
}

// init initializes the means to use the Info, Warning, and Error Loggers which are outputted to logs.txt.
func init() {
    file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
		ErrorLogger.Println("init | Fatal error with opening logs.txt file.")
    }

    InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// main initializes an empty database: Users that uses the struct []UserInfo.
func main() {
	u := &UserDatabase{} 
	
	u.databaseReader()

	http.ListenAndServe(":8080", u.handler())
	InfoLogger.Println("main | Server was initialized")
}

// handler provides the response writer and requests for the process function if the URL path is correct based on the -cmd retrieved from watch.go.
func (u *UserDatabase) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/signup" {
			u.signup(w, r)
			InfoLogger.Println("handler | /signup request was received from cli")
		} else if r.URL.Path == "/login" {
			u.login(w, r)
			InfoLogger.Println("handler | /login request was received from cli")
		} else if r.URL.Path == "/logout" {
			u.logout(w, r)
			InfoLogger.Println("handler | /logout request was received from cli")
		} else if r.URL.Path == "/delete" {
			u.delete(w, r)
			InfoLogger.Println("handler | /delete request was received from cli")
		} else if r.URL.Path == "/add" {
			u.addWatch(w, r)
			InfoLogger.Println("handler | /addWatch request was received from cli")
		} else if r.URL.Path == "/remove" {
			u.removeWatch(w, r)
			InfoLogger.Println("handler | /removeWatch request was received from cli")
		} else if r.URL.Path == "/mark" {
			u.markWatch(w, r)
			InfoLogger.Println("handler | /markWatch request was received from cli")
		} else if r.URL.Path == "/list" {
			u.watchList(w, r)
			InfoLogger.Println("handler | /watchList request was received from cli")
		} else {
			message := "\n\n\nBad Request: Error 404. url " + r.URL.Path + " was not found on this server."
			http.Error(w, message, http.StatusBadRequest)
			ErrorLogger.Println(message)
		}
	}
}

// displayCollection sets up the template used to display the username, number of watches collected, and the deatils of each watch in the user's collection.
func displayCollection (u *UserDatabase, watch WatchInfo, username string, numWatches int, index int, w http.ResponseWriter) {
	watchesCollected := 0

	for _, watch := range u.Users[index].Watches {
		if watch.Collected {
			watchesCollected++
		}
	} 
	
	fmt.Fprintln(w, "\n-------------------------------------")
	fmt.Fprintf(w, " User: %v | Watches Collected: %v/%v", username, watchesCollected, numWatches)
	fmt.Fprintln(w, "\n-------------------------------------")
	fmt.Fprintln(w, "\n~~~~~ CURRENT WATCH COLLECTION ~~~~~~")
    fmt.Fprintln(w, "")
	
	for _, watch := range u.Users[index].Watches {
		fmt.Fprintf(w, "|\t%v %v\n", watch.Brand, watch.Model)
		fmt.Fprintf(w, "|\tCase Size: %v\n", watch.Width)	
		fmt.Fprintf(w, "|\t%v\n", watch.Price)
		fmt.Fprintln(w, "|")
		if watch.Collected {
			fmt.Fprintf(w, "|\tCollected.\n")
		} else {
			fmt.Fprintf(w, "|\tNot yet collected.\n")
		}
		fmt.Fprintln(w, "_____________________________________")
		fmt.Fprintln(w, "")
	}
	
	InfoLogger.Println("displayCollection | Current watch collection list was sent.")
}

// logCheck is a checker that returns the user's index and True if a user is logged in, else it returns -1 and False.
func logCheck (u *UserDatabase) (int, bool) {
	for i, user := range u.Users {
		if user.IsLoggedIn {
			return i, true
		}
	}
	InfoLogger.Println("logCheck | Login check boolean and index was sent.")
	return -1, false
}

// signup registers a user into the database with a unique username.
func (u *UserDatabase) signup(w http.ResponseWriter, r *http.Request) {
	

	var ui UserInfo
	switch r.Method {
	case "POST":
		u.databaseReader()

		w.Header().Set("Content-Type", "application/json")
		
		if err := json.NewDecoder(r.Body).Decode(&ui); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			ErrorLogger.Println("signup | Error 400: Bad Request. Error on decoding json sign up.")
			return
		}

		isDuplicate := false

		for _, user := range u.Users {
			if ui.Username == user.Username {
				isDuplicate = true
				http.Error(w, "\n\n\nBad Request: Error 409. User already exists in Database.\n", http.StatusConflict)
				ErrorLogger.Println("signup | Bad Request: Error 409. User already exists in Database.")
			} 
		}

		if !isDuplicate && ui.Username != "" && ui.Password != "" {
			u.mu.Lock()
			u.Users = append(u.Users, ui)
			u.mu.Unlock()
			http.Error(w, "\n\n\nUser successfully added to database.\n", http.StatusCreated)
			InfoLogger.Printf("signup | User: %v was successfully added to the database.\n", ui.Username)
		} else if ui.Username == "" || ui.Password == "" {
			http.Error(w, "\n\n\nBad Request: Error 405: Please enter your desired credentials properly.\n", http.StatusMethodNotAllowed)
			ErrorLogger.Println("signup | Bad Request: Error 405. Incorrect input.")
		}
		u.databaseWriter()
		
	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "\n\n\nBad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
		ErrorLogger.Println("signup | Bad Request: Error 405. Method other than POST was used.")
	}
}

// login turns the IsLoggedIn boolean of the user whose credentials are correctly provided to True.
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
		
		if len(u.Users) == 0 {
			http.Error(w, "\n\n\nBad Request: Error 404. Database not found.", http.StatusNotFound)
		} else {
			if err := json.NewEncoder(w).Encode(u.Users); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				ErrorLogger.Println("login | Bad Request: Error 500. Status Internal Server Error.")
				return
			}
		}

	case "POST":
		u.databaseReader()
		w.Header().Set("Content-Type", "application/json")
		
		if err := json.NewDecoder(r.Body).Decode(&ui); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			ErrorLogger.Println("login | Error 400: Bad Request. Error on decoding json sign up.")
			return
		}
		
		for _, user := range u.Users {
			if user.IsLoggedIn {
				amtLogged++
			}
		}

		u.mu.Lock()
		for j, user := range u.Users {
			if ui.Username == user.Username && ui.Password == user.Password && amtLogged == 0 {
				u.Users[j].IsLoggedIn = true
				isLogged = true
				index = j
				userCorrect = true
				passCorrect = true
				amtLogged++
			} else if ui.Username == user.Username && ui.Password == user.Password && user.IsLoggedIn && amtLogged == 1 {
				isLogged = true
				index = j
				userCorrect = true
				passCorrect = true
				alreadyLogged = true
			} else if ui.Username == user.Username && ui.Password != user.Password {
				userCorrect = true
			} else if ui.Username == user.Username && ui.Password == user.Password && amtLogged == 1 {
				amtLogged++
			} 
		}
		u.mu.Unlock()
		
		if alreadyLogged {
			fmt.Fprintf(w, "\n\n\nYou are already logged in, %v\n",  u.Users[index].Username)
			WarningLogger.Printf("login | %v attemped to log in again.\n", ui.Username)
		} else if isLogged && amtLogged == 1 && userCorrect && passCorrect {
			fmt.Fprintf(w, "\n\n\nLogged in successfully.\nWelcome to your Watch Collection, %v.\n", u.Users[index].Username)
			InfoLogger.Printf("login | %v succesfully logged in.\n", ui.Username)
		} else if !isLogged && amtLogged == 0 && userCorrect {
			http.Error(w, "\n\n\nError 400: Password entered is incorrect.\nPlease try again.\n", http.StatusBadRequest)
			ErrorLogger.Println("login | Error 400: Bad Request. Incorrect password.")
		} else if !isLogged && amtLogged == 0 && !userCorrect && !passCorrect {
			http.Error(w, "\n\n\nError 400: Account does not exist.\n", http.StatusBadRequest)
			ErrorLogger.Println("login | Error 400: Bad Request. Account does not exist.")
		} else {
			http.Error(w, "\n\n\nError 400: Please log out of your current account before attempting to log in again.\n", http.StatusBadRequest)
			ErrorLogger.Println("login | Error 400: Bad Request. Attempted to log in a different account without logging out.")
		}
		u.databaseWriter()
		
	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "\n\n\nBad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
		ErrorLogger.Println("login | Bad Request: Error 405. Method other than GET and POST was used.")
	}
}

// logout turns the IsLoggedIn boolean of the currently logged in user to false.
func (u *UserDatabase) logout(w http.ResponseWriter, r *http.Request) {
	var ui UserInfo
	
	switch r.Method {
	case "GET": 
		w.Header().Set("Content-Type", "application/json")
		
		if len(u.Users) == 0 {
			http.Error(w, "\n\n\nBad Request: Error 404. Database not found.", http.StatusNotFound)
			WarningLogger.Println("delete | Database not found.")
		} else {
			if err := json.NewEncoder(w).Encode(u.Users); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				ErrorLogger.Println("logout | Bad Request: Error 500. Status Internal Server Error.")
				return
			}
		}	

	case "POST":
		u.databaseReader()
		w.Header().Set("Content-Type", "application/json")
		
		allOut := false
		if err := json.NewDecoder(r.Body).Decode(&ui); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			ErrorLogger.Println("logout | Error 400: Bad Request. Error on decoding json sign up.")
			return
		}

		u.mu.Lock()
		for j := range u.Users {
			if u.Users[j].IsLoggedIn {
				u.Users[j].IsLoggedIn = false
				fmt.Fprintf(w, "\n\n\nSuccessfully logged out.\nThank you for using our application, %v.\n", u.Users[j].Username)
				allOut = true
				InfoLogger.Printf("logout | %v succesfully logged out\n.", ui.Username)
				break
			} 
		}
		u.mu.Unlock()

		if !allOut {
			fmt.Fprintln(w, "\n\n\nYou are not logged in.")
			WarningLogger.Println("logout | Attempted to log out while not logged in.")
		}
		u.databaseWriter()
		
	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "\n\n\nBad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
		ErrorLogger.Println("logout | Bad Request: Error 405. Method other than GET and POST was used.")
	}
}

// delete removes a registered user from the database after providing that user's credentials.
func (u *UserDatabase) delete(w http.ResponseWriter, r *http.Request) {
	var ui UserInfo
	userCorrect := false
	passCorrect := false
	switch r.Method {
	case "GET": 
		w.Header().Set("Content-Type", "application/json")
		
		if len(u.Users) == 0 {
			http.Error(w, "n\n\nBad Request: Error 404. Database not found.", http.StatusNotFound)
			WarningLogger.Println("delete | Database not found.")
		} else {
			if err := json.NewEncoder(w).Encode(u.Users); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				ErrorLogger.Println("delete | Bad Request: Error 500. Status Internal Server Error.")
				return
			}
		}	

	case "POST":
		u.databaseReader()
		w.Header().Set("Content-Type", "application/json")
		
		if err := json.NewDecoder(r.Body).Decode(&ui); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			ErrorLogger.Println("delete | Error 400: Bad Request. Error on decoding json sign up.")
			return
		}

		if len(u.Users) == 0 {
			http.Error(w, "\n\n\nThere are currently no Users in the database to delete", http.StatusBadRequest)
			ErrorLogger.Println("delete | Error 400: Bad Request. No users in database to delete.")
		} else {
			u.mu.Lock()
			for i, user := range u.Users {
				if ui.Username == user.Username && ui.Password == user.Password  {
					u.Users = append(u.Users[:i], u.Users[i+1:]...)
					fmt.Fprintf(w, "\n\n\nThank you for trying our application, %v.\nYour account has been successfully deleted.\n", user.Username)
					InfoLogger.Printf("delete | User %v was deleted from the database.\n", user.Username)
					userCorrect = true
					passCorrect = true
					break
				} else if ui.Username == user.Username && ui.Password != user.Password {
					userCorrect = true
				}
			}
			u.mu.Unlock()
				
			if userCorrect && !passCorrect {
					http.Error(w, "\n\n\nError 400: Password entered is incorrect. Please try again.", http.StatusBadRequest)
					ErrorLogger.Println("delete | Error 400: Bad Request. Password entered is incorrect.")
			} else if !userCorrect && !passCorrect {
					http.Error(w, "\n\n\nError 400: Specified user wasn't found", http.StatusBadRequest)
					ErrorLogger.Println("delete | Error 400: Bad Request. User to delete was not found.")
			}


		}
		u.databaseWriter()
	
		
	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "\n\n\nBad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
		ErrorLogger.Println("delete | Bad Request: Error 405. Method other than GET and POST was used.")
	}
}

// watchList displays all the watches in the user's database.
func (u *UserDatabase) watchList (w http.ResponseWriter, r *http.Request) {
	var wi WatchInfo
	switch r.Method {
	case "GET":
		u.databaseReader()
		w.Header().Set("Content-Type", "application/json")

		if len(u.Users) == 0 {
			http.Error(w, "\n\n\nUser database is currently empty.", http.StatusNotFound)
			WarningLogger.Println("watchList | Database not found.")
		} else {
			index, isLogged := logCheck(u)

			if index == -1 {
				fmt.Fprintln(w, "\n\n\nPlease log in before attempting to add a watch to your collection.")
				WarningLogger.Println("watchList | Login before adding a watch to collection.")
			}
			
			if isLogged {
				if len(u.Users[index].Watches) == 0 {
					fmt.Fprintln(w, "-------------------------------------\nYour collection is currently empty.\n-------------------------------------")
					WarningLogger.Println("watchList | Watch collection is currently empty.")
				} else {
					
					displayCollection(u, wi, u.Users[index].Username, len(u.Users[index].Watches), index, w)

					InfoLogger.Printf("watchList | Current watch collection of User: %v was displayed.\n", u.Users[index].Username)
				}
			} 
		}
		
	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "\n\n\nBad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
		ErrorLogger.Println("watchList | Bad Request: Error 405. Method other than GET was used.")
	}
}

// addWatch appends a watch to the logged in user's watch collection based on the details specified.
func (u *UserDatabase) addWatch(w http.ResponseWriter, r *http.Request) {
	var wi WatchInfo
	switch r.Method {
	case "POST":
		u.databaseReader()
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewDecoder(r.Body).Decode(&wi); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			ErrorLogger.Println("addWatch | Error 400: Bad Request. Error on decoding json sign up.")
			return
		}

		index, isLogged := logCheck(u)

		isDuplicate := false

		if index == -1 {
			fmt.Fprintln(w, "\n\n\nPlease log in before attempting to add a watch to your collection.")
			WarningLogger.Println("addWatch | Attempted to add a watch before logging in.")
		}

		if isLogged {
			for _, watch := range u.Users[index].Watches {
				if wi.Brand == watch.Brand && wi.Model == watch.Model {
					isDuplicate = true

				}
			}
			
			u.mu.Lock()
			if wi.Brand != "" && wi.Model != "" && wi.Width != "" && wi.Price != "" && !isDuplicate {
				u.Users[index].Watches = append(u.Users[index].Watches, wi)
				fmt.Fprintf(w, "\n\n\nSuccessfully added %v %v.\n", wi.Brand, wi.Model)	
				InfoLogger.Printf("addWatch | Added %v %v to the collection of User: %v.\n", wi.Brand, wi.Model, u.Users[index].Username)

				displayCollection(u, wi, u.Users[index].Username, len(u.Users[index].Watches), index, w)

				InfoLogger.Printf("addWatch | Current watch collection of User: %v was displayed.\n", u.Users[index].Username)
			} else if wi.Brand != "" && wi.Model != "" && wi.Width != "" && wi.Price != "" && isDuplicate {
				http.Error(w, "\n\n\nThis watch is already a part of your collection.", http.StatusMethodNotAllowed)
				ErrorLogger.Println("addWatch | Bad Request: Error 405. User attempted to add a watch that's already part of their collection.")

				displayCollection(u, wi, u.Users[index].Username, len(u.Users[index].Watches), index, w)

				InfoLogger.Printf("addWatch | Current watch collection of User: %v was displayed.\n", u.Users[index].Username)
			} else {
				http.Error(w, "\n\n\nPlease include the desired brand, model, width, and price of the item.", http.StatusMethodNotAllowed)
				ErrorLogger.Println("addWatch | Bad Request: Error 405. User attempted to add a watch with incomplete details.")

				displayCollection(u, wi, u.Users[index].Username, len(u.Users[index].Watches), index, w)

				InfoLogger.Printf("addWatch | Current watch collection of User: %v was displayed.\n", u.Users[index].Username)
			}
			u.mu.Unlock()
		} 
		u.databaseWriter()

	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "\n\n\nBad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
		ErrorLogger.Println("addWatch | Bad Request: Error 405. Method other than POST was used.")
	}
}

// removeWatch deletes the watch from the logged in user's watch collection based on the provided Brand and Model.
func (u *UserDatabase) removeWatch(w http.ResponseWriter, r *http.Request) {
	var wi WatchInfo
	removed := false

	switch r.Method {
	case "POST":
		u.databaseReader()
		w.Header().Set("Content-Type", "application/json")
		
		if err := json.NewDecoder(r.Body).Decode(&wi); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			ErrorLogger.Println("removeWatch | Error 400: Bad Request. Error on decoding json sign up.")
			return
		}
		
		index, isLogged := logCheck(u)
		
		if index == -1 {
			fmt.Fprintln(w, "\n\n\nPlease log in before atemmpting to remove a watch from your collection.")
			WarningLogger.Println("removeWatch | Attempted to remove a watch before logging in.")
		}
		
		if isLogged {
			for j, watch := range u.Users[index].Watches {
				u.mu.Lock()
				if wi.Brand == u.Users[index].Watches[j].Brand && wi.Model == u.Users[index].Watches[j].Model {
					u.Users[index].Watches = append(u.Users[index].Watches[:j], u.Users[index].Watches[j+1:]...)
					fmt.Fprintf(w, "\n\n\nSuccessfully removed %v %v.\n", watch.Brand, watch.Model)		
					InfoLogger.Printf("removeWatch | Removed %v %v from the collection of User: %v.\n", watch.Brand, watch.Model, u.Users[index].Username)
					removed = true
					break
				}
				u.mu.Unlock()
			}
			
			if !removed {
				fmt.Fprintln(w, "\n\n\nThe watch with the provided details was not found in your collection.")
				WarningLogger.Printf("removeWatch | %v attempted to delete a watch that is not part of the collection.\n", u.Users[index].Username)
			}
			if len(u.Users[index].Watches) > 0 {
				displayCollection(u, wi, u.Users[index].Username, len(u.Users[index].Watches), index, w)
				InfoLogger.Printf("removeWatch | Current watch collection of User: %v was displayed.\n", u.Users[index].Username)
			} else {
				fmt.Fprintln(w, "-------------------------------------\nYour collection is currently empty.\n-------------------------------------")
				WarningLogger.Println("removeWatch | Watch collection is currently empty.")
			}
		}
		u.databaseWriter()

	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "\n\n\nBad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
		ErrorLogger.Println("removeWatch | Bad Request: Error 405. Method other than POST was used.")
	}
}

// markWatch toggles the Collected bool of a watch based on the provided Brand and Model.
func (u *UserDatabase) markWatch(w http.ResponseWriter, r *http.Request) {
	var wi WatchInfo
	switch r.Method {
	case "POST":
		u.databaseReader()
		w.Header().Set("Content-Type", "application/json")
		
		if err := json.NewDecoder(r.Body).Decode(&wi); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			ErrorLogger.Println("markWatch | Error 400: Bad Request. Error on decoding json sign up.")
			return
		}
		
		index, isLogged := logCheck(u)
		watchIndex := -1

		exists := false

		if index == -1 {
			fmt.Fprintln(w, "\n\n\nPlease log in before trying to mark a watch from your collection.")
			WarningLogger.Println("markWatch | Attempted to mark a watch before logging in.")
		}
		
		if isLogged {
			if wi.Brand == "" || wi.Model == "" {
				http.Error(w, "\n\n\nPlease enter the watch brand and model properly.\n", http.StatusMethodNotAllowed)
				WarningLogger.Println("markWatch | Bad Request: Error 405. Incorrect input.")
			} else {
				for j, watch := range u.Users[index].Watches {
					if wi.Brand == watch.Brand && wi.Model == watch.Model {
						exists = true

						u.mu.Lock()
						if !watch.Collected {
							watchIndex = j
							u.Users[index].Watches[j].Collected = true
							InfoLogger.Printf("markWatch | %v %v was marked as COLLECTED from the collection of User: %v.\n", wi.Brand, wi.Model, u.Users[index].Username)
							break
						} else if watch.Collected {
							watchIndex = j
							u.Users[index].Watches[j].Collected = false
							InfoLogger.Printf("markWatch | %v %v was marked as NOT COLLECTED from the collection of User: %v.\n", wi.Brand, wi.Model, u.Users[index].Username)
							break
						}	
						u.mu.Unlock()
					} 
				}
				
				if !exists {
					fmt.Fprintf(w, "\n\n\n%v is not in your collection.\n", wi.Model)
					WarningLogger.Printf("markWatch | %v %v is not in the collection of User: %v.\n", wi.Brand, wi.Model, u.Users[index].Username)
				} else {
					if u.Users[index].Watches[watchIndex].Collected {
						fmt.Fprintf(w, "\n\n\n%v %v successfully marked as COLLECTED.\n", wi.Brand, wi.Model)
					} else {
						fmt.Fprintf(w, "\n\n\n%v %v successfully marked as NOT COLLECTED.\n", wi.Brand, wi.Model)
					}
				}
	
				if len(u.Users[index].Watches) > 0 {
					displayCollection(u, wi, u.Users[index].Username, len(u.Users[index].Watches), index, w)
					InfoLogger.Printf("markWatch | Current watch collection of User: %v was displayed.\n", u.Users[index].Username)
				} else {
					fmt.Fprintln(w, "-------------------------------------\nYour collection is currently empty.\n-------------------------------------")
					WarningLogger.Println("markWatch | Watch collection is currently empty.")
				}
			}

			
		}
		u.databaseWriter()

	default:
		w.Header().Set("Content-Type", "application/json")

		http.Error(w, "\n\n\nBad Request: Error 405. Action is not allowed.", http.StatusMethodNotAllowed)
		ErrorLogger.Println("markWatch | Bad Request: Error 405. Method other than POST was used.")
	}
}