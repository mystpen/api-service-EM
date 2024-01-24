-- +migrate Up
CREATE TABLE nationalities (
    nationality_id SERIAL PRIMARY KEY,
    nationality_name UNIQUE VARCHAR(50)
);

CREATE TABLE genders (
    gender_id SERIAL PRIMARY KEY,
    gender_name UNIQUE VARCHAR(10)
);

CREATE TABLE people (
    person_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    patronymic VARCHAR(50),
    age INT,
    gender_id INT,
    nationality_id INT,
    FOREIGN KEY (nationality_id) REFERENCES nationalities(nationality_id),
    FOREIGN KEY (gender_id) REFERENCES genders(gender_id)
);
