## Get Profile
Request
```
curl --location --request GET 'https://themonkeys.tech/api/v1/profile/user/23' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc2NTg3NTIsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJtYWlsLnRoZW1vbmtleXMubGlmZUBnbWFpbC5jb20ifQ.6tAe7sgLJCkwPsQ9ZZJ_TZaPqVr4E7ZNihjD-0Md1nc'
```

Response
```
{
    "id": 23,
    "firstName": "Jon",
    "lastName": "Doe",
    "email": "john-doe@mail.com",
    "isActive": true,
    "mobile": "none",
    "about": "none",
    "insagram": "none",
    "twitter": "none"
}
```


## Upload Profile details
Request
```
curl --location --request POST 'https://themonkeys.tech/api/v1/profile/user/23' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc2NTg3NTIsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJtYWlsLnRoZW1vbmtleXMubGlmZUBnbWFpbC5jb20ifQ.6tAe7sgLJCkwPsQ9ZZJ_TZaPqVr4E7ZNihjD-0Md1nc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "Dave",
    "last_name": "A",
    "country_code": "+91",
    "mobile": "55665548",
    "about": "The Monkeys inc",
    "instagram": "isndfs",
    "twitter": "adsfddd",
    "email": "new_email@mmail.com"
}'
```

Response
```
{
    "status": 200
}
```

## Upload Profile Picture
Request
```
curl --location 'https://themonkeys.tech/api/v1/profile/user/pic/12' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc2NTg3NTIsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJtYWlsLnRoZW1vbmtleXMubGlmZUBnbWFpbC5jb20ifQ.6tAe7sgLJCkwPsQ9ZZJ_TZaPqVr4E7ZNihjD-0Md1nc' \
--form 'image=@"/home/dave/Pictures/kI92gPS8_400x400.jpg"'
```

Response
```
   "uploaded" 
```

## Get Profile pics
Request
```
curl --location --request GET 'https://themonkeys.tech/api/v1/profile/user/pic/23' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk0NzkwNjksImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJtYWlsLnRoZW1vbmtleXMubGlmZUBnbWFpbC5jb20ifQ.7O9NHabJqTHowQorUZ2ToSjkkgyGg3RljjLm07RBHf4'
```

Response
```
Binary file
```