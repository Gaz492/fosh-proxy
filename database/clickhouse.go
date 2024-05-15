package database

import (
	"context"
	"fmt"
	"fosh-proxy/config"
	"fosh-proxy/structs"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/pterm/pterm"
	"time"
)

type row struct {
	Timestamp      time.Time `ch:"timestamp"`
	StationId      string    `ch:"station_id"`
	SensorId       string    `ch:"sensor_id"`
	AbsBaroMin     float64   `ch:"abs_baro_min"`
	BaroMin        float64   `ch:"baro_min"`
	TempF          float64   `ch:"temp_f"`
	DewpointF      float64   `ch:"dewpoint_f"`
	Humidity       float64   `ch:"humidity"`
	IndoorTempF    float64   `ch:"indoor_temp_f"`
	IndoorHumidity float64   `ch:"indoor_humidity"`
	SolarRadiation float64   `ch:"solar_radiation"`
	Uv             int32     `ch:"uv"`
	WindChillF     float64   `ch:"wind_chill_f"`
	WindDir        int32     `ch:"wind_dir"`
	WindSpeedMph   float64   `ch:"wind_speed_mph"`
	WindGustMph    float64   `ch:"wind_gust_mph"`
	RainIn         float64   `ch:"rain_in"`
	DailyRainIn    float64   `ch:"daily_rain_in"`
	WeeklyRainIn   float64   `ch:"weekly_rain_in"`
	MonthlyRainIn  float64   `ch:"monthly_rain_in"`
	YearlyRainIn   float64   `ch:"yearly_rain_in"`
	LowBattery     int32     `ch:"low_battery"`
	RealTime       bool      `ch:"real_time"`
	RtFreq         int32     `ch:"rt_freq"`
	SoftwareType   string    `ch:"software_type"`
}

func InsertData(wData structs.EcowittData) error {
	conn, err := connect()
	if err != nil {
		pterm.Error.Println("Error connecting to ClickHouse", err)
		return err
	}
	defer conn.Close()
	v, err := conn.ServerVersion()
	pterm.Info.Println("ClickHouse version:", v)

	batch, err := conn.PrepareBatch(context.Background(), "INSERT INTO data")
	if err != nil {
		return err
	}

	//pterm.Info.Println(d)
	err = batch.AppendStruct(&row{
		Timestamp:      time.Now(),
		StationId:      "home",
		SensorId:       wData.ID,
		AbsBaroMin:     wData.Absbaromin,
		BaroMin:        wData.Baromin,
		TempF:          wData.Tempf,
		DewpointF:      wData.Dewptf,
		Humidity:       wData.Humidity,
		IndoorTempF:    wData.Indoortempf,
		IndoorHumidity: wData.Indoorhumidity,
		SolarRadiation: wData.Solarradiation,
		Uv:             wData.UV,
		WindChillF:     wData.Windchillf,
		WindDir:        wData.Winddir,
		WindSpeedMph:   wData.Windspeedmph,
		WindGustMph:    wData.Windgustmph,
		RainIn:         wData.Rainin,
		DailyRainIn:    wData.Dailyrainin,
		WeeklyRainIn:   wData.Weeklyrainin,
		MonthlyRainIn:  wData.Monthlyrainin,
		YearlyRainIn:   wData.Yearlyrainin,
		LowBattery:     wData.Lowbatt,
		RealTime:       wData.Realtime,
		RtFreq:         wData.Rtfreq,
		SoftwareType:   wData.Softwaretype,
	})
	if err != nil {
		pterm.Error.Println("Error appending data to batch", err)
		return err
	}
	return batch.Send()
}

func connect() (driver.Conn, error) {
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{fmt.Sprintf("%s:%d", config.Cfg.Database.Host, config.Cfg.Database.Port)},
			Auth: clickhouse.Auth{
				Database: config.Cfg.Database.Database,
				Username: config.Cfg.Database.User,
				Password: "",
			},
			Debug: true,
			Debugf: func(format string, v ...interface{}) {
				pterm.Warning.Printfln(format, v)
			},
		})
	)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			pterm.Error.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	return conn, nil
}
