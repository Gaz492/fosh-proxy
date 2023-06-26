package config

type (
	Config struct {
		SaveToDatabase bool     `yaml:"saveToDatabase"`
		Database       Database `yaml:"database"`
		//DatabaseType   string   `yaml:"databaseType"`
		StationKey  string   `yaml:"stationKey"`
		EnableRelay bool     `yaml:"enableRelay"`
		Relays      []Relays `yaml:"relays"`
	}
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
	Relays struct {
		Enabled       bool   `yaml:"enabled"`
		Host          string `yaml:"host"`
		Format        string `yaml:"format"`
		StationID     string `yaml:"stationID"`
		StationKey    string `yaml:"stationKey"`
		SendFrequency int64  `yaml:"sendFrequency"`
	}
)
