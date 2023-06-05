# Реализация 

Было принято решение реализации базы данных через создание отдельных таблиц: Order, item , delivery, payment. Отношения всех таблиц 1к1, исключя 1 Order ко многим Item. В качестве кеша используется встроенная в Go структура данных - map.

## API

## GET - получить главную страницу
````
/api
````

## POST - запрос на получение заказа по id
```
/api/search 
```

Проект построен по Clean Architecture.

# Задание 

## В БД: 

Развернуть локально postgresql

Создать свою бд

Настроить своего пользователя.

Создать таблицы для хранения полученных данных.


## В сервисе:

1. Подключение и подписка на канал в nats-streaming
2. Полученные данные писать в Postgres
3. Так же полученные данные сохранить in memory в сервисе (Кеш)
4. В случае падения сервиса восстанавливать Кеш из Postgres
5. Поднять http сервер и выдавать данные по id из кеша
6. Сделать простейший интерфейс отображения полученных данных, для
их запроса по id    

## Доп инфо:

• Данные статичны, исходя из этого подумайте насчет модели хранения в Кеше и в pg. Модель в файле model.json

• В канал могут закинуть что угодно, подумайте как избежать проблем из-за этого

• Чтобы проверить работает ли подписка онлайн, сделайте себе отдельный скрипт, для публикации данных в канал

• Подумайте как не терять данные в случае ошибок или проблем с сервисом

• Nats-streaming разверните локально ( не путать с Nats )


Бонус задание
1. Покройте сервис автотестами. Будет плюсик вам в карму
2. Устройте вашему сервису стресс тест, выясните на что он способен - 
воспользуйтесь утилитами WRK, Vegeta. Попробуйте оптимизировать код


