# Basic api service

Missing:
1. Validation
2. Rate limiting
3. Response writing is not fully done
4. Error handling (struct for error)
5. Middleware logging
6. Configuration subpackage
7. Deployment flow (dockerfile)
8. Make file targets

AND most importantly:
1. Unit tests, mocks, e2e tests (ginko in example)


alex.ermishkin@hubuc.com

# Original task:
```
A link to github with project that has 1 endpoint 
POST /user and accepts JSON with

username  //required
email //required, email
password //required 6-120 characters

and returns user_id in response

With (fake)repository which does a dummy check if such user already exists

What we'd like to see:
 - Good structured code
 - Error handling
 - input data validation
 - as many production-ready features as you can think of (graceful shutdown/logs/etc)
 
 ```

# Basic manual test

 curl -X POST http://localhost:8080/user \
   -H 'Content-Type: application/json' \
   -d '{
  "username": "bob",
  "email": "email@corp.com",
  "password": "secure"
}'

