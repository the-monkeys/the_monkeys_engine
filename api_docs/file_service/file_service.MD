

## Upload file from blog text editor

Request:
```
curl --location 'https://127.0.0.1:8080/api/v1/files/post/123' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.asddffwer342121.LOWXYurnPaSWljfygdyY0Xuurn69dwBZcv5z0x9SaMY' \
--form 'file=@"/home/dave/Pictures/signal-2023-01-26-235210_002.jpeg"'
```

Response
```
{
    "status": 200,
    "newFileName": "signal20230126235210002.jpeg"
}
```

## Get file from blog text editor

Request:
```
curl --location 'https://127.0.0.1:8080/api/v1/files/post/123/signal20230126235210002.jpeg'
```

Response
```
images/signal2023012623521.jpg
```

## Delete a file
Request
```
curl --location --request DELETE 'https://127.0.0.1:8080/api/v1/files/post/123/signal20230126235210002.jpeg' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTI1Nzg0NDMsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJkYXZldHdlZXRsaXZlQGdtYWlsLmNvbSJ9.LOWXYurnPaSWljfygdyY0Xuurn69dwBZcv5z0x9SaMY'
```

Response
```
{
    "message": "successfully deleted",
    "status": 200
}
```