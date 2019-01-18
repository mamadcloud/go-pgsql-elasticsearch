package ads

import (
	"time"
	"fmt"

	"API/pkg/entities"
	pgsql "API/pkg/configs/db/postgres"
)

func Create(ads entities.Ads) (entities.Ads, error) {
	db, _ := pgsql.PGConnect()
	modifiedDate := time.Now()
	id := 0

	sqlStatement := `INSERT INTO ads VALUES(DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
		$11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
		$21, $22, $23, $24, $25, $26) RETURNING ad_id`
	err := db.QueryRow(sqlStatement, ads.ListID, ads.OriginalListTime, ads.ListTime, ads.Status, ads.Type, ads.Name, ads.Phone, ads.Region,
									ads.City, ads.Category, ads.UserID, nil, ads.PhoneHidden, ads.NoSalesmen, ads.CompanyAd, ads.Subject, ads.Body, ads.Price,
									ads.OldPrice, ads.Image, ads.InfoPage, ads.InfoPageTitle, ads.StoreID, nil, modifiedDate, ads.AdType).Scan(&id)

	if(err != nil) {
		fmt.Println("error insert ==>", err)
		return ads, nil
	}

	ads.ID = id
	defer pgsql.PGClose(db)
	return ads, nil
}

func Edit(id int, ads entities.Ads) (entities.Ads, error) {
	db, _ := pgsql.PGConnect()
	modifiedDate := time.Now()

	sqlStatement := `UPDATE ads SET orig_list_time=$1, list_time=$2, status=$3, type=$4, name=$5, phone=$6, region=$7, city=$8,
		category=$9, phone_hidden=$10, no_salesmen=$11, company_ad=$12, subject=$13, body=$14, price=$15, old_price=$16, image=$17,
		infopage=$18, infopage_title=$19, modified_at=$20, adtype=$21 WHERE ad_id=$22`
	_, err := db.Exec(sqlStatement, ads.OriginalListTime, ads.ListTime, ads.Status, ads.Type, ads.Name, ads.Phone, ads.Region, ads.City,
		ads.Category, ads.PhoneHidden, ads.NoSalesmen, ads.CompanyAd, ads.Subject, ads.Body, ads.Price, ads.OldPrice, ads.Image,
		ads.InfoPage, ads.InfoPageTitle, modifiedDate, ads.AdType, id)
		
	if(err != nil) {
		fmt.Println("error update ==>", err)
		return ads, nil
	}

	ads.ModifiedAt = &modifiedDate
	ads.ID = id
	defer pgsql.PGClose(db)
	return ads, nil
}

func Delete(id int) (error) {
	db, _ := pgsql.PGConnect()
	
	sqlStatement := `DELETE FROM ads WHERE ad_id = $1`
	_, err := db.Exec(sqlStatement, id)

	if(err != nil) {
		return err
	}

	defer pgsql.PGClose(db)
	return nil
}