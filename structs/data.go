package structs

type (
	EcowittData struct {
		Absbaromin     string `form:"absbaromin" binding:"required"`
		Action         string `form:"action" binding:"required"`
		Baromin        string `form:"baromin" binding:"required"`
		Dailyrainin    string `form:"dailyrainin" binding:"required"`
		Dateutc        string `form:"dateutc" binding:"required"`
		Dewptf         string `form:"dewptf" binding:"required"`
		Humidity       string `form:"humidity" binding:"required"`
		ID             string `form:"ID" binding:"required"`
		Indoorhumidity string `form:"indoorhumidity" binding:"required"`
		Indoortempf    string `form:"indoortempf" binding:"required"`
		Lowbatt        string `form:"lowbatt" binding:"required"`
		Monthlyrainin  string `form:"monthlyrainin" binding:"required"`
		Password       string `form:"PASSWORD" binding:"required"`
		Rainin         string `form:"rainin" binding:"required"`
		Realtime       string `form:"realtime" binding:"required"`
		Rtfreq         string `form:"rtfreq" binding:"required"`
		Softwaretype   string `form:"softwaretype" binding:"required"`
		Solarradiation string `form:"solarradiation" binding:"required"`
		Tempf          string `form:"tempf" binding:"required"`
		UV             string `form:"UV" binding:"required"`
		Weeklyrainin   string `form:"weeklyrainin" binding:"required"`
		Windchillf     string `form:"windchillf" binding:"required"`
		Winddir        string `form:"winddir" binding:"required"`
		Windgustmph    string `form:"windgustmph" binding:"required"`
		Windspeedmph   string `form:"windspeedmph" binding:"required"`
		Yearlyrainin   string `form:"yearlyrainin" binding:"required"`
	}

	WeatherParamsStruct struct {
		Action         string `url:"action"`
		Baromin        string `url:"baromin"`
		Dailyrainin    string `url:"dailyrainin"`
		Dateutc        string `url:"dateutc"`
		Dewptf         string `url:"dewptf"`
		Humidity       string `url:"humidity"`
		ID             string `url:"ID"`
		Indoorhumidity string `url:"indoorhumidity"`
		Indoortempf    string `url:"indoortempf"`
		Password       string `url:"PASSWORD"`
		Rainin         string `url:"rainin"`
		Softwaretype   string `url:"softwaretype"`
		Solarradiation string `url:"solarradiation"`
		Tempf          string `url:"tempf"`
		UV             string `url:"UV"`
		Winddir        string `url:"winddir"`
		Windgustmph    string `url:"windgustmph"`
		Windspeedmph   string `url:"windspeedmphurl"`
	}
)
