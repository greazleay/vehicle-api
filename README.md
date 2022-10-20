# Vehicle API

## About

This is a Simple Go CRUD API that manages information about Vehicles, most routes are protected with JWT Authentication.

## Authors

- [Lekan Adetunmbi](https://www.github.com/greazleay)


## Tech Stack

* [Go](https://go.dev/)
* [Gin](https://gin-gonic.com/)
* [PostgreSQL](https://www.postgresql.org/)
* [GORM](https://gorm.io/)
* [JWT-Go](https://github.com/golang-jwt/jwt)

## Running the app

```bash
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
| `firstName` | `string` | **Required**. User's First name |
| `lastName` | `string` | **Required**. User's Last name |

##### Get All Users

```http
  GET /users
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |

##### Get User By ID

```http
  GET /users/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |
| `id`      | `string` | **Required**. Valid UUID in Request Params|

##### Update User

```http
  PATCH /users/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |
| `id`      | `string` | **Required**. Valid UUID in Request Params|
| `firstName` | `string` | **Required**. User's First name in Request Body|
| `lastName` | `string` | **Required**. User's Last name in Request Body|

##### Delete User

```http
  DELETE /users/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |
| `id`      | `string` | **Required**. Valid UUID in Request Params|

#### Make Routes

##### Create Make

```http
  POST /makes
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |
| `name` | `string` | **Required**. Name of the Make in Request Body|
| `country` | `string` | **Required**. Country of the Make in Request Body|

##### Get All Makes

```http
  GET /makes
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |

##### Get Make By ID

```http
  GET /makes/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |
| `id`      | `string` | **Required**. Valid UUID in Request Params|

##### Get Make By Name

```http
  GET /makes/names?name={name}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |
| `name`      | `string` | **Required**. Make Name in Request Query|

##### Get Make By Country

```http
  GET /makes/countries?country={country}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |
| `country`      | `string` | **Required**. Make Country in Request Query|

##### Update Make

```http
  PATCH /makes/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |
| `id`      | `string` | **Required**. Valid UUID in Request Params|
| `name` | `string` | **Required**. Make Name in Request Body|
| `country` | `string` | **Required**. Make Country in Request Body|

##### Delete Make

```http
  DELETE /makes/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |
| `id`      | `string` | **Required**. Valid UUID in Request Params|

#### Vehicle Routes

##### Create Vehicle

```http
  POST /vehicles
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |
| `model` | `string` | **Required**. Model of the Vehicle in Request Body|
| `category` | `string` | **Required**. Category of the Make in Request Body|

##### Get All Vehicles

```http
  GET /vehicles
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |

##### Get Vehicle By ID

```http
  GET /vehicles/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |
| `id`      | `string` | **Required**. Valid UUID in Request Params|

##### Update Vehicle

```http
  PATCH /vehicles/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |
| `model` | `string` | **Required**. Model of the Vehicle in Request Body|
| `category` | `string` | **Required**. Category of the Make in Request Body|

##### Delete Vehicle

```http
  DELETE /vehicles/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer_token`      | `string` | **Required**. Valid Bearer Token |
| `id`      | `string` | **Required**. Valid UUID in Request Params|

## License

[MIT](https://choosealicense.com/licenses/mit/)


## 🔗 Links
[![portfolio](https://img.shields.io/badge/my_portfolio-000?style=for-the-badge&logo=ko-fi&logoColor=white)](https://pollaroid.net/)
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/siezes)


## Badges

[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)
[![Language](https://img.shields.io/github/languages/count/greazleay/thrifty-api)](https://github.com/greazleay/thrifty-api/)