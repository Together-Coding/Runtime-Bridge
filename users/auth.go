package users

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type VerifiedUser struct {
	UserId      int64     `json:"userId"`
	Email       string    `json:"email"`
	IssuedAt    time.Time `json:"issuedAt"`
	ExpiredAt   time.Time `json:"expiredAt"`
	Valid       bool      `json:"valid"`
	ErrorReason *string   `json:"error_reason"`
}

// VerifyUser verifies the given token by asking the auth server.
func VerifyUser(c *gin.Context, token string) (verifiedUser VerifiedUser) {
	// Send the token to the auth server
	payload, _ := json.Marshal(map[string]string{
		"token": token,
	})
	resp, err := http.Post(os.Getenv("API_URL")+"/auth/token", "application/json", bytes.NewBuffer(payload))
	defer resp.Body.Close()
	if err != nil {
		reason := "The authorization server is not available."
		verifiedUser.ErrorReason = &reason
		verifiedUser.Valid = false
		return
	}

	// When verification is failed, process must be stopped.
	if resp.StatusCode != 200 {
		reason := "Invalid token"
		verifiedUser.ErrorReason = &reason
		verifiedUser.Valid = false
	} else {
		err = json.NewDecoder(resp.Body).Decode(&verifiedUser)
		if err != nil {
			// TODO: sentry
			reason := "Malformed data from authorization server."
			verifiedUser.ErrorReason = &reason
			verifiedUser.Valid = false
		} else {
			verifiedUser.Valid = true
		}
	}

	return
}
