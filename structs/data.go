package structs

type (
	EcowittData struct {
		Absbaromin     float64 `form:"absbaromin" binding:"required"`
		Action         string  `form:"action" binding:"required"`
		Baromin        float64 `form:"baromin" binding:"required"`
		Dailyrainin    float64 `form:"dailyrainin" binding:"required"`
		Dateutc        string  `form:"dateutc" binding:"required"`
		Dewptf         float64 `form:"dewptf" binding:"required"`
		Humidity       float64 `form:"humidity" binding:"required"`
		ID             string  `form:"ID" binding:"required"`
		Indoorhumidity float64 `form:"indoorhumidity" binding:"required"`
		Indoortempf    float64 `form:"indoortempf" binding:"required"`
		Lowbatt        int32   `form:"lowbatt" binding:"required"`
		Monthlyrainin  float64 `form:"monthlyrainin" binding:"required"`
		Password       string  `form:"PASSWORD" binding:"required"`
		Rainin         float64 `form:"rainin" binding:"required"`
		Realtime       bool    `form:"realtime" binding:"required"`
		Rtfreq         int32   `form:"rtfreq" binding:"required"`
		Softwaretype   string  `form:"softwaretype" binding:"required"`
		Solarradiation float64 `form:"solarradiation" binding:"required"`
		Tempf          float64 `form:"tempf" binding:"required"`
		UV             int32   `form:"UV" binding:"required"`
		Weeklyrainin   float64 `form:"weeklyrainin" binding:"required"`
		Windchillf     float64 `form:"windchillf" binding:"required"`
		Winddir        int32   `form:"winddir" binding:"required"`
		Windgustmph    float64 `form:"windgustmph" binding:"required"`
		Windspeedmph   float64 `form:"windspeedmph" binding:"required"`
		Yearlyrainin   float64 `form:"yearlyrainin" binding:"required"`
	}

	WUParams struct {
		Absbaromin     float64 `url:"absbaromin, omitempty"`
		Action         string  `url:"action, omitempty"`
		Baromin        float64 `url:"baromin, omitempty"`
		Dailyrainin    float64 `url:"dailyrainin, omitempty"`
		Dateutc        string  `url:"dateutc, omitempty"`
		Dewptf         float64 `url:"dewptf, omitempty"`
		Humidity       float64 `url:"humidity, omitempty"`
		ID             string  `url:"ID, omitempty"`
		Indoorhumidity float64 `url:"indoorhumidity, omitempty"`
		Indoortempf    float64 `url:"indoortempf, omitempty"`
		Lowbatt        int32   `url:"lowbatt, omitempty"`
		Monthlyrainin  float64 `url:"monthlyrainin, omitempty"`
		Password       string  `url:"PASSWORD, omitempty"`
		Rainin         float64 `url:"rainin, omitempty"`
		Realtime       bool    `url:"realtime, omitempty"`
		Rtfreq         int32   `url:"rtfreq, omitempty"`
		Softwaretype   string  `url:"softwaretype, omitempty"`
		Solarradiation float64 `url:"solarradiation, omitempty"`
		Tempf          float64 `url:"tempf, omitempty"`
		UV             int32   `url:"UV, omitempty"`
		Weeklyrainin   float64 `url:"weeklyrainin, omitempty"`
		Windchillf     float64 `url:"windchillf, omitempty"`
		Winddir        int32   `url:"winddir, omitempty"`
		Windgustmph    float64 `url:"windgustmph, omitempty"`
		Windspeedmph   float64 `url:"windspeedmph, omitempty"`
		Yearlyrainin   float64 `url:"yearlyrainin, omitempty"`
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
		Baromin float64 `url:"baromin, omitempty"`
		//Dateutc      string `url:"dateutc, omitempty"` #TODO While not needed should probably send this
		Dewpoint     float64 `url:"dewpoint, omitempty"`
		Humidity     float64 `url:"humidity, omitempty"`
		StationID    string  `url:"station, omitempty"`
		Rainin       float64 `url:"rainin, omitempty"`
		Tempf        float64 `url:"tempf, omitempty"`
		UV           int32   `url:"uv, omitempty"`
		Winddir      int32   `url:"winddir, omitempty"`
		Windgustmph  float64 `url:"windgustmph, omitempty"`
		Windspeedmph float64 `url:"windspeedmph, omitempty"`
	}
)
