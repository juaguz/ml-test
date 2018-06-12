## Precondiciones
- El estado inicial es en el dia 0, en el angulo 0
- El año esta comprendido por 360 dias, lo que tarda el planeta mas cercano al sol en completar una vuelta
- El pico maximo de lluvia y lluevia se consideran ambos como "Lluvia"


## Como Usar

```
go get github.com/juaguz/ml-test
```

```
go build -o main main.go
```

* El programa soporta los siguientes parametros:
  * -api=true   para correr el servidor web
  * -fill=true  para llenar la base de datos
  * -print=true para imprmir los resultados de las prediciones

* La aplicacion guarda y muestra los datos de una instancia de postgres. La configuracion se realiza por variables de entornos

*Ejemplo
```
WEATHER_DATABASE=postgres://postgres:ml@localhost:5432/weather?sslmode=disable WEATHER_DRIVER=postgres main -print=true -api=true -fill=true
```