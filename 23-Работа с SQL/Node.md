# Работа с SQL

## Маршрут вебинара

- Установка и работа с PostgreSQL
- Подключение к СУБД и настройка пула соединений
- Выполнение запросов
- Транзакции и SQL инъекции

## Цели вебинара

1. Изучить как выполнять запросы к PostgreSQL из Go
2. Понять особенности применения `sql.DB`, `sql.Rows`, `sql.Tx`

## Установка и работа с PostgreSQL

### Работаем с PostgreSQL локально

Устанавливаем сервер из консоли (пример для Ubuntu):

```sh
# обновить пакеты
$ sudo apt-get update

# установить PostgreSQL сервер и клиент
$ sudo apt-get -y install postgresql

# запустить PostgreSQL
$ sudo systemctl start postgresql

# подключиться под пользователем, созданным по умолчанию
$ sudo -u postgres psql
```

### Работаем с PostgreSQL локально через Docker

1. Запускаем контейнер с сервером PostgreSQL:

```docker
docker run -d \
    --name pg \
    -e POSTGRES_PASSWORD=postgres \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v pg_data:/var/lib/postgresql/data \
    -p 5432:5432 \
    postgres
```

2. Ждём немного, пока СУБД поднимется
3. Подключаемся к серверу:
```docker
docker exec -it pg psql -Upostgres -dpostgres
```

### Создание database и пользователя в PostgreSQL

**Работаем в клиенте СУБД из терминала:**

```postgresql
postgres=# create database exampledb;
postgres=# create user otus_user with encrypted password 'otus_password';
postgres=# grant all privileges on database exampledb to otus_user;
```


**Удобный клиент с графическим интерфейсом:**
https://www.pgadmin.org/download/

![alt text](image.png)

## Подкатегории SQL

### Data Definition Language (DDL)

Создание простой таблицы и индекса (файл 001.sql):

```postgresql
create table events (
    id serial primary key,
    owner bigint,
    title text,
    descr text,
    start_date date not null,
    start_time time,
    end_date date not null,
    end_time time
);
create index owner_idx on events (owner);
create index start_idx on events using btree (start_date, start_time);
```

Вýполнение SQL скрипта из консоли:

```sh
psql 'host=localhost user=myuser password=mypass dbname=mydb' < 001.sql
```

### Data Manipulation Language (DML)

Добавление строки:

```postgresql
insert into events(owner, title, descr, start_date, end_date)
values(
    42, 'new year', 'watch the irony of fate' , '2024-12-31' , '2024-12-31'
    )
returning id;
```

Обновление строки:

```postgresql
update events
set end_date = '2025-01-01'
where id = 1;
```

### Data Query Language (DQL)

Получение одной строки:

```postgresql
select * from events where id = 1;
```

Получение нескольких строк:

```postgresql
select id, title, descr
from events
where owner = 42 and start_date = '2020-01-01';
```

| id(bigint)    | title(text)   | descr(text)               |
|----           |-------        |------                     |
|       2       |    new year   |   watch the irony of fate |
|       3       |    prepare ny |   make some olive         |

## Подключаем Go

### Подключение к PostgreSQL из Go

Создание подключения:

```go
import "database/sql"
import _ "github.com/jackc/pgx/stdlib"

func main() {
    dsn := "..."
    db, err := sql.Open("pgx", dsn) // *sql.DB
    if err != nil {
        log.Fatalf( "failed to load driver: %v" , err)
    }
}
// создан пул соединений
```

http://go-database-sql.org/importing.html
http://go-database-sql.org/accessing.html

#### DataSourceName


DSN - строка подключения к базе, содержит все необходимые опции.
Синтаксис DSN зависит от используемой базы данных и драйвера.
Например для PostgreSQL:

```sh
"postgres://myuser:mypass@localhost:5432/mydb?sslmode=verify-full"
```

или

```sh
"user=myuser dbname=mydb sslmode=verify-full password=mypass"
```

- **host** - Сервер базы данных или путь к UNIX-сокету (по-умолчанию *localhost*)
- **port** - Порт базы данных (по-умолчанию 5432)
- **dbname** - Имя базы данных
- **user** - Пользователь в СУБД (по умолчанию - пользователь OS)
- **password** - Пароль пользователь

[Подробнее](https://godoc.org/github.com/lib/pq)


#### Пул соединений

![alt text](image-1.png)

![alt text](image-2.png)

- **sql.DB** - это пул соединений с базой данных. Соединения будут открываться по мере необходимости.
- **sql.DB** - безопасен для конкурентного использования (так же как http.Client)

Настройки пула:

```go
// Макс. число открытых соединений от этого процесса
// (по умолчанию - 0, без ограничений)
db.SetMaxOpenConns(n int)
// Макс. число открытых неиспользуемых соединений
db.SetMaxIdleConns(n int)
// Макс. время жизни одного подключения
db.SetConnMaxLifetime(d time.Duration)
// Макс. время ожидания подключения в пуле
db.SetConnMaxIdleTime(d time.Duration)
```

[Подробнее](http://go-database-sql.org/connection-pool.html)


### Выполнение запросов

```go
query := `insert into events(owner, title, descr, start_date, end_date) values($1, $2, $3, $4, $5)`
result, err := db.ExecContext(ctx, query, 42, "new year", "watch the irony of fate" , "2019-12-31" , "2019-12-31") // sql.Result
if err != nil {
// обработать ошибку
}

// Авто-генерируемый ID (SERIAL)
eventId, err := result.LastInsertId() // int64

// Количество измененных строк
rowsAffected, err := result.RowsAffected() // int64
```

[Подробнее](http://go-database-sql.org/retrieving.html)

### Получение результатов

```go
query := `
 select id, title, descr
 from events
 where owner = $1 and start_date = $2
`

rows, err := db.QueryContext(ctx, query, owner, date)

if err != nil {
// ошибка при выполнении запроса
}

defer rows.Close()

for rows.Next() {
    var id int64
    var title, descr string
    if err := rows.Scan(&id, &title, &descr); err != nil {
    // ошибка сканирования
    }
    // обрабатываем строку
    fmt.Printf( "%d %s %s\n" , id, title, descr)
}
```