package http

import (
	"fmt"
	"github.com/Lineblaze/GwentGallery/app/internal/domain"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (s *server) CreateCard(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateCardRequest
	//if err := s.readJSON(&req, r); err != nil {
	//	s.sendError(err, w)
	//	return
	//}
	//Name
	req.Name = strings.TrimSpace(r.FormValue("name"))
	//CategoryId
	categoryIdString := r.FormValue("category_id")
	if categoryIdString != "" {
		categoryId, err := strconv.Atoi(categoryIdString)
		if err != nil {
			s.sendError(fmt.Errorf("parsing category_id: %w", err), w)
			return
		}
		temp := uint8(categoryId)
		req.CategoryId = &temp
	}
	//Description
	req.Description = strings.TrimSpace(r.FormValue("description"))
	//Provision
	provisionString := r.FormValue("provision")
	provision, err := strconv.Atoi(provisionString)
	if err != nil {
		s.sendError(fmt.Errorf("parsing provision: %w", err), w)
		return
	}
	req.Provision = uint8(provision)
	//Power
	powerString := r.FormValue("power")
	if powerString != "" {
		power, err := strconv.Atoi(powerString)
		if err != nil {
			s.sendError(fmt.Errorf("parsing power: %w", err), w)
			return
		}
		temp := uint8(power)
		req.Power = &temp
	}
	//Faction
	factionString := r.FormValue("faction")
	faction, err := strconv.Atoi(factionString)
	if err != nil {
		s.sendError(fmt.Errorf("parsing faction: %w", err), w)
		return
	}
	req.Faction = domain.Faction(int8(faction))
	//Rarity
	rarityString := r.FormValue("rarity")
	rarity, err := strconv.Atoi(rarityString)
	if err != nil {
		s.sendError(fmt.Errorf("parsing rarity: %w", err), w)
		return
	}
	req.Rarity = domain.Rarity(int8(rarity))
	//CardType
	cardTypeString := r.FormValue("card_type")
	cardType, err := strconv.Atoi(cardTypeString)
	if err != nil {
		s.sendError(fmt.Errorf("parsing type: %w", err), w)
		return
	}
	req.CardType = domain.CardType(int8(cardType))
	//Color
	colorString := r.FormValue("color")
	color, err := strconv.Atoi(colorString)
	if err != nil {
		s.sendError(fmt.Errorf("parsing color: %w", err), w)
		return
	}
	req.Color = domain.Color(int8(color))
	//Set
	setString := r.FormValue("set")
	set, err := strconv.Atoi(setString)
	if err != nil {
		s.sendError(fmt.Errorf("parsing set: %w", err), w)
		return
	}
	req.Color = domain.Color(int8(set))
	//ImageAuthor
	req.ImageAuthor = strings.TrimSpace(r.FormValue("image_author"))
	//AuthorId
	authorIdString := r.FormValue("author_id")
	authorId, err := strconv.Atoi(authorIdString)
	if err != nil {
		s.sendError(fmt.Errorf("parsing author_id: %w", err), w)
		return
	}
	req.AuthorId = int64(authorId)
	//Status
	statusString := r.FormValue("status")
	status, err := strconv.Atoi(statusString)
	if err != nil {
		s.sendError(fmt.Errorf("parsing status: %w", err), w)
		return
	}
	req.Status = domain.Status(int8(status))
	// Получаем файл из формы запроса
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, fmt.Sprintf("parsing inage file: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Создаем новый файл на сервере для сохранения изображения
	imagePath := "./images/" + header.Filename
	out, err := os.Create(imagePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error saving image: %s", err), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Копируем данные из загруженного файла в новый файл на сервере
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error coping image: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Image uploaded successfully.")
	return

	card, err := s.core.CreateCard(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusCreated, card, w)
}

func (s *server) GetCard(w http.ResponseWriter, r *http.Request) {
	req := domain.GetCardRequest{CardId: s.parseParamInt64("card_id", r)}
	card, err := s.core.GetCard(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, card, w)
}

func (s *server) UpdateCard(w http.ResponseWriter, r *http.Request) {
	var req domain.UpdateCardRequest
	if err := s.readJSON(&req, r); err != nil {
		s.sendError(err, w)
		return
	}

	req.CardId = s.parseParamInt64("card_id", r)

	card, err := s.core.UpdateCard(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, card, w)
}

func (s *server) DeleteCard(w http.ResponseWriter, r *http.Request) {
	req := domain.DeleteCardRequest{CardId: s.parseParamInt64("card_id", r)}

	_, err := s.core.DeleteCard(r.Context(), &req)
	if err != nil {
		s.sendError(err, w)
		return
	}

	s.sendJSON(http.StatusOK, nil, w)
}
