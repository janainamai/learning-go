package dto

type (
	CreateProductInput struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	CreateUserInput struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	GetJWTInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	GetJWTOutput struct {
		AccessToken string `json:"access_token"`
	}
)
