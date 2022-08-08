package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RemiPail/moncha/pkg/models"
)

func (h handler) GetAllShopItems(w http.ResponseWriter, r *http.Request) {
	var shopItems []models.ShopItem

	if result := h.DB.Find(&shopItems); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shopItems)
}
