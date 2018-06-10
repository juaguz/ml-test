package weather


type WeatherCondition string

const  (
	Rain WeatherCondition = "Lluvia"
	Drought WeatherCondition = "Sequia"
	OptimalCondition WeatherCondition = "Condiciones Optimas"
	UnknowCondition WeatherCondition = "Condiciones desconocidas"
)



type Weather struct {
	Id int64 `json:"-"`
	Day uint `json:"dia"`
	Weather WeatherCondition `json:"clima"`
}



