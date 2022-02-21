## About 

#### Go Lang crud application, It has three APIs which helps to fetch data from database, update database and post database.  

## Setup

* "Setup mongo in local"
* git clone "https://gitlab.com/bansaltushar014/Golang_Crud"
* cd Golang_Crud
* go run main.go 


## APIs

##### Fetch data from mongoDB
`GET`
> localhost:9000?Title=ramayan

##### Add data to mongoDB
`POST - JSON in raw body (postman)`
> localhost:9000
```JSON
{
    "Title": "ramayan",
    "Body": "legacy"
}
```

##### Update book through ID
`Update`
> localhost:9000
```JSON
{
    "Title": "ramayan",
    "Body": "legacy2"
}
```