CREATE OR REPLACE FUNCTION log_update_ads_table()
  RETURNS trigger AS
$body$ 
BEGIN
 INSERT INTO delta_update VALUES(DEFAULT, NEW.ad_id, 'update', 'ads', NOW());
 RETURN NEW;
END;
$body$
language plpgsql; 