data "weather_cities" "cities" { }

resource "weather_server" "east" {
  address = "510551"
  city    = "singapore"
}

output "city" {
  value = "${weather_server.east.city}"
}

output "foo" {
  value = "${data.weather_cities.cities.foo}"
}