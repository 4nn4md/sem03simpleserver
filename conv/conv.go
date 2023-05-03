package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Formler for kovertering av temperaturer
func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 1.8) + 32
}

func CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func FarhenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * (0.5556)
}

func FarhenheitToKelvin(farhenheit float64) float64 {
	return (farhenheit-32)*(0.55555556) + 273.15
}

func KelvintoCelcius(kelvin float64) float64 {
	return kelvin - 273.15
}

func KelvinToFarhenheit(kelvin float64) float64 {

	return 1.8*(kelvin-273) + 32

}

// De andre konverteringsfunksjonene implementere her
// ...
