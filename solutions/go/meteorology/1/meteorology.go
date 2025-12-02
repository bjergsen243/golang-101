package meteorology
import "fmt"
type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (tu TemperatureUnit) String() string {
    if tu == 0 {
        return "°C"
    }
    return "°F"
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (tu Temperature) String() string {
    unit := tu.unit.String()
	return fmt.Sprintf("%d %s", tu.degree, unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (tu SpeedUnit) String() string {
	if tu == 0 {
        return "km/h"
    }
    return "mph"
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (tu Speed) String() string {
	unit := tu.unit.String()
	return fmt.Sprintf("%d %s", tu.magnitude, unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (tu MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %v, %v%% Humidity", tu.location, tu.temperature.String(),tu.windDirection, tu.windSpeed, tu.humidity)
}