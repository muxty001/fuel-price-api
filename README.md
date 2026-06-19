# Fuel Price API

A REST API built with Go, Gin and PostgreSQL for managing fuel stations and fuel prices.

## Features

- User Registration
- User Login (JWT Authentication)
- Create Fuel Stations
- Update Fuel Stations
- Delete Fuel Stations
- View Fuel Stations
- Add Fuel Prices
- Delete Fuel Prices
- Filter Fuel Prices By Station

## Technologies

- Go
- Gin
- PostgreSQL
- JWT
- Thunder Client

## Endpoints

POST /register

POST /login

GET /stations

GET /stations/:id

POST /stations

PUT /stations/:id

DELETE /stations/:id

GET /fuel-prices

GET /fuel-prices/:id

GET /fuel-prices/station/:id

GET /fuel-prices-details

POST /fuel-prices

DELETE /fuel-prices/:id