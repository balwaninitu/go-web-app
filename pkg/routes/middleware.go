package routes

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

//middleware
func WriteToConsole(next http.Handler) http.Handler {
	//return has anonymous func which allows to use middleware
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the Page")
		next.ServeHTTP(w, r)
	})
}

//nosrf new will create handler
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	/*set value, it uses cookies to make sure token is generated is
	  available on per page basis*/
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
