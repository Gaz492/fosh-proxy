package database

import (
	"context"
	"errors"
	"fmt"
	"fosh-proxy/config"
	"fosh-proxy/structs"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pterm/pterm"
	"os"
)

var (
	Conn    *pgxpool.Pool
	ConnCtx context.Context
)

func Initialization() {
	ConnCtx = context.Background()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.Cfg.Database.User, config.Cfg.Database.Password, config.Cfg.Database.Host, config.Cfg.Database.Port, config.Cfg.Database.Database)
	pterm.Info.Println("Con str:", connStr)
	var err error
	Conn, err = pgxpool.New(ConnCtx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	pterm.Success.Println("Connected to DB")
}

func SaveToDB(wData structs.EcowittData) error {
	insertData := `INSERT INTO sensor_data (
absbaromin, baromin, dailyrainin, dateutc, dewptf, humidity, sensor_id, indoorhumidity, indoortempf, lowbatt, monthlyrainin, rainin, realtime, rtfreq, softwaretype, solarradiation, tempf, UV, weeklyrainin, windchillf, winddir, windgustmph, windspeedmph, yearlyrainin) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24);`

	//Execute INSERT command
	_, err := Conn.Exec(ConnCtx, insertData, wData.Absbaromin, wData.Baromin, wData.Dailyrainin, wData.Dateutc, wData.Dewptf, wData.Humidity, wData.ID, wData.Indoorhumidity, wData.Indoortempf, wData.Lowbatt, wData.Monthlyrainin, wData.Rainin, wData.Realtime, wData.Rtfreq, wData.Softwaretype, wData.Solarradiation, wData.Tempf, wData.UV, wData.Weeklyrainin, wData.Windchillf, wData.Winddir, wData.Windgustmph, wData.Windspeedmph, wData.Yearlyrainin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert data into database: %v\n", err)
		return errors.New(fmt.Sprintf("Unable to insert data into database: %v\n", err))
	}
	return nil
}
