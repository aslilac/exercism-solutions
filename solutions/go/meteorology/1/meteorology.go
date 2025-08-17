package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (self TemperatureUnit) String() string {
    return map[TemperatureUnit]string{
        0: "°C",
        1: "°F",
    }[self]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (self Temperature) String() string {
    return fmt.Sprintf("%v %v", self.degree, self.unit)
}


type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (self SpeedUnit) String() string {
    return map[SpeedUnit]string{
        0: "km/h",
        1: "mph",
    }[self]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (self Speed) String() string {
    return fmt.Sprintf("%v %v", self.magnitude, self.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (self MeteorologyData) String() string {
    return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", self.location, self.temperature, self.windDirection, self.windSpeed, self.humidity)
}
