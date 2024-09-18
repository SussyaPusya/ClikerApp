package handlers

import (
	"clickercopro/app/logic"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var sessionKey = "e8550643e8f2"

var store = sessions.NewCookieStore([]byte(os.Getenv(sessionKey)))

func CreateSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "coprosession")
	as := logic.GetPointTab()
	session.Values["Count"] = as
	fmt.Println(session.Values["Count"])
	err := session.Save(r, w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
