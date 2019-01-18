package ads

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"API/pkg/entities"
	adsService "API/pkg/services/ads"
	adParamsService "API/pkg/services/ad_param"
)

func Create(w http.ResponseWriter, r *http.Request) {
	ads := entities.Ads{}
	_ = json.NewDecoder(r.Body).Decode(&ads)

	ads, _ = adsService.Create(ads)
	ads.AdParams, _ = adParamsService.Create(ads.ID, ads.AdParams)

	json.NewEncoder(w).Encode(ads)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	adsID, _ := strconv.Atoi(id)

	ads := entities.Ads{}
	_ = json.NewDecoder(r.Body).Decode(&ads)

	ads, _ = adsService.Edit(adsID, ads)
	ads.AdParams, _ = adParamsService.Edit(ads.ID, ads.AdParams)

	json.NewEncoder(w).Encode(ads)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	adsID, _ := strconv.Atoi(id)

	_ = adsService.Delete(adsID)
	_ = adParamsService.Delete(adsID)

	f := map[string]interface{}{
		"status": "OK",
	}
	json.NewEncoder(w).Encode(f)
}