# Go Api

Build a Api which use method Get, Post, Put, Del and including Docker and Database.

### Folders
* `/config`: connect with database.
* `/api`: contains all the methods GET, POST etc which using process the data.
* `/model`: contains all the data which needed in order to do the request.
* `/databse`: contains all the request for the database.

### Build Dependencies

* Go Lang 1.19+

### Database

In order tyo connect your postgres db in the project should file config and you write your details
like `username:password@localhost:port/postgres`

After that you can take the tables from the file `schama.sql` and you create to in your db.
Some you can make yet the data simple you take data from the file `insertData.sql`

### Sawgger
The project have swagger in order to you see the api `URL` which there are in the APIs.

So when you run the app, going in this `url http://localhost:8000/docs/index.html#/` in order to you see all the app.

### Create a request (POST)

Do Post in postman in this url http://localhost:8000/api/createTourism with body from the json file which have in data file.

### Read a request (GET)

Do Get in postman in this url http://localhost:8000/api/readTourism/{id}.

### Update a request (PUT)

Do Put in postman in this url http://localhost:8000/api/updateTourism/{id} with body a different json file.

### Delete  a request (DEL)

Do Del in postman in this url http://localhost:8000/api/deleteTourism/{id}.



