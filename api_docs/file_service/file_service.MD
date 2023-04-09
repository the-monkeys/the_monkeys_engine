

## Upload file from blog text editor

Request:
```
curl --location 'https://themonkeys.tech/api/v1/files/post/0242ac120002' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTE5OTc3NjgsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJkYXZldHdlZXRsaXZlQGdtYWlsLmNvbSJ9.Mjj9S4oOSN77cwIBNkGhdPF1BI88El5OFYfp8M2d_6k' \
--form 'file=@"/home/dave/Pictures/20211007_180313_62.jpg"'
```

Response
```
"uploaded"
```

## Get file from blog text editor

Request:
```
curl --location 'https://themonkeys.tech/api/v1/files/post/0242ac120002/20211007_180313_62.jpg' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTE5OTc3NjgsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJkYXZldHdlZXRsaXZlQGdtYWlsLmNvbSJ9.Mjj9S4oOSN77cwIBNkGhdPF1BI88El5OFYfp8M2d_6k'
```

Response
```
binary
```

## Delete a file
Request
```
curl --location --request DELETE 'https://themonkeys.tech/api/v1/files/post/anmpthjerfiletest/20211007_180313_62.jpg' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk0NzkwNjksImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJtYWlsLnRoZW1vbmtleXMubGlmZUBnbWFpbC5jb20ifQ.7O9NHabJqTHowQorUZ2ToSjkkgyGg3RljjLm07RBHf4'
```

Response
```
{
    "message": "successfully deleted",
    "status": 200
}
```