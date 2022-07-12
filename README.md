# Usage
1. Start postgresql, rabbitMQ, envoy:
```
make docker-compose
```

2. Create base auth, users
3. Migrations:
```
make migrate-up
```

4. Start app:
```
make docker-services
```

# TODO LIST

## Auth service
- [ ] Реализовать стрим обновления списка пользователей
- [ ] Проверка входных параметров
- [ ] Возвращаемые ошибки
- [ ] Добавить тесты

## Access service
- [ ] Реализовать стрим обновления списка групп
- [ ] Проверка входных параметров
- [ ] Возвращаемые ошибки
- [ ] Добавить тесты

## Internal
- [ ] Рефактор брокера. Возвращать ошибки.

# Dev
При разработке новго сервиса необходимо указать сервис в app/internal/service/service.go func Services():
 ```
    s = append(s, auth.Auth_ServiceDesc)
 ```

