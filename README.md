# API in GO

## How to use

- `git clone https://github.com/lucasmarques73/api-go.git`
- `cp .env.examlple .env`
- `docker-compose up --build`
- `docker-compose exec web go run Infra/creatingDatabase.go`

## Routes

### Unprotected

#### Home
GET `/` http://localhost/
#### Login
POST `/auth/login` http://localhost/auth/login

### Protected via JWT

To access add in header
` Bearer Token :token`

#### List of users
GET `/users` http://localhost/users

#### Data of one user
GET `/users/:id` http://localhost/users/:id

#### List of widgets
GET `/widgets` http://localhost/widgets

#### Data of one widget
GET `/widgets/:id` http://localhost/widgets/:id

#### Creating new widget
POST `/widgets` http://localhost/widgets

#### Updating existing widgets
PUT `/widgets/:id` http://localhost/widgets/:id

## Structs

### Login
```JSON
{
 "email": "email@example.com",
 "password":"secret"
}
```

### User

```JSON
{
    "name": "Todd Weaver",
    "avatar": "https://loremflickr.com/320/240/cats",
    "email": "email@example.com",
    "pass": "secret"
}
```

### Widget
```JSON
{
    "name": "Helen Lee",
    "color": "Aquamarine",
    "price": 60.4,
    "inventory": 12
}
```

## Password Encrypt

The golang.org/x/crypto/bcrypt package was used to encrypt passwords using bcrypt algorithm

Ex.:
```
secret === $2a$10$4Ndqo.lqBHhaZUTPc2L7veU1Xb9JBHGjD74CfSHl2aeOxN3tx1jD2