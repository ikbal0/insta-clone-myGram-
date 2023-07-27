# Insta Clone

### Framework
used framework :
1. Gin Gonic
2. ORM GORM

### Others
* jwt-go
* crypto
* swagger
* image server https://github.com/ikbal0/image-server

### Image server
uploaded photo are saved on image server (https://github.com/ikbal0/image-server) and image server give image url to be saved on photo table, delete photo will delete photo and image url on image server. 

### Validation
1 email
- email format validation
- unique index
- email can't be empty

2 username
- unique index
- username can't be empty

3 password
- password can't be empty
- greater than 6 character

4 Age
- can't be empty
- value must greater than 8