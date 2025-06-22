package model

type SystemHealthResp struct {
	DBOnline bool `json:"db_online"`
}

type SystemTimeResp struct {
	CurrentTimeUnix int64 `json:"current_time_unix"`
}
