/* ----------------------------------
*  @author suyame 2022-10-17 16:28:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package file

import "net/http"

func Dirhandler(patern, filedir string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		handler := http.StripPrefix(patern, http.FileServer(http.Dir(filedir)))
		handler.ServeHTTP(w, req)
	}
}
