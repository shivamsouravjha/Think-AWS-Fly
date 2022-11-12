# PiePay

**Tech Used**
* Golang 
* Sentry
* Elasticsearch

**Postman Collection**
* [Import from the link](https://www.getpostman.com/collections/09fbd18d2eee5e446a64)

**How to start the project?**
* ```go run main.go```

**Basic Functionalities**
* Goroutine to fetch data from Youtube API every 10 second and upload on Elasticsearch.
* ```/get``` to get paginated data from Elasticsearch sorted on basis of dates(latest first).
* ```/search``` to search data on basis of description and title ( ```A video with 'title How to make tea?'  matches for the search query 'tea how'```)  
* Multiple key support (in case one expires)
* Sentry to log files and track errors.

**Env values**
  ```APP_ENV=PROD
    SENTRY_DSN=
    SENTRY_SAMPLING_RATE=0.8
    SERVER_PORT=4000
    ES_URL=
    Key=
    YoutubeKey=
    Index=
```
**Sentry Dashboard**
![Screenshot (429)](https://user-images.githubusercontent.com/60891544/173463417-1ef75f39-3249-41be-a690-638da76a5452.png)
![Screenshot (430)](https://user-images.githubusercontent.com/60891544/173463424-db8c0dc5-811c-418e-a75d-9d3e328e3719.png)
