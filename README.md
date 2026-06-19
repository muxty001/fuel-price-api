# Fuel Price API

A REST API built with Go, Gin, and PostgreSQL for managing fuel stations and fuel prices.

## Features

- User Registration
- User Login
- JWT Authentication
- Station Management
- Fuel Price Management
- PostgreSQL Database

## Endpoints

POST /register
POST /login

GET /stations
POST /stations
PUT /stations/:id
DELETE /stations/:id

GET /fuel-prices
GET /fuel-prices/:id
GET /fuel-prices-details
POST /fuel-prices
PUT /fuel-prices/:id
