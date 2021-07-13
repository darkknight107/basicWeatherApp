package types

type coordinates struct {
	lon float32
	lat float32
}

type weatherDescription struct {
	id          int
	main        string
	description string
	icon        string
}

type temperature struct {
	Temp       float32
	feels_Like float32
	temp_min   float32
	temp_max   float32
	pressure   float32
	humidity   float32
}

type wind struct {
	speed float32
	deg   float32
	gust  float32
}

type clouds struct {
	all int
}

type sys struct {
	sysType int
	id      int
	country string
	sunrise int64
	sunset  int64
}

type WeatherResponse struct {
	coord      coordinates
	weather    []weatherDescription
	base       string
	Main       temperature
	visibility int
	wind       wind
	clouds     clouds
	dt         int64
	sys        sys
	timezone   int64
	id         int64
	name       string
	cod        int
}
