package skyalert

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

type (
	CloudCondition    int
	WindCondition     int
	RainCondition     int
	DarknessCondition int
	AlertCondition    int
)

const (
	// Clouds
	CloudUnknown     CloudCondition = 0
	CloudClear       CloudCondition = 1
	CloudLightClouds CloudCondition = 2
	CloudVeryCloudy  CloudCondition = 3
	CloudDisabled    CloudCondition = 4

	// WindLimit
	WindUnknown   WindCondition = 0
	WindCalm      WindCondition = 1
	WindWindy     WindCondition = 2
	WindVeryWindy WindCondition = 3
	WindDisabled  WindCondition = 4

	// RainFlag
	RainUnknown  RainCondition = 0
	RainDry      RainCondition = 1
	RainDamp     RainCondition = 2
	RainRain     RainCondition = 3
	RainDisabled RainCondition = 4

	// Darkness
	DarknessUnknown   DarknessCondition = 0
	DarknessDark      DarknessCondition = 1
	DarknessLight     DarknessCondition = 2
	DarknessVeryLight DarknessCondition = 3
	DarknessDisabled  DarknessCondition = 4

	// Alert
	NoWeatherAlert AlertCondition = 0
	WeatherAlert   AlertCondition = 1
)

func parseNumber[T int | float64](input []byte) T {
	var zero T

	input = bytes.Trim(input, " ")

	switch any(zero).(type) {
	case int:
		if i, err := strconv.Atoi(string(input)); err == nil {
			return any(i).(T)
		}
	case float64:
		if i, err := strconv.ParseFloat(string(input), 64); err == nil {
			return any(i).(T)
		}
	}

	return zero
}

// Generates a new Data struct from an io.reader.
// See the examples directory
func New(in []byte) *Data {
	zone, _ := time.LoadLocation("Local")

	d := &Data{
		GeneratedAt: time.Now(),

		location: zone,
		file:     in,
	}

	return d
}

// The SkyAlert adheres to the the Boltwood II cloud sensor standard.
//
// Following are the docs from SBIG.
//
// https://diffractionlimited.com/wp-content/uploads/2016/04/Cloud-SensorII-Users-Manual.pdf
//
// +====+==========+============+====================================================+
// |    | Heading  | Col’s      | Meaning                                            |
// |====|==========|============|====================================================|
// | 0  | Date     | 1-10 		| local date yyyy-mm-dd                              |
// |----|----------|------------|----------------------------------------------------|
// | 1  | Time     | 12-22 	    | local time hh:mm:ss.ss (24 hour clock)             |
// |----|----------|------------|----------------------------------------------------|
// | 2  | T 	   | 24 		| temperature units displayed and in this            |
// |    |          |            | data, 'C' for Celsius or 'F' for Fahrenheit        |
// |----|----------|------------|----------------------------------------------------|
// | 3  | V 	   | 26         | wind velocity units displayed and in this          |
// |    |          |            | data, 'K' for km/hr or 'M' for mph or 'm' for m/s  |
// |----|----------|------------|----------------------------------------------------|
// | 4  | SkyT     | 28-33      | sky-ambient temperature                            |
// |----|----------|------------|----------------------------------------------------|
// | 5  | AmbT     | 35-40      | ambient temperature                                |
// |----|----------|------------|----------------------------------------------------|
// | 6  | SenT     | 41-47      | sensor case temperature                            |
// |----|----------|------------|----------------------------------------------------|
// | 7  | Wind     | 49-54      | wind speed                                         |
// |----|----------|------------|----------------------------------------------------|
// | 8  | Hum      | 56-58      | relative humidity in %                             |
// |----|----------|------------|----------------------------------------------------|
// | 9  | DewPt    | 60-65      | dew point temperature                              |
// |----|----------|------------|----------------------------------------------------|
// | 10 | Hea      | 67-69      | heater setting in %                                |
// |----|----------|------------|----------------------------------------------------|
// | 11 | R        | 71         | rain flag (0=dry, 1=rain last minute, 2=rain now)  |
// |----|----------|------------|----------------------------------------------------|
// | 12 | W        | 73         | wet flag (0=dry, 1=wet last minute, 2=wet now)     |
// |----|----------|------------|----------------------------------------------------|
// | 13 | Since    | 75-79      | seconds since the last valid data                  |
// |----|----------|------------|----------------------------------------------------|
// | 14 | Now()    | 81-92      | date/time given as the VB6 Now() function result   |
// |    |          |            | (in days) when Clarity II last wrote this file     |
// |----|----------|------------|----------------------------------------------------|
// | 15 | c        | 94         | cloud condition                                    |
// |    |          |            |   0=unknown, 1=clear, 2=cloudy, 3=veryCloudy       |
// |----|----------|------------|----------------------------------------------------|
// | 16 | w        | 96         | wind condition                                     |
// |    |          |            |   0=unknown, 1=calm, 2=windy, 3=veryWindy          |
// |----|----------|------------|----------------------------------------------------|
// | 17 | r        | 98         | rain condition                                     |
// |    |          |            |   0=unknown, 1=dry, 2=wet, 3=rain                  |
// |----|----------|------------|----------------------------------------------------|
// | 18 | d        | 100        | daylight condition                                 |
// |    |          |            |   0=unknown, 1=dark, 2=light, 3=verylight          |
// |----|----------|------------|----------------------------------------------------|
// | 19 | C        | 102        | roof close, =0 not requested, =1 close requested   |
// |----|----------|------------|----------------------------------------------------|
// | 20 | A        | 104        | alert, =0 when not alerting, =1 when alerting      |
// +----+----------+------------+----------------------------------------------------+
//
//	Sample when the unit was sitting on my server rack:
//	2024-06-11 14:09:57.00 F M 79.6   92.8  93      0      42  66.3   000 0 0 00019 045454.59025 3 1 1 1 1 1
type Data struct {
	Timestamp           time.Time         `json:"timestamp" yaml:"timestamp" xml:"timestamp"`                               // 2024-06-11 14:09:57.00
	TempScale           string            `json:"tempScale" yaml:"tempScale" xml:"tempScale"`                               // F
	WindScale           string            `json:"windScale" yaml:"windScale" xml:"windScale"`                               // M
	SkyTemp             float64           `json:"skyTemp" yaml:"skyTemp" xml:"skyTemp"`                                     // 79.6
	AmbientTemp         float64           `json:"ambientTemp" yaml:"ambientTemp" xml:"ambientTemp"`                         // 92.8
	SensorTemp          float64           `json:"sensorTemp" yaml:"sensorTemp" xml:"sensorTemp"`                            // 93
	WindSpeed           float64           `json:"windSpeed" yaml:"windSpeed" xml:"windSpeed"`                               // 0
	Humidity            int               `json:"humidity" yaml:"humidity" xml:"humidity"`                                  // 42
	DewPoint            float64           `json:"dewPoint" yaml:"dewPoint" xml:"dewPoint"`                                  // 66.3
	DewHeaterPercentage int               `json:"dewHeaterPercentage" yaml:"dewHeaterPercentage" xml:"dewHeaterPercentage"` // 000
	RainFlag            int               `json:"rainFlag" yaml:"rainFlag" xml:"rainFlag"`                                  // 0
	WetFlag             int               `json:"wetFlag" yaml:"wetFlag" xml:"wetFlag"`                                     // 0
	SinceGoodData       int               `json:"secSinceGoodData" yaml:"secSinceGoodData" xml:"secSinceGoodData"`          // 00019
	DaysSinceLastWrite  float64           `json:"daysSinceLastWrite" yaml:"daysSinceLastWrite" xml:"daysSinceLastWrite"`    // 045454.59025
	CloudCondition      CloudCondition    `json:"cloudCondition" yaml:"cloudCondition" xml:"cloudCondition"`                // 3
	WindCondition       WindCondition     `json:"windCondition" yaml:"windCondition" xml:"windCondition"`                   // 1
	RainCondition       RainCondition     `json:"rainCondition" yaml:"rainCondition" xml:"rainCondition"`                   // 1
	DarknessCondition   DarknessCondition `json:"darknessCondition" yaml:"darknessCondition" xml:"darknessCondition"`       // 1
	RoofCloseRequested  int               `json:"roofCloseRequested" yaml:"roofCloseRequested" xml:"roofCloseRequested"`    // 1
	AlertCondition      AlertCondition    `json:"alertCondition" yaml:"alertCondition" xml:"alertCondition"`                // 1

	// used internally to create SinceGoodData
	GeneratedAt time.Time `json:"-" yaml:"-" xml:"-"`
	location    *time.Location
	file        []byte
}

// Parse handles parsing the read one-line data file
// into the Data struct
func (d *Data) Parse() (*Data, error) {
	var err error

	d.GeneratedAt = time.Now()

	d.Timestamp, err = time.ParseInLocation(
		"2006-01-02 15:04:05.00",
		fmt.Sprintf("%s %s", d.file[0:10], d.file[11:22]),
		d.location,
	)
	if err != nil {
		return nil, err
	}
	d.TempScale = string(d.file[23])
	d.WindScale = string(d.file[25])
	d.SkyTemp = parseNumber[float64](d.file[27:32])
	d.AmbientTemp = parseNumber[float64](d.file[34:39])
	d.SensorTemp = parseNumber[float64](d.file[40:46])
	d.WindSpeed = parseNumber[float64](d.file[48:53])
	d.Humidity = parseNumber[int](d.file[55:57])
	d.DewPoint = parseNumber[float64](d.file[59:64])
	d.DewHeaterPercentage = parseNumber[int](d.file[66:68])
	d.RainFlag = parseNumber[int](d.file[70:71])
	d.WetFlag = parseNumber[int](d.file[72:73])
	d.SinceGoodData = parseNumber[int](d.file[74:79])

	d.CloudCondition = CloudCondition(
		parseNumber[int](d.file[93:94]),
	)

	d.WindCondition = WindCondition(
		parseNumber[int](d.file[95:96]),
	)

	d.RainCondition = RainCondition(
		parseNumber[int](d.file[97:98]),
	)

	d.DarknessCondition = DarknessCondition(
		parseNumber[int](d.file[99:100]),
	)

	d.RoofCloseRequested = parseNumber[int](d.file[101:102])

	d.AlertCondition = AlertCondition(
		parseNumber[int](d.file[103:104]),
	)

	return d, nil
}

// WithLocation Sets a non-local location for time parsing
//
// One probably is running this on a server in the same rack
// or a small raspberry pi that has the same timezone as
// is set on the SkyAlert. By default, this assumes a local
// timezone, but can be overridden with this function.
func (d *Data) WithLocation(loc *time.Location) *Data {
	d.location = loc
	return d
}
