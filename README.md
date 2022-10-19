# Vehicle API

## About

This is a simple API that holds information about vehicles

## Authors

- [Lekan Adetunmbi](https://www.github.com/greazleay)


## Tech Stack

* [Go](https://go.dev/)
* [Gin](https://gin-gonic.com/)
* [PostgreSQL](https://www.postgresql.org/)
* [GORM](https://gorm.io/)

## Installation

```bash
  go get
```

## Running the app

```bash
# production
$ go run src/main.go
```

## Documentation

Full API Documentation is available [here]()

## API Reference

Some of the available routes are listed below:

#### Authentication Routes

##### Auth Login

```http
  POST /auth/login
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required**. Your Valid Email |
| `password` | `string` | **Required**. Your Valid Password |


#### User Routes

##### Register

```http
  POST /users/register
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required**. Valid Email Address|
| `password` | `string` | **Required**. Password |
| `fullName` | `string` | **Required**. User's full name |

##### Get User Info

```http
  GET /users/userinfo
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `access_token`      | `string` | **Required**. Valid Access Token |

#### Make Routes

##### Create Make

```http
  POST /accounts/open-account
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `access_token`      | `string` | **Required**. Valid Access Token |
| `name` | `string` | **Required**. Name of the Make in Request Body|
| `country` | `string` | **Required**. Country of the Make in Request Body|

##### Get All Makes

```http
  GET /makes
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `access_token`      | `string` | **Required**. Valid Access Token |

#### Vehicle Routes

##### Create Vehicle

```http
  POST /accounts/open-account
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `access_token`      | `string` | **Required**. Valid Access Token |
| `model` | `string` | **Required**. Model of the Vehicle in Request Body|
| `category` | `string` | **Required**. Category of the Make in Request Body|

##### Get All Vehicles

```http
  GET /vehicles
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `access_token`      | `string` | **Required**. Valid Access Token |



## License

[MIT](https://choosealicense.com/licenses/mit/)


## 🔗 Links
[![portfolio](https://img.shields.io/badge/my_portfolio-000?style=for-the-badge&logo=ko-fi&logoColor=white)](https://pollaroid.net/)
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/siezes)


## Badges

[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)
[![Language](https://img.shields.io/github/languages/count/greazleay/thrifty-api)](https://github.com/greazleay/thrifty-api/)