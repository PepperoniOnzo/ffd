package models

type EnvConfigs struct {
	Version       string `mapstructure:"version"`
	FirebaseAppId string `mapstructure:"firebase_app_id"`
}
