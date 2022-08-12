# beearning-test

Thiết kế DB và viết api để đảm bảo curd user.
Làm api login và logout bằng jwt

* API Postman:
1 Register user: 
URL: http://localhost:8080/api/auth/register
Method: POST
Parameter: 
    "name":"username",
    "email":"useremail",
    "password":"userpassword"
    
2 Login user:
URL: http://localhost:8080/api/auth/login
Method: POST
Parameter:
    "email":"useremail",
    "password":"userpassword"

3 Profile User:
URL: http://localhost:8080/api/user/profile
Method: GET
Parameter: No

4 Update Page:
URL: http://localhost:3000//api/pages/id
Method: PUT
Parameter: 
    Header: Key: Authorization
            Value:   "keylogin"
    
5 Update User:
URL: http://localhost:8080/api/user/profile
Method: PUT
Parameter: 
     "name":"username",
    "email":"useremail",
    "password":"userpassword"
