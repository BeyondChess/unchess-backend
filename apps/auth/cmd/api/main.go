package main

import (
	"apps/auth/internal/server"
	"fmt"
)

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	//key := "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" // your client secret
	//maxAge := 86400 * 30                         // 30 days
	//isProd := true                               // Set to true when serving over https
	//
	//store := sessions.NewCookieStore([]byte(key))
	//store.MaxAge(maxAge)
	//store.Options.Path = "/"
	//store.Options.HttpOnly = true // HttpOnly should always be enabled
	//store.Options.Secure = isProd
	//
	//gothic.Store = store
	//goth.UseProviders(
	//	google.New("623077223352-dbv82pd0qo18i145dlo0up8n4lihpv1j.apps.googleusercontent.com", "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", "http://localhost:3000/auth/google/callback", "email", "profile"),
	//)

	//p := pat.New()
	//
	//p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {
	//	user, err := gothic.CompleteUserAuth(res, req)
	//	if err != nil {
	//		res.WriteHeader(http.StatusInternalServerError)
	//		json.NewEncoder(res).Encode(map[string]string{"error": err.Error()})
	//		return
	//	}
	//
	//	// Send user data as JSON response
	//	res.WriteHeader(http.StatusOK)
	//	json.NewEncoder(res).Encode(user)
	//})
	//
	//p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
	//	gothic.BeginAuthHandler(res, req)
	//})
	//
	//log.Println("listening on localhost:3000")
	//log.Fatal(http.ListenAndServe(":3000", p))
}
