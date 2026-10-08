package reviews

import (
	"log"
	"net/http"
)

func (h *Handler) GetReviews(w http.ResponseWriter, r *http.Request) {
	log.Println("Get reviews Handler called")
}
