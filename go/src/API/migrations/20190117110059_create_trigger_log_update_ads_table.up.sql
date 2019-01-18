CREATE TRIGGER log_update_ads_table
  AFTER UPDATE
  ON ads
  FOR EACH ROW
  EXECUTE PROCEDURE log_update_ads_table();