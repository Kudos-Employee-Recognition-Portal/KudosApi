# KudosApi
## Data Objects
#### award
#### user
#### manager
#### admin
## Endpoints
- GET /awards
    - returns JSON list of award objects.
    - ![GET users postman call](documentation/images/postmen/GETAwards.PNG)
- GET /awards/{id}
    - returns JSON award object specified by unique identifier, id.
    - ![GET users postman call](documentation/images/postmen/GETAward.PNG)
- GET /users
    - returns JSON list of user objects.
    - ![GET users postman call](documentation/images/postmen/GETUsers.PNG)
- GET /users/{email}
    - returns JSON user object specified by unique identifier, email.
    - ![GET users postman call](documentation/images/postmen/GETUser.PNG)
- GET /users/admins
    - returns JSON list of admin objects.
    - ![GET users postman call](documentation/images/postmen/GETAdmins.PNG)
- GET /users/admins/{id}
    - returns JSON admin object specified by unique identifier, id.
    - ![GET users postman call](documentation/images/postmen/GETAdmin.PNG)
- GET /users/managers
    - returns JSON list of manager objects.
    - ![GET users postman call](documentation/images/postmen/GETManagers.PNG)
- GET /users/managers/{id}
    - returns JSON manager object specified by unique identifier, id.
    - ![GET users postman call](documentation/images/postmen/GETManager.PNG)
- GET /users/managers/{id}/awards
    - In progress...
    - returns JSON list of award objects specified by unique manager identifier, id.
##
- POST /awards
    - creates a new award with request body data and returns unique row id.
    - FUTURE: will trigger certificate generation and email functionality. 
- POST /users/managers
    - creates a new manager with request body data and returns unique row id.
- POST /users/admins
    - creates a new admin with request body data and returns unique row id.
##
- PUT /users/admins/{id}
    - modifies an admin specified by unique identifier, id, and returns the modified JSON object.
- PUT /users/managers/{id}
    - modifies a manager specified by unique identifier, id, and returns the modified JSON object.
##
- DELETE /awards/{id}
    - deletes the award row specified by unique identifier, id, from the database.
- DELETE /users/admins/{id}
    - deletes the admin row specified by unique identifier, id, from the database.
- DELETE /users/managers/{id}
    - deletes the manager rows specified by unique identifier, id, from the database.