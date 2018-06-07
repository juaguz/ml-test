package weather


type WeatherCondition string

const  (
	rain WeatherCondition = "Lluvia"
	drought WeatherCondition = "Sequia"
	rainPeak WeatherCondition = "Pico de Lluvias"
	optimalCondition WeatherCondition = "Condiciones Optimas"
)



type Weather struct {
	Day uint
	Weather WeatherCondition
}


