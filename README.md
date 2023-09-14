# GO Simple REST API

## Golang Simple REST API using Golang and PostgreSQL

Simple API using Golang. Course from myskill.id

### Environment

Check PostgreSQL config in .env file
Database should be exist
Table will created while server started

```sql
CREATE TABLE IF NOT EXISTS public.students
(
    student_id bigint NOT NULL DEFAULT nextval('students_student_id_seq'::regclass),
    student_name text COLLATE pg_catalog."default" NOT NULL,
    student_age bigint NOT NULL,
    student_address text COLLATE pg_catalog."default" NOT NULL,
    student_phone_no text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT students_pkey PRIMARY KEY (student_id)
)
```

### Start Server

```
go run main.go
```

server will run at http://localhost:8080

### Index

<details>
    <summary>
        <code>GET</code> <code>/</code>
    </summary>

#### response

```javascript
{
    "message": "welcome to simple api gorm auth using golang. course from myskill.id"
}
```
</details>

### Login

<details>
    <summary>
        <code>POST</code> <code>/login</code>
    </summary>

#### body

```javascript
{
    "username" : "admin",
    "password" : "password123"
}
```

#### failed response

```javascript
{
    "message": "anauthorized. invalid password"
}
```

#### success response

```javascript
{
    "message": "success",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQ2NjM4OTYsImlhdCI6MTY5NDY2MzI5NiwiaXNzIjoidGVzdCJ9.TcNz80OvkGGTuIwLcLe1CLCX7HEGAnQ-1okt_KixCPk"
}
```

Use "token" to access other APIs endpoints
</details>