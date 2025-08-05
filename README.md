# date_chose_bot_v2
[![Frontend CI/CD](https://github.com/NRF24l01/date_chose_bot_v2/actions/workflows/frontend.yml/badge.svg)](https://github.com/NRF24l01/date_chose_bot_v2/actions/workflows/frontend.yml)
[![Backend CI/CD](https://github.com/NRF24l01/date_chose_bot_v2/actions/workflows/backend.yml/badge.svg)](https://github.com/NRF24l01/date_chose_bot_v2/actions/workflows/backend.yml)
[![Tg Bot CI/CD](https://github.com/NRF24l01/date_chose_bot_v2/actions/workflows/bot.yml/badge.svg)](https://github.com/NRF24l01/date_chose_bot_v2/actions/workflows/bot.yml)

Сервис для выбора даты пикника, основной функционал через web app в телеграме.

## Запуск в проде
```shell
wget https://raw.githubusercontent.com/NRF24l01/date_chose_bot_v2/main/docker-compose.yml
docker compose up
```
*Обязательно нужен проброс с домена с https на 9002 порт локальной машины*

## Запуск для разработки
```shell
git clone https://github.com/NRF24l01/date_chose_bot_v2
cd date_chose_bot_v2
docker compose -f test.docker-compose.yml up --build
```
*Обязательно нужен проброс с домена с https на 6000 порт локальной машины*