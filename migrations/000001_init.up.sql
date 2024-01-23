CREATE TABLE people (
    person_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    patronymic VARCHAR(50),
    age INT,
    gender_id INT,
    nationality_id INT,
    FOREIGN KEY (nationality_id) REFERENCES nationalities(nationality_id),
    FOREIGN KEY (gender_id) REFERENCES genders(gender_id)
);

CREATE TABLE nationalities (
    nationality_id INT PRIMARY KEY AUTO_INCREMENT,
    nationality_name VARCHAR(50)
);

CREATE TABLE genders (
    gender_id INT PRIMARY KEY AUTO_INCREMENT,
    gender_name VARCHAR(10)
);