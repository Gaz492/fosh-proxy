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

	WUParams struct {
		Action         string `url:"action, omitempty"`
		Baromin        string `url:"baromin, omitempty"`
		Dailyrainin    string `url:"dailyrainin, omitempty"`
		Dateutc        string `url:"dateutc, omitempty"`
		Dewptf         string `url:"dewptf, omitempty"`
		Humidity       string `url:"humidity, omitempty"`
		ID             string `url:"ID, omitempty"`
		Indoorhumidity string `url:"indoorhumidity, omitempty"`
		Indoortempf    string `url:"indoortempf, omitempty"`
		Password       string `url:"PASSWORD, omitempty"`
		Rainin         string `url:"rainin, omitempty"`
		Softwaretype   string `url:"softwaretype, omitempty"`
		Solarradiation string `url:"solarradiation, omitempty"`
		Tempf          string `url:"tempf, omitempty"`
		UV             string `url:"UV, omitempty"`
		Winddir        string `url:"winddir, omitempty"`
		Windgustmph    string `url:"windgustmph, omitempty"`
		Windspeedmph   string `url:"windspeedmph, omitempty"`
	}
	WOWParams struct {
		Action         string `url:"action, omitempty"`
		Baromin        string `url:"baromin, omitempty"`
		Dailyrainin    string `url:"dailyrainin, omitempty"`
		Dateutc        string `url:"dateutc, omitempty"`
		Dewptf         string `url:"dewptf, omitempty"`
		Humidity       string `url:"humidity, omitempty"`
		ID             string `url:"ID, omitempty"`
		Indoorhumidity string `url:"indoorhumidity, omitempty"`
		Indoortempf    string `url:"indoortempf, omitempty"`
		Password       string `url:"PASSWORD, omitempty"`
		Rainin         string `url:"rainin, omitempty"`
		Softwaretype   string `url:"softwaretype, omitempty"`
		Solarradiation string `url:"solarradiation, omitempty"`
		Tempf          string `url:"tempf, omitempty"`
		UV             string `url:"UV, omitempty"`
		Winddir        string `url:"winddir, omitempty"`
		Windgustmph    string `url:"windgustmph, omitempty"`
		Windspeedmph   string `url:"windspeedmph, omitempty"`
	}
	WindyParams struct {
		Baromin string `url:"baromin, omitempty"`
		//Dateutc      string `url:"dateutc, omitempty"` #TODO While not needed should probably send this
		Dewpoint     string `url:"dewpoint, omitempty"`
		Humidity     string `url:"humidity, omitempty"`
		StationID    string `url:"station, omitempty"`
		Rainin       string `url:"rainin, omitempty"`
		Tempf        string `url:"tempf, omitempty"`
		UV           string `url:"uv, omitempty"`
		Winddir      string `url:"winddir, omitempty"`
		Windgustmph  string `url:"windgustmph, omitempty"`
		Windspeedmph string `url:"windspeedmph, omitempty"`
	}
)
