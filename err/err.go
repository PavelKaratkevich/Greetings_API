package err

type Error struct {
	Message string `json:"message"`
	Status int `json:"status"`
}

//func SendError (w http.ResponseWriter, e Error) {
//		w.Header().Add("Content-Type", "application/json")
//		w.WriteHeader(e.Status)
//		json.NewEncoder(w).Encode(e)
//	}
//


