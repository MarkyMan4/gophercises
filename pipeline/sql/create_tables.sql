create table location (
    id integer primary key autoincrement,
    lat decimal,
    lon decimal
);

create table current_condition (
    id integer primary key autoincrement,
    location_id integer,
    weather_id integer,
    main varchar,
    description varchar,
    icon varchar,
    temp decimal,
    feels_like decimal,
    temp_min decimal,
    temp_max decimal,
    pressure integer,
    humidity integer,
    visibility integer,
    wind_speed decimal,
    wind_deg integer,
    cloud_cover integer,
    dt integer,
    sunrise integer,
    sunset integer,
    timezone integer,
    owm_location_id integer,
    location_name varchar,
    foreign key(location_id) references location(id)
);
