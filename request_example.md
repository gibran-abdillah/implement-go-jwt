
## Authentication

### Register
```sh
curl -X POST \
  http://localhost:8080/api/auth/register \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "example_user",
    "password": "secretpassword",
    "name": "John Doe"
}'
```
Expected response : 
```json
{
    "message":"Created",
    "user":{
        "id":127,
        "name":"John Doe",
        "username":"example_user"
    }
 }
```

### Login
```sh
curl -X POST \
  http://localhost:8080/api/auth/login \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "example_user",
    "password": "secretpassword"
}'
```
Expected response : 
```json
{"message":"Success Logged in!","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2OTIyNTYxNzMsInN1YiI6MSwidXNlcm5hbWUiOiJleGFtcGxlX3VzZXIifQ.rdrh_KHJUgH8fXCk6z9_epSS51EiVpmj_p-Rma_Vv5Y"}
```

## User Handler

### List User

```sh

curl \
  http://localhost:8080/api/users \
  -H 'Cookie: Authorization=<your_token>'
```
### Get Me
```
curl \
  http://localhost:8080/api/me \
  -H 'Cookie: Authorization=<your_token>'

```
### Get user 

```sh
curl \
  http://localhost:8080/api/users/48 \
  -H 'Cookie: Authorization=<your_token>'
```

### Create user 
```sh
curl -X POST \
  http://localhost:8080/api/users \
  -H 'Content-Type: application/json' \
  -H 'Cookie: Authorization=<your_token> \
  -d '{
    "username": "new_username",
    "password": "new_password",
    "name": "New User"
}'
```

### Delete user 
```sh

curl -X DELETE \
  http://localhost:8080/api/users/<user_id> \
  -H 'Cookie: Authorization=<your_token>'
```
Expected response : 
```json
{"message":"deleted, bye!"}
```

### Edit user

```sh
curl -X PUT \
  http://localhost:8080/api/users/<user_id> \
  -H 'Content-Type: application/json' \
  -H 'Cookie: Authorization=<your_token>' \
  -d '{
    "username": "new_username",
    "name": "Updated User"
}'
```


