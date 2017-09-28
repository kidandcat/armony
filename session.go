package armony

import (
	"net/http"
	"strings"
	"time"
)

// Session : Object representing a session
type Session struct {
	w *http.ResponseWriter
	r *http.Request
}

func (ss *Session) setCookie(key string, value string) {
	expire := time.Now().AddDate(0, 0, 30)
	ck := http.Cookie{Name: key, Value: value, Expires: expire}
	http.SetCookie(*ss.w, &ck)
}

func (ss *Session) getCookie(key string) string {
	cookie, err := ss.r.Cookie(key)
	if err != nil {
		return ""
	}
	return cookie.Value
}

//Set : Set a key:value pair in the session
func (ss *Session) Set(key string, value string) {
	if ss.getCookie("_ID") == "" {
		return
	}
	kvSet(strings.Trim(ss.getCookie("_ID")+key, " "), value)
}

//Get : Get a value from the session
func (ss *Session) Get(key string) string {
	k := strings.Trim(ss.getCookie("_ID")+key, " ")
	value := kvGet(k)
	return value
}

func (ss *Session) setSessionID() {
	ss.setCookie("_ID", RandString(10))
}

// LoadSession : Look for a existing session || load a new one
func LoadSession(w *http.ResponseWriter, r *http.Request) Session {
	ss := Session{
		w: w,
		r: r,
	}
	if ss.getCookie("_ID") == "" {
		ss.setSessionID()
	}
	return ss
}
