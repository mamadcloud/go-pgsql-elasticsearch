# go-pgsql-elasticsearch

## Requirements
1. Make sure you've installed `docker` and `docker-compose`.
2. Make sure you've installed `Go`.
3. Download or clone this `Repo`.
4. `Postman Collection` is in root folder

## Codes
The source code for the API and service to update `ElasticSearch` are in folder `go/src/`.

## Stacks
1. Programming Language: `GOLANG`
2. Database: `Postgresql`
3. Full Text Search Engine: `ElasticSearch`
4. Container: `Docker`

## Run the code outside the Container
1. Make sure to install `dep`.
2. Run command `dep ensure` to download the dependencies.
3. Make sure the container for `Database` is running.
4. On folder `pkg/configs/db/postgres` open file `postgres.go`.
5. Change `host` from `database_postgres` to `localhost
6. Change `port` from `5432` to `5433`

## Run the code inside contaianer
After you download/clone this `Repo`, go to `dockers` folder. Run command `docker-compose up --build -d`

### Access containers
1. API: [`localhost:8001`](http://localhost:8001/)
2. Database: `localhost:5433`
3. ElasticSearch: [`localhost:9201`](http://locahost:9201/)

## Troubleshoot
If there are some changes made to the code for API
- [x] Make sure you are in root folder of API (`go/src/API`)
- [x] Run command `./build.sh`
- [x] Rebuild the containers by running this command `docker-compose up --build -d`

## Update to ElasticSearch on data change
Steps:
1. INSERT/UPDATE/DELETE on `ads` table will trigger a function to add record to `delta_update`.
2. It will store `ads ID` and `action` with value [add|update|delete].
3. A service that running in the background will check on table `delta_update` for every 5 seconds.
4. It will grab 1000 data and pump to ElasticSearch depend on the action on the row.