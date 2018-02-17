package space

const earthYear = 31557600

// Planet is the name of the planet
type Planet string

// Age calculates your age for that planet.
func Age(age float64, planet Planet) float64 {
	planets := createPlanetReference()

	// Calculate eatch years
	years := age / float64(earthYear)

	return years / planets[planet]
}

func createPlanetReference() map[Planet]float64 {
	ref := map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	return ref
}
