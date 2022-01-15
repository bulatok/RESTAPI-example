###Запуск приложение с docker-compose
```bash
    git clone https://github.com/bulatok/task.git
    make start
```
###Про реализацию
Во время запуска приложения поднимается бд(PostgreSQL) в который имеетя всего одно отношение - users(id, name, surname, phone_number)

Есть один хендлер '/', который 
1) При GET запросе возвращает всех юзеров из бд в формате JSON
```
    curl -X GET http://localhost:8080
```
2) При POST запросе добавляет юзера в бд
```
    curl -X POST http://localhost:8080 -H 'Content-Type: application/json' -d '{"name":"Jack", "surname":"Jk12", "phone_num":"+79522947691"}'
```
3) При DELETE запросе удаляет юзера из бд по id
```
    curl -X DELETE http://localhost:8080 -H 'Content-Type: application/json' -d '{"id":1}'
```