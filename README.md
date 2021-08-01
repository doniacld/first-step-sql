# First step with SQL

Building a dummy Golang service using SQL.

## Prerequisites

## How to run


    docker exec app ping db

    
curl -v -X POST http://0.0.0.0:8080/students -d@test/add-student-request.jsoN


### Postgres
    psql -h db -d test_db -U root

INSERT INTO students (first_name, last_name, city, phone, email, address, postcode, date_of_birth) VALUES ('titi', 'toto', 'toto', '05144461974', 'toto', 'toto', 123, '2001-12-16');

## To improve

## Resources

### Postgres

https://hub.docker.com/_/postgres

https://www.calhoun.io/updating-and-deleting-postgresql-records-using-gos-sql-package/

https://towardsdatascience.com/how-to-run-postgresql-using-docker-15bf87b452d4

https://gist.github.com/Kartones/dd3ff5ec5ea238d4c546

https://www.postgresql.org/docs/9.1/app-psql.html

https://docs.tibco.com/pub/bwce/2.6.0/doc/html/GUID-E7191578-7E50-4460-93BA-2FF1621CD5B9.html

Driver
https://github.com/jackc/pgx
https://pkg.go.dev/github.com/jackc/pgx/v4/pgxpool

### Docker
https://docs.docker.com/compose/gettingstarted/

https://docs.docker.com/compose/environment-variables/