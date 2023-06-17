CREATE TABLE IF NOT EXISTS hotel
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(128) NOT NULL,
    description TEXT,
    facilities  TEXT         NOT NULL,
    address     VARCHAR(256) NOT NULL,
    price       decimal      not null,
    rating      decimal      not null
);

CREATE TABLE IF NOT EXISTS booking
(
    id SERIAL PRIMARY KEY,
    hotelId INTEGER REFERENCES hotel (id),
    arrival TIMESTAMP not null,
    departure TIMESTAMP not null
);

CREATE TABLE IF NOT EXISTS review
(
    id SERIAL PRIMARY KEY,
    hotelId INTEGER REFERENCES hotel (id),
    body TEXT not null,
    rating decimal not null
)