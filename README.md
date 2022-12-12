# SIMPLE API MUSIC

## 1. List Music

#### Request Data
- localhost:8080/api/music
- HTTP Request Method : GET

## 2. Search Music

#### Request Data
- localhost:8080/api/music/{id}
- HTTP Request Method : GET

## 3. Add Music

#### Request Data
- localhost:8080/api/music
- HTTP Request Method : POST
- Body : raw/json
- Parameter masukan:
>name : String\
>album: String\
>album_art: String\
>publish_date: Date


## 3. Update Music

#### Request Data
- localhost:8080/api/music/{id}
- HTTP Request Method : PUT
- Body : raw/json
- Parameter masukan:
>name : String\
>album: String\
>album_art: String\
>publish_date: Date


## 4. Delete Music

#### Request Data
- localhost:8080/api/music/{id}
- HTTP Request Method : DELETE
