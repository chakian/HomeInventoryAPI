package controllers

import (
	"log"
	"net/http"

	"HomeInventoryAPI/models"
	u "HomeInventoryAPI/utils"
)

// GetAllInventory ...
var GetAllInventory = func(w http.ResponseWriter, r *http.Request) {
	log.Print("Started: inventoryController.GetAllInventory")

	var invItems [3]models.InventoryItem
	invItems[0].ID = 1
	invItems[0].Name = "Bira"

	invItems[1].ID = 3
	invItems[1].Name = "Ton Balığı"

	invItems[2].ID = 4
	invItems[2].Name = "Makarna"

	data := invItems
	resp := u.Message(true, "ok")
	resp["data"] = data
	u.Respond(w, resp)

	log.Print("Finished: inventoryController.GetAllInventory")
}
