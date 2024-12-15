# REST API in Go
a REST api in go using Echo and SQLite

# api:

## `POST` `/register`:
registers a user with given data

**required params:**
- `name`
- `email`
- `password`

**returns:**
- `message`: `"User created."`
- `token`: a JWT

## `POST` `/login`:
log in with given credentials

**required params:**
- `email`
- `password`


**returns (if authenticated successfully):**
- `message`: `"Logged in successfully."`
- `token`: a JWT


## `GET` `/profile`:
test authenticated route

**required params:**
- JWT as a Bearer token (authorization header)

**returns (if authenticated):**
- `message`: `"Welcome, [user's name]!"`
