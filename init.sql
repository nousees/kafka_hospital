CREATE TABLE patients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE doctors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE appointments (
    id SERIAL PRIMARY KEY,
    patient_id INT NOT NULL,
    doctor_id INT NOT NULL,
    appointment_date STRING NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES patients(id),
    FOREIGN KEY (doctor_id) REFERENCES doctors(id)
);

INSERT INTO doctors (name) VALUES 
    ('Терапевт Петров'), 
    ('Стоматолог Иванов'), 
    ('Невролог Смирнов'), 
    ('Лор Сидоров'), 
    ('Кардиолог Антонов');

INSERT INTO patients (name) VALUES 
    ('Иванов И. И.'), 
    ('Антонов А. А.'), 
    ('Сидоров. С. С.'), 
    ('Петров П. П.'), 
    ('Смирнов С. С.');

INSERT INTO appointments (patient_id, doctor_id, appointment_date) VALUES
    (1, 1, '2025-03-10'), 
    (2, 2, '2025-03-11'), 
    (3, 1, '2025-03-12'),
    (1, 3, '2025-03-13'), 
    (4, 2, '2025-03-14'), 
    (5, 4, '2025-03-15');