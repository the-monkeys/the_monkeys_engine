

## Upload file from blog text editor

Method: POST
Request:
```
curl --location 'https://themonkeys.tech/api/v1/files/post/123' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTI1Nzg0NDMsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJkYXZldHdlZXRsaXZlQGdtYWlsLmNvbSJ9.LOWXYurnPaSWljfygdyY0Xuurn69dwBZcv5z0x9SaMY' \
--form 'file=@"/home/dave/Pictures/kI92gPS8_400x400.jpg"'
```

Response
```
{
    "status": 200,
    "newFileName": "kI92gPS8400x400.jpg"
}
```

## Get file from blog text editor
Method: GET
Request:
```
curl --location 'https://themonkeys.tech/api/v1/files/post/123/kI92gPS8400x400.jpg'
```

Response
```
images/signal2023012623521.jpg
```

## Delete a file
Method: DELETE
Request
```
curl --location --request DELETE 'https://themonkeys.tech/api/v1/files/post/123/kI92gPS8400x400.jpg' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTI1Nzg0NDMsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJkYXZldHdlZXRsaXZlQGdtYWlsLmNvbSJ9.LOWXYurnPaSWljfygdyY0Xuurn69dwBZcv5z0x9SaMY'
```

Response
```
{
    "message": "successfully deleted",
    "status": 200
}
```