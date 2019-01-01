package config

import "time"

type Config struct {
	Version         int                    `yaml:"Version"`         // must be 0
	MqttClients     []MqttClientConfig     `yaml:"MqttClient"`      // mandatory: at least 1 must be defined
	InfluxDbClients []InfluxDbClientConfig `yaml:"InfluxDbClients"` // mandatory: at least 1 must be defined
	Converters      []ConverterConfig      `yaml:"Converters"`      // mandatory: at least 1 must be defined
	HttpServer      HttpServerConfig       `yaml:"HttpServer"`      // optional: default Disabled
	Statistics      StatisticsConfig       `yaml:"Statistics"`      // optional: default Disabled
	LogConfig       bool                   `yaml:"LogConfig"`       // optional: default False
	LogWorkerStart  bool                   `yaml:"LogWorkerStart"`  // optional: default False
	LogMqttDebug    bool                   `yaml:"LogMqttDebug"`    // optional: default False
}

type MqttClientConfig struct {
	name              string // defined automatically by map key
	broker            string // mandatory
	user              string // optional: default empty
	password          string // optional: default empty
	clientId          string // optional: default go-mqtt-to-influxdb
	qos               byte   // optional: default 0, must be 0, 1, 2
	availabilityTopic string // optional: default %Prefix%tele/%clientId%/LWT
	topicPrefix       string // optional: default empty
	logMessages       bool   // optional: default False
}

type InfluxDbClientConfig struct {
	name            string        // defined automatically by map key
	address         string        // mandatory
	user            string        // optional: default empty
	password        string        // optional: default empty
	database        string        // optional: default go-mqtt-to-influxdb
	writeInterval   time.Duration // optional: default 200ms
	timePrecision   time.Duration // optional: default 1s
	logLineProtocol bool          // optional: default False
}

type ConverterConfig struct {
	name              string   // defined automatically by map key
	implementation    string   // mandatory
	targetMeasurement string   // optional: default depends on implementation
	mqttTopics        []string // mandatory: at least 1 must be defined
	mqttClients       []string // optional: defaults to all defined clients
	influxDbClients   []string // optional: defaults to all defined clients
	logHandleOnce     bool     // optional: default False
}

type HttpServerConfig struct {
	enabled     bool   // defined automatically if HttpServer section exists
	bind        string // optional: defaults to ::1 (ipv6 loopback)
	port        int    // optional: defaults to 8042
	logRequests bool   // optional:  default False
}

type StatisticsConfig struct {
	enabled bool // defined automatically if Statistics section exists
}

// Read structs are given to yaml for decoding and are slightly less exact in types
type configRead struct {
	Version         *int                        `yaml:"Version"`
	MqttClients     mqttClientConfigReadMap     `yaml:"MqttClients"`
	InfluxDbClients influxDbClientConfigReadMap `yaml:"InfluxDbClients"`
	Converters      converterReadMap            `yaml:"Converters"`
	HttpServer      *httpServerConfigRead       `yaml:"HttpServer"`
	Statistics      *statisticsConfigRead       `yaml:"Statistics"`
	LogConfig       *bool                       `yaml:"LogConfig"`
	LogWorkerStart  *bool                       `yaml:"LogWorkerStart"`
	LogMqttDebug    *bool                       `yaml:"LogMqttDebug"`
}

type mqttClientConfigRead struct {
	Broker            string  `yaml:"Broker"`
	User              string  `yaml:"User"`
	Password          string  `yaml:"Password"`
	ClientId          string  `yaml:"ClientId"`
	Qos               *byte   `yaml:"Qos"`
	AvailabilityTopic *string `yaml:"AvailabilityTopic"`
	TopicPrefix       string  `yaml:"TopicPrefix"`
	LogMessages       *bool   `yaml:"LogMessages"`
}

type mqttClientConfigReadMap map[string]mqttClientConfigRead

type influxDbClientConfigRead struct {
	Address         string `yaml:"Address"`
	User            string `yaml:"User"`
	Password        string `yaml:"Password"`
	Database        string `yaml:"Database"`
	WriteInterval   string `yaml:"WriteInterval"`
	TimePrecision   string `yaml:"TimePrecision"`
	LogLineProtocol *bool  `yaml:"LogLineProtocol"`
}

type influxDbClientConfigReadMap map[string]influxDbClientConfigRead

type converterConfigRead struct {
	Implementation    string   `yaml:"Implementation"`
	TargetMeasurement string   `yaml:"TargetMeasurement"`
	MqttTopics        []string `yaml:"MqttTopics"`
	MqttClients       []string `yaml:"MqttClients"`
	InfluxDbClients   []string `yaml:"InfluxDbClients"`
	LogHandleOnce     *bool    `yaml:"LogHandleOnce"`
}

type converterReadMap map[string]converterConfigRead

type httpServerConfigRead struct {
	Bind        string `yaml:"Bind"`
	Port        *int   `yaml:"Port"`
	LogRequests *bool  `yaml:"LogRequests"`
}

type statisticsConfigRead struct {
	Enabled *bool `yaml:"Enabled"`
}
