package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InsertData(weatherData map[string]interface{}) {
	stmt, prepErr := db.Prepare(`
		insert into current_condition (
			location_id,
			weather_id,
			main,
			description,
			icon,
			temp,
			feels_like,
			temp_min,
			temp_max,
			pressure,
			humidity,
			visibility,
			wind_speed,
			wind_deg,
			cloud_cover,
			dt,
			sunrise,
			sunset,
			timezone,
			owm_location_id,
			location_name
		) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);
	`)

	if prepErr != nil {
		log.Fatal("Failed to prepare insert statement")
	}

	weatherList := weatherData["weather"].([]interface{})
	weather := weatherList[0].(map[string]interface{})
	mainSection := weatherData["main"].(map[string]interface{})
	wind := weatherData["wind"].(map[string]interface{})
	clouds := weatherData["clouds"].(map[string]interface{})
	sys := weatherData["sys"].(map[string]interface{})

	_, insertErr := stmt.Exec(
		1,
		weather["id"],
		weather["main"],
		weather["description"],
		weather["icon"],
		mainSection["temp"],
		mainSection["feels_like"],
		mainSection["temp_min"],
		mainSection["temp_max"],
		mainSection["pressure"],
		mainSection["humidity"],
		weatherData["visibility"],
		wind["speed"],
		wind["deg"],
		clouds["all"],
		weatherData["dt"],
		sys["sunrise"],
		sys["sunset"],
		weatherData["timezone"],
		weatherData["id"],
		weatherData["name"],
	)
}
