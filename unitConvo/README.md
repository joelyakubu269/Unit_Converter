# Unit Converter

A web-based unit converter built with Go. The application allows users to convert values between different units of measurement through a simple browser interface.

## Features

* Length conversion

  * Metres
  * Kilometres
  * Centimetres

* Weight conversion

  * Kilograms
  * Grams
  * Pounds

* Temperature conversion

  * Celsius
  * Kelvin
  * Fahrenheit

* Input validation

* Error handling

* Generic conversion engine using reusable unit definitions

* Separation of concerns between handlers, services, and unit systems

## Project Structure

```text
.
├── main.go
├── handler.go
├── service.go
├── units.go
├── models.go
├── length.html
├── weight.html
├── temperature.html
└── style.css
```

### Responsibilities

* **main.go** – application entry point and route registration.
* **handler.go** – HTTP request handling and template rendering.
* **service.go** – generic conversion logic.
* **units.go** – unit definitions and conversion systems.
* **models.go** – shared structs and types.

## Installation

Clone the repository:

```bash
git clone <https://github.com/joelyakubu269/Unit_Converter.git>
cd <Unit_Converter>
```

## Running the Application

Start the server:

```bash
go run .
```

The application will start on:

```text
http://localhost:8080
```

Available routes:

```text
/length
/weight
/temperature
```

## Example

Convert:

```text
2 kilometres → metres
```

Result:

```text
2000 metres
```

## Design

The application uses a generic conversion engine based on a shared base unit for each measurement system.

Conversion flow:

```text
Input Unit → Base Unit → Output Unit
```

Example for length conversions:

```text
Kilometres → Metres → Centimetres
```

This approach avoids writing conversion logic for every possible pair of units and makes the system easier to extend.

## Future Improvements

* Add support for additional unit systems

  * Time
  * Speed
  * Area
  * Volume

* Replace unit selection switches with a registry-based system

* Add automated tests

* Improve UI styling and responsiveness

## Author

Joel Yakubu
