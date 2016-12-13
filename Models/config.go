package models

var Conf Config

type Config struct {
	Port  int   `json:"port"`
	Redis RConf `json:"rconf"`
}

type RConf struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port,omitempty"`
	Password string `json:"password,omitempty"`
}
