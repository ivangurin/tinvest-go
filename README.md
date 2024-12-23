# О боте

Бот позволяет рассчитать реальную доходность по всему портфелю включая как открытые, так и уже закрытые позиции. Так же можно увидеть подробную детализацию, как по всему портфелю, так и по каждой отдельной позиции.

Пример работы бота можно увидеть в [телеграмм](https://t.me/tinkoff_invest_robot).

# Как запустить своего бота

1. Создать бота в телеграмм и получить токен;
2. Склонировать репозитарий:

```zsh
git clone https://github.com/ivangurin/tinvest-go
```

3. Создать .env файл в директории с репозитария со следующим содержимым, где вместо %TELEGAM_BOT_TOKEN% указать полученный токен из телеграмм:

```zsh
TINVEST_BOT_TOKEN=%TELEGAM_BOT_TOKEN%
```

4. Выполнить последовательно в директории репозитария:

```zsh
go mod download
```

```zsh
go run ./cmd/bot
```

5. Перейти в бота и написать команду

```zsh
/start
```

и следовать инструкциям бота;

6. Для получения данных по счету написать команду:

```zsh
/accounts
```

7. Следить за обновлениями бота)

# Запуск с помощью докера

1. Запустить

```zsh
docker compose up -d
```

2. Остановить

```zsh
docker compose down
```
