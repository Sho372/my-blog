package handlers

import (
	"net/http"
)

func CheckAuth(w http.ResponseWriter, r *http.Request) {
    session, err := store.Get(r, "session-name")
    if err != nil {
        http.Error(w, "Failed to retrieve session", http.StatusInternalServerError)
        return
    }

    if auth, ok := session.Values["authenticated"].(bool); ok && auth {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"authenticated": true}`))
    } else {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"authenticated": false}`))
    }
}
