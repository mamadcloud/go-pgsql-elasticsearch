package ad_param

import (
	"fmt"
	"strings"

	"API/pkg/entities"
	pgsql "API/pkg/configs/db/postgres"
)

func Create(adsID int, adsParams []entities.AdParam) ([]entities.AdParam, error) {
	db, _ := pgsql.PGConnect()
	noOfColumns := 3
	valueStrings := make([]string, 0, len(adsParams))
	valueArgs := make([]interface{}, 0, len(adsParams) * noOfColumns)
	for i, param := range adsParams {
		adsParams[i].AdID = adsID
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i*noOfColumns+1, i*noOfColumns+2, i*noOfColumns+3))
		valueArgs = append(valueArgs, adsID)
		valueArgs = append(valueArgs, param.Name)
		valueArgs = append(valueArgs, param.Value)
	}
	sqlStatement := fmt.Sprintf("INSERT INTO ad_params VALUES %s", strings.Join(valueStrings, ","))
	_, err := db.Exec(sqlStatement, valueArgs...)
	if err != nil {
		return nil, err
	}

	defer pgsql.PGClose(db)
	return adsParams, nil
}

func Edit(adsID int, adsParams []entities.AdParam) ([]entities.AdParam, error) {
	err := Delete(adsID)
	if err != nil {
		return nil, err
	}

	adsParams, err = Create(adsID, adsParams)
	if err != nil {
		return nil, err
	}

	return adsParams, nil
}

func Delete(adsID int) (error) {
	db, _ := pgsql.PGConnect()
	
	sqlStatement := `DELETE FROM ad_params WHERE ad_id = $1`
	_, err := db.Exec(sqlStatement, adsID)

	if err != nil {
		return err
	}

	defer pgsql.PGClose(db)
	return nil
}