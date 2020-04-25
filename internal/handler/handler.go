package handler

import (
	"fmt"
	"net/http"
	"workshop/internal/api"
)

type Handler struct {
	jokeClient api.Client
}

func NewHandler(jokeClient api.Client) *Handler {
	return &Handler{
		jokeClient: jokeClient,
	}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	fmt.Fprint(w, "hello golang Saint-Petersburg!")
=======
	joke, err := h.jokeClient.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, joke.Joke)
>>>>>>> d95f61c0e8559840c8a19b0f32967204dddc5b09
}
