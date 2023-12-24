



func initRoutesHandler(){

}

func validateStockHandler([] ) {

}


func GetStockHandler( w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	
	// get the query params
	queryParams := r.URL.Query()

	//get the stock symbol from the query params
	stockSymbol := queryParams.Get()
}