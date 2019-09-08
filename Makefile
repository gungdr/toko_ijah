APP=toko_ijah
SKU_PREFIX=SSI
DB_DRIVER=sqlite3 
DB_CONNECTION=toko_ijah_sqlite.db 
HTTP_PORT=1337
build:
	go get ./...
	go build -o ${APP} main.go
test :
	go test -cover ./...
run :
	SKU_PREFIX=SSI \
	DB_DRIVER=${DB_DRIVER} \
	DB_CONNECTION=${DB_CONNECTION} \
	HTTP_PORT=${HTTP_PORT} \
	go run main.go
clean :
	if [ -f ${APP} ] ; then rm ${APP} ; fi