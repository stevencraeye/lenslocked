Web Development with GO - from Jon Calhoun

https://courses.calhoun.io/courses


```bash
docker exec -it lenslocked-db-1 /usr/bin/psql -U baloo -d lenslocked
```



4 steps to secure passwords:
1. Use HTTPS to secure the domain
2. Store hashed passwords. Never store encrypted or plaintext passwords.
3. Add a salt to passwords before hashing.
4. Using time-constant functions during authentication.