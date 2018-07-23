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
To get the token use the struct `Login` below  
POST `/auth/login` http://localhost/auth/login

### Protected via JWT

To access add in header
`Authorization: Bearer :token`

#### Authenticated user data
GET `/auth/user` http://localhost/auth/user

#### List of users
GET `/users` http://localhost/users

#### Data from a user

GET `/users/:id` http://localhost/users/:id

#### Creating new user
To create user use the struct `User` below  
POST `/users` http://localhost/users

#### Updating existing users
To update user use the struct `User` below  
PUT `/users/:id` http://localhost/users/:id

#### Deleting existing users
DELETE `/users/:id` http://localhost/users/:id

#### List of widgets
GET `/widgets` http://localhost/widgets

#### Data from a widget
GET `/widgets/:id` http://localhost/widgets/:id

#### Creating new widget
To create widget use the struct `Widget` below  
POST `/widgets` http://localhost/widgets

#### Updating existing widgets
To update widget use the struct `Widget` below  
PUT `/widgets/:id` http://localhost/widgets/:id

#### Deleting existing widgets
DELETE `/widgets/:id` http://localhost/widgets/:id

## Structs

### Login
```JSON
{
 "email": "user@user.com",
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
    "melts": true
    "inventory": 12
}
```

## Password Encrypt

The [golang.org/x/crypto/bcrypt](golang.org/x/crypto/bcrypt) package was used to encrypt passwords using bcrypt algorithm

Ex.:
```
secret === $2a$10$4Ndqo.lqBHhaZUTPc2L7veU1Xb9JBHGjD74CfSHl2aeOxN3tx1jD2
