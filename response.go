package pangea

import "net/http"

type Response struct {
	OriginRes http.ResponseWriter
}
