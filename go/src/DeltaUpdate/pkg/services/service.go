package service

import (
	"fmt"

	"DeltaUpdate/pkg/entities"
	pgsql "DeltaUpdate/pkg/configs/db/postgres"
)

var LIMIT = 1000

func FindDeltaUpdates() ([]entities.DeltaUpdate, error) {
	deltaUpdates := []entities.DeltaUpdate{}
	db, _ := pgsql.PGConnect()

	sqlStatement := `SELECT * FROM delta_update LIMIT $1`
	rows, err := db.Query(sqlStatement, LIMIT)

	if err != nil {
		return nil, nil
	}

	for rows.Next(){
		deltaUpdate := entities.DeltaUpdate{}
		err := rows.Scan(&deltaUpdate.ID, &deltaUpdate.ItemID, &deltaUpdate.Action, &deltaUpdate.TableName, &deltaUpdate.Timestamp)
    if err != nil {
			fmt.Println(err)
		} 
    deltaUpdates = append(deltaUpdates, deltaUpdate)
	}

	defer pgsql.PGClose(db)
	return deltaUpdates, nil
}

func FindAd(id int) (*entities.Ads, error) {
	ads := entities.Ads{}
	adsTemp := entities.Ads{}
	ads.AdParams = []entities.AdParam{}
	db, _ := pgsql.PGConnect()

	sqlStatement := `SELECT ads.ad_id, list_id, orig_list_time, list_time, status, type, ads.name, phone, region, city, category, user_id, phone_hidden, no_salesmen, company_ad, subject, body, price, old_price, image, infopage, infopage_title, store_id, modified_at, adtype, ad_params.name, value FROM ads INNER JOIN ad_params ON ad_params.ad_id = ads.ad_id WHERE ads.ad_id = $1`
	rows, err := db.Query(sqlStatement, id)

	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	i := 0;

	name := ""
	value := ""

	for rows.Next(){
		err := rows.Scan(&adsTemp.ID, &adsTemp.ListID, &adsTemp.OriginalListTime, &adsTemp.ListTime, &adsTemp.Status, &adsTemp.Type, &adsTemp.Name, &adsTemp.Phone, &adsTemp.Region,
								&adsTemp.City, &adsTemp.Category, &adsTemp.UserID, &adsTemp.PhoneHidden, &adsTemp.NoSalesmen, &adsTemp.CompanyAd, &adsTemp.Subject, &adsTemp.Body, &adsTemp.Price,
								&adsTemp.OldPrice, &adsTemp.Image, &adsTemp.InfoPage, &adsTemp.InfoPageTitle, &adsTemp.StoreID, &adsTemp.ModifiedAt, &adsTemp.AdType, &name, &value)
    if err != nil {
			fmt.Println(err)
		}

		if i == 0 {
			ads = adsTemp
		}
		ads.AdParams = append(ads.AdParams, entities.AdParam{ AdID: id, Name: name, Value: value })
		i++
	}

	defer pgsql.PGClose(db)
	return &ads, nil
}

func DeleteDeltaUpdate(id int) (error) {
	db, _ := pgsql.PGConnect()
	
	sqlStatement := `DELETE FROM delta_update WHERE id = $1`
	_, err := db.Exec(sqlStatement, id)

	if(err != nil) {
		return err
	}

	defer pgsql.PGClose(db)
	return nil
}