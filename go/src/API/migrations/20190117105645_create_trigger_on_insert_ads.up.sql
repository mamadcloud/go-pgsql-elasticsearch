CREATE TRIGGER log_add_ads_table
  AFTER INSERT
  ON ads
  FOR EACH ROW
  EXECUTE PROCEDURE log_add_ads_table();