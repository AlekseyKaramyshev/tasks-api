**Golang REST API example.**

---

```zsh
go run .\cmd\server\
2026/02/08 15:12:42 server listening on :8080
```

---

Endpoint table

- /tasks        - **task storage**
- /tasks/${i}   - **specific task**

---

- curl tests ( Good/Bad examples )

# create
```zsh
$ curl -iLH 'Content-type: application/json' -d '{"title":"test1"}' 'localhost:8080/tasks'
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 08 Feb 2026 12:13:08 GMT
Content-Length: 79

{"id":1,"title":"test1","done":false,"created_at":"2026-02-08T15:13:08+03:00"}

$ curl -iLH 'Content-type: application/json' -d '{"name":"test1"}' 'localhost:8080/tasks'
HTTP/1.1 400 Bad Request
Content-Type: application/json
Date: Sun, 08 Feb 2026 12:14:26 GMT
Content-Length: 30

{"error":"title is required"}
```

# list all
```zsh
$ curl -iL 'localhost:8080/tasks'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 08 Feb 2026 12:15:26 GMT
Content-Length: 81

[{"id":1,"title":"test1","done":false,"created_at":"2026-02-08T15:13:08+03:00"}]

$ curl -iIL 'localhost:8080/tasks'
HTTP/1.1 405 Method Not Allowed
Content-Type: application/json
Date: Sun, 08 Feb 2026 12:15:54 GMT
Content-Length: 31
```

# get one
```zsh
$ curl -iL 'localhost:8080/tasks/1'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 08 Feb 2026 12:16:54 GMT
Content-Length: 79

{"id":1,"title":"test1","done":false,"created_at":"2026-02-08T15:13:08+03:00"}

$ curl -iL 'localhost:8080/tasks/2'
HTTP/1.1 404 Not Found
Content-Type: application/json
Date: Sun, 08 Feb 2026 12:16:56 GMT
Content-Length: 27

{"error":"task not found"}
```

# update one
```zsh
$ curl -iLH 'Content-type: application/json' -X PUT -d '{"title":"test1_1"}' 'localhost:8080/tasks/1'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 08 Feb 2026 12:19:11 GMT
Content-Length: 40

{"id":1,"title":"test1_1","done":false}

$ curl -iLH 'Content-type: application/json' -X PUT -d '{"title":"test1_1"}' 'localhost:8080/tasks/2'
HTTP/1.1 404 Not Found
Content-Type: application/json
Date: Sun, 08 Feb 2026 12:19:24 GMT
Content-Length: 27

{"error":"task not found"}

$ curl -iL 'localhost:8080/tasks'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 08 Feb 2026 12:19:50 GMT
Content-Length: 42

[{"id":1,"title":"test1_1","done":false}]
```

# delete one
```zsh
$ curl -iLH 'Content-type: application/json' -X DELETE -d '{"title":"test1_1"}' 'localhost:8080/tasks/1'
HTTP/1.1 204 No Content
Date: Sun, 08 Feb 2026 12:20:18 GMT

$ curl -iLH 'Content-type: application/json' -X DELETE -d '{"title":"test1_1"}' 'localhost:8080/tasks/1'
HTTP/1.1 404 Not Found
Content-Type: application/json
Date: Sun, 08 Feb 2026 12:20:22 GMT
Content-Length: 27

{"error":"task not found"}
```
