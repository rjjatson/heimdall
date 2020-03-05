package hub

import (
	"encoding/json"
	"heimdall/internal/router"
	"heimdall/pkg/model"
)

const (
	authMethoddebug = "debug"
)

// AuthorizationRequest model for auth request
type AuthorizationRequest struct {
	model.Request
	Type   string `json:"type"`
	Method string `json:"method"`
	Token  string `json:"token"`
}

// AuthorizationResponse model for auth response
type AuthorizationResponse struct {
	model.Response
	UserID string `json:"user_id"`
}

// AuthorizeUser special handlers called by router
func (h *Hub) AuthorizeUser(msg []byte, rw *router.ResponseWriter) {
	var authReq AuthorizationRequest
	json.Unmarshal(msg, &authReq)

	switch authReq.Method {
	case authMethoddebug:
		h.authenticateDebug(authReq.SenderID, authReq.Token)
		rw.WriteResponse(&AuthorizationResponse{
			UserID:   authReq.Token,
			Response: model.Response{ReceiverID: authReq.Token},
		})
	}

}

func (h *Hub) authenticateDebug(oldID, newID string) error {
	c, err := h.GetClient(oldID)
	if err != nil {
		return err
	}
	h.RemoveClient(oldID)
	c.SetAuthentication(true)
	h.AddClient(newID, c)
	return nil
}
