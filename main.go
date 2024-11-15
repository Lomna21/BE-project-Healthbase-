package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/gorilla/sessions" //-------------------
)

// ----------------------------------------------------------------
var store = sessions.NewCookieStore([]byte("secret-key"))

// When the "Logout" button is clicked in the dashboard, the logout() JavaScript function is triggered.
// This sends a POST request to the /logout endpoint.
// The Go backend processes this by clearing the session data and redirecting the user back to the login page.

func logout(w http.ResponseWriter, r *http.Request) {
	// Destroy the session
	session, _ := store.Get(r, "user-session")   // retrieves the current session using store.Get().
	session.Values["username"] = nil     // Clears the session data by setting session.Values["username"] to nil
	session.Save(r, w)

	// Redirect the user to the login page after logout
	http.Redirect(w, r, "/login", http.StatusSeeOther)  // Redirects the user to the login page using http.Redirect.
}

//------------------------------------------------------------------

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login func is called")
	fileName := "E:/.vscode/BE Project/htmlpages/login.html"
	fileName2 := "login.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Printf("Error message while parsing %s: %v\n", fileName, err)
		return
	}
	err = t.ExecuteTemplate(w, fileName2, nil)
	if err != nil {
		fmt.Printf("Error message while executing %s: %v\n", fileName2, err)
		return
	}
}

// Updated userDB with roles for each user
var userDB = map[string]struct {
	Password string
	Role     string
}{
	"harshit": {Password: "21591", Role: "doctor"},
	"riya":    {Password: "21547", Role: "patient"},
	"anmol": {Password: "21575", Role: "patient"},
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginSubmit func is called")

	username := r.FormValue("username")
	password := r.FormValue("password")
	role := r.FormValue("role")
	fmt.Println(username, password, role)

	// Validate user credentials and role
	if user, exists := userDB[username]; exists && user.Password == password && user.Role == role {
		w.WriteHeader(http.StatusOK)

		// Choose the appropriate dashboard based on role
		var fileName, fileName2 string
		if role == "doctor" {
			fileName = "E:/.vscode/BE Project/htmlpages/doctordashboard.html"
			fileName2 = "doctordashboard.html"
		} else if role == "patient" {
			fileName = "E:/.vscode/BE Project/htmlpages/userdashboard.html"
			fileName2 = "userdashboard.html"
		}

		// Parse and execute the appropriate dashboard template
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Printf("Error message while parsing %s: %v\n", fileName, err)
			return
		}

		type User struct {
			Name     string
			Password string
			Role     string
		}

		err = t.ExecuteTemplate(w, fileName2, User{Name: username, Password: password, Role: role})
		if err != nil {
			fmt.Printf("Error message while executing %s: %v\n", fileName2, err)
			return
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid username, password, or role.")
	}
}

func appointments(w http.ResponseWriter, r *http.Request) {

}
func profile(w http.ResponseWriter, r *http.Request) {

}
func activities(w http.ResponseWriter, r *http.Request) {

}
func settings(w http.ResponseWriter, r *http.Request) {

}
func medicalRecords(w http.ResponseWriter, r *http.Request) {

}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	req := r.URL.Path

// 	switch req {
// 	case "/login":
// 		login(w, r)
// 	case "/login-submit":
// 		loginSubmit(w, r)
// 	default:
// 		fmt.Fprintf(w, "HOME Page")
// 	}
// }

func handler(w http.ResponseWriter, r *http.Request) {
	req := r.URL.Path

	switch req {
	case "/login":
		login(w, r)
	case "/login-submit":
		loginSubmit(w, r)
	case "/logout": //---------------
		logout(w, r)
	case "/appointments":
		appointments(w, r)
	case "/profile":
		profile(w, r)
	case "/activities":
		activities(w, r)
	case "/settings":
		settings(w, r)
	case "/medicalRecords":
		medicalRecords(w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("server is starting\n")
	http.ListenAndServe(":8000", nil)
}
