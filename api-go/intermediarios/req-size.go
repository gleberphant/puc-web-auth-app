package intermediarios

import "net/http"

func ReqSizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// prepara a resposta
		res.Header().Set("Content-Type", "application/json")

		// limita tamanho do body Request para a 1MB para não travar servidor
		req.Body = http.MaxBytesReader(res, req.Body, 1048576)

		// agenda fechamento do body Request para liiberar recursos
		defer req.Body.Close()

		next.ServeHTTP(res, req)
	})
}
