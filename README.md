# Система управления записями

### 1. Клонирование

```bash
git clone https://github.com/nousees/kafka_hospital.git
cd <repository-name>
```

### 2. Настройка переменных

Скопируйте и настройте `.env`:

```bash
cp .env.example .env
```

### 3. Запуск

Соберите и запустите:

```bash
docker-compose up --build
```

Проверьте статус:

```bash
docker ps
```

### 4. Запросы к API

API доступно на порту `8080`. Запросы приведены в формате Curl.

#### Создать пациента

  ```bash
  curl -X POST http://localhost:8080/patients -H "Content-Type: application/json" -d '{"name": "Иванов Иван Иванович"}'
  ```
- **Ожидаемый ответ**:
  ```json
  {"message": "Patient created"}
  ```

#### Создать врача

  ```bash
  curl -X POST http://localhost:8080/doctors -H "Content-Type: application/json" -d '{"name": "Доктор Докторов"}'
  ```
- **Ожидаемый ответ**:
  ```json
  {"message": "Doctor created"}
  ```

#### Создать запись

  ```bash
  curl -X POST http://localhost:8080/appointment -H "Content-Type: application/json" -d '{"patient_id": 1, "doctor_id": 1, "appointment_date": "2025-05-10"}'
  ```
- **Ожидаемый ответ**:
  ```json
  {"message": "Appointment created"}
  ```

#### Поиск записей

  ```bash
  curl http://localhost:8080/search
  ```

- **Ожидаемый ответ**:
  ```json
  [
    {
      "appointment_date": "2025-05-10",
      "patient_name": "Терапевт Петров",
      "doctor_name": "Иванов И. И."
    },
    {
      "appointment_date": "2025-05-10",
      "patient_name": "Стоматолог Иванов",
      "doctor_name": "Антонов А. А."
    }
  ]
  ```

#### Отчет: Топ врачей

  ```bash
  curl http://localhost:8080/reports/top-doctors
  ```
  
- **Ожидаемый ответ**:
  ```json
  [
    {
      "count": 2,
      "name": "Терапевт Петров"
    },
    {
      "count": 3,
      "name": "Стоматолог Иванов"
    }
  ]
  ```

#### Отчет: Топ пациентов


  ```bash
  curl http://localhost:8080/reports/top-patients
  ```

- **Ожидаемый ответ**:
  ```json
  [
    {
      "count": 3,
      "name": "Олег"
    },
    {
      "count": 1,
      "name": "Антонов А. А."
    },
  ]
  ```

#### Отчет: Количество записей на день

  ```bash
  curl http://localhost:8080/reports/daily-counts
  ```
- **Ожидаемый ответ**:
  ```json
  [
    {
      "count": 1,
      "date": "2025-05-10"
    },
    {
      "count": 5,
      "date": "2025-05-15"
    },
  ]
  ```

### 6. Остановка

```bash
docker-compose down
```

### 7. Примечания

- Топики Kafka: `appointments`, `patients`, `doctors`.
- Проверить топики:
  ```bash
  docker exec -it <repository-name>_kafka_1 kafka-topics --bootstrap-server localhost:9092 --list
  ```
