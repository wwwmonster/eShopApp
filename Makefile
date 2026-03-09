server:
# 	nodemon --watch '.**/*.go' --signal SIGTERM --exec APP_ENV=dev go run main.go	


	nodemon --signal SIGTERM --ext go --exec APP_ENV=dev go run main.go	




# 	nodemon --exec APP_ENV=dev go run main.go --signal SIGTERM --ext go,js,json
