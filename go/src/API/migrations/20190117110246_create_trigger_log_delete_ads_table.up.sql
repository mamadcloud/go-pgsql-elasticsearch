CREATE TRIGGER log_delete_ads_table
  AFTER DELETE
  ON ads
  FOR EACH ROW
  EXECUTE PROCEDURE log_delete_ads_table();