package response

type registerResponse struct {
	Fullname   string `json:"fullname"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Token      string `json:"token"`
}
