package relays

import (
	"errors"
	"fmt"
	"fosh-proxy/config"
	"fosh-proxy/structs"
	"fosh-proxy/utils"
	"github.com/google/go-querystring/query"
	"github.com/pterm/pterm"
	"io"
	"time"
)

var (
	relayLastSent = make(map[string]time.Time)
)

func RelayHandler(data structs.EcowittData) error {
	for _, relay := range config.Cfg.Relays {
		lastSent := time.Now().Sub(relayLastSent[relay.Host]).Seconds()
		if relay.Enabled && (int64(lastSent) >= relay.SendFrequency) {
			pterm.Debug.Println(fmt.Sprintf("Sending data to %s", relay.Host))
			switch relay.Format {
			case "wunderground":
				pterm.Debug.Println("Formatting data as wunderground")
				err := wuFormat(data, relay)
				if err != nil {
					pterm.Error.Println("Error relaying wunderground data\n", err)
					return err
				}
			case "wow":
				pterm.Debug.Println("Formatting data as wow")
				err := wowFormat(data, relay)
				if err != nil {
					pterm.Error.Println("Error relaying wow data\n", err)
					return err
				}
			case "windy":
				pterm.Debug.Println("Formatting data as windy")
				err := windyFormat(data, relay)
				if err != nil {
					pterm.Error.Println("Error relaying windy data\n", err)
					return err
				}
			default:
				pterm.Error.Println("Unknown data format, must be 'wunderground', 'wow' or 'windy', not relaying data")
				return errors.New("unknown data format, must be 'wunderground', 'wow' or 'windy', not relaying data")
			}
			relayLastSent[relay.Host] = time.Now()
		} else {
			pterm.Debug.Println(fmt.Sprintf("%s is disabled, skipping", relay.Host))
		}
	}
	return nil
}

func wuFormat(data structs.EcowittData, relay config.Relays) error {
	params := structs.WUParams{
		Absbaromin:     data.Absbaromin,
		Action:         data.Action,
		Baromin:        data.Baromin,
		Dailyrainin:    data.Dailyrainin,
		Dateutc:        data.Dateutc,
		Dewptf:         data.Dewptf,
		Humidity:       data.Humidity,
		ID:             data.ID,
		Indoorhumidity: data.Indoorhumidity,
		Indoortempf:    data.Indoortempf,
		Lowbatt:        data.Lowbatt,
		Monthlyrainin:  data.Monthlyrainin,
		Password:       data.Password,
		Rainin:         data.Rainin,
		Realtime:       data.Realtime,
		Rtfreq:         data.Rtfreq,
		Softwaretype:   data.Softwaretype,
		Solarradiation: data.Solarradiation,
		Tempf:          data.Tempf,
		UV:             data.UV,
		Weeklyrainin:   data.Weeklyrainin,
		Windchillf:     data.Windchillf,
		Winddir:        data.Winddir,
		Windgustmph:    data.Windgustmph,
		Windspeedmph:   data.Windspeedmph,
		Yearlyrainin:   data.Yearlyrainin,
	}

	v, _ := query.Values(params)
	queryString := v.Encode()

	request, err := utils.HandleGetRequest(fmt.Sprintf("%s?%s", relay.Host, queryString), map[string][]string{})
	if err != nil {
		pterm.Error.Println("Error sending wunderground request\n", err)
		return err
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		pterm.Error.Println("Error reading request body\n", err)
	}
	pterm.Debug.Println(string(body))
	return nil
}

func wowFormat(data structs.EcowittData, relay config.Relays) error {
	return nil
}

func windyFormat(data structs.EcowittData, relay config.Relays) error {
	params := structs.WindyParams{Baromin: data.Baromin, Dewpoint: utils.FarenheitToCelsius(data.Dewptf), Humidity: data.Humidity, StationID: relay.StationID, Rainin: data.Rainin, Tempf: data.Tempf, UV: data.UV, Winddir: data.Winddir, Windgustmph: data.Windgustmph, Windspeedmph: data.Windspeedmph}

	v, _ := query.Values(params)
	queryString := v.Encode()

	request, err := utils.HandleGetRequest(fmt.Sprintf("%s/%s?%s", relay.Host, relay.StationKey, queryString), map[string][]string{})
	if err != nil {
		pterm.Error.Println("Error sending windy request\n", err)
		return err
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		pterm.Error.Println("Error reading request body\n", err)
	}
	pterm.Debug.Println(string(body))
	return nil
}
