CREATE OR REPLACE FUNCTION log_delete_ads_table()
  RETURNS trigger AS
$body$ 
BEGIN
 INSERT INTO delta_update VALUES(DEFAULT, OLD.ad_id, 'delete', 'ads', NOW());
 RETURN NEW;
END;
$body$
language plpgsql; 