package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/janainamai/study-api-go/internal/dto"
	"github.com/janainamai/study-api-go/internal/entity"
	"github.com/janainamai/study-api-go/internal/infra/database"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	userDB       database.UserDatabaseInterface
	jwt          *jwtauth.JWTAuth
	jwtExpiresIn int
}

func NewUserHandler(userDB database.UserDatabaseInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		userDB:       userDB,
		jwt:          jwt,
		jwtExpiresIn: jwtExpiresIn,
	}
}

// Get user JWT token
// @Summary 	Get a user JTW token
// @Description Get a user JTW token to access the products resources
// @Tags 		users
// @Accept 		json
// @Produce 	json
// @Param 		request body dto.GetJWTInput true "user credentials"
// @Success 	200 {object} dto.GetJWTOutput
// @Failure 	400 {object} Error
// @Failure		401 {object} Error
// @Failure 	404 {object} Error
// @Failure 	500 {object} Error
// @Router 		/users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var input dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	user, err := h.userDB.FindByEmail(input.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: "user not found"}
		json.NewEncoder(w).Encode(err)
		return
	}

	if !user.ValidatePassword(input.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenSecrets := map[string]interface{}{
		"sub": user.ID,
		"exp": time.Now().Add(time.Second * time.Duration(h.jwtExpiresIn)).Unix(),
	}
	_, tokenString, _ := h.jwt.Encode(tokenSecrets)

	accessToken := dto.GetJWTOutput{
		AccessToken: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user
// @Summary 	Create user
// @Description Create a user for receive access for the system
// @Tags 		users
// @Accept 		json
// @Produce 	json
// @Param 		request body dto.CreateUserInput true "user request"
// @Success 	201
// @Failure 	500 {object} Error
// @Router 		/users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	err = h.userDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
