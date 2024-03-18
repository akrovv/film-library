# Фильмотека

```
Возможности:
    - Добавить актера (POST /actor)  
    - Редактировать актера (PUT /actor)
    - Удалить актера (DELETE /actor)
    - Получить актера и всего фильмы (GET /actor)
    - Добавить фильм (POST /movie)
    - Редактировать фильм (PUT /movie)
    - Удалить фильм (DELETE /movie)
    - Получить все фильмы с сортировкой по полю (GET /movie/all)
    - Поиск фильм по фрагменту названия, по фрагменту имени актера (GET /movie)
    - Регистрация (POST /register)
    - Авторизация (POST /login)
```

## Инструкция по запуску
```
git clone https://github.com/akrovv/film-library.git
docker-compose up
```

## Стек технологий
Backend: Golang, PostgreSQL,, Redis, Rest API, docker-compose.

## Особенности
- Различные БД: PostgreSQL, Redis для сессий
- Casbin для ролей (admin, user)
- Mock сервисов для тестирования handler'ов
- CI-CD **(lint+test)**

## Makefile
- Make build - сборка проекта
- Make lint - linter 
- Make test-api - тестирование с покрытием

## Особенности:
```
1. Необходимо дополнительно к запросу прикладывать cookie, чтобы приложение могло распознать роль  
2. Уже заранее был создан администратор  
Curl для входа:  

curl -v -X POST -H "Content-Type: application/json" -d '{"username": "admin", "password": "admin"}' localhost:8080/login  

В дальнейшем нужно будет прикладывать cookie, чтобы выполнять действия от администратор.  
Схожая логика и у пользователей, но их надо регистрировать:  
curl -v -X POST -H "Content-Type: application/json" -d '{"username": "user", "password": "user"}' localhost:8080/register
````