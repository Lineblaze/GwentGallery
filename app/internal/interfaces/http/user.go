package http

import (
	"github.com/Lineblaze/GwentGallery/app/internal/domain"
	"net/http"
)

func (s *server) GetUser(w http.ResponseWriter, r *http.Request) {
	userId := s.parseParamInt64("user_id", r)

	user, err := s.core.GetUser(r.Context(), &domain.GetUserRequest{UserId: userId})
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, user, w)
}
