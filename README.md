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