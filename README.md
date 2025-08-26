# Rental App
Веб-приложение для аренды квартир.  
**Стек:** Golang, PostgreSQL, Redis, JWT, React (Vite).
## Функционал
- Регистрация и авторизация пользователей (JWT)
- Личный кабинет арендодателя и арендатора
- Добавление, редактирование и удаление объявлений
- Добавление, редактирование и удаление квартир
- Поиск и фильтрация по параметрам, пагинация
- Кэширование через Redis

## Установка и запуск
1. Клонируем репозиторий:
```bash
git clone https://github.com/scmbr/renting-app
```
2. Создаем .env-файл в основной директории (renting-app)
```bash
POSTGRES_PASSWORD=postgres
POSTGRES_USER=postgres
POSTGRES_DB=postgres
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=redis
PGADMIN_DEFAULT_EMAIL=admin@example.com
PGADMIN_DEFAULT_PASSWORD=password
PGADMIN_LISTEN_ADDRESS=0.0.0.0
```
3. Создаем .env-файл в директории backend (renting-app/backend)
```bash
POSTGRES_PASSWORD = postgres
PASSWORD_SALT=password-salt
JWT_SIGNING_KEY=secret-signing-key
URL=http://localhost:3000
APP_ENV=local
REDIS_PASSWORD=redis
```
4. Получаем бесплатные ключи API Геосаджеста и API Геокодера в https://developer.tech.yandex.ru/services 
5. Создаем .env-файл в директории frontend-web (renting-app/frontend-web)
```bash
VITE_YANDEX_GEOCODER_KEY="ключ API Геокодера"
VITE_API_URL=http://localhost:8000/api
VITE_YANDEX_GEOSUGGEST_API_KEY="ключ API Геосаджеста"
VITE_BACKEND_URL=http://localhost:8000
```
6. Запускаем
```bash
docker-compose up --build
```
## Сервисы
Frontend: 
```
http://localhost:3000
```
Backend API: 
```
http://localhost:8000
```
pgAdmin: 
```
http://localhost:5050(логин/пароль из .env)
```
Redis:
```
localhost:6379
```
## Документация API
```
http://localhost:8000/swagger/index.html
```
## Страницы
### Главная
<img width="3840" height="2392" alt="localhost_3000_yakutsk" src="https://github.com/user-attachments/assets/67085c08-6a78-4de6-9554-4fdeb47b01c2" />

### Страница объявления
<img width="3840" height="4230" alt="localhost_3000_advert_542" src="https://github.com/user-attachments/assets/59eabcba-241c-4eac-b642-6b84f766e8bd" />

