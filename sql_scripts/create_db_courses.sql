-- -----------------------------------------------------
-- Schema db_courses
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `db_courses` DEFAULT CHARACTER SET utf8 ;
USE `db_courses` ;

-- -----------------------------------------------------
-- Table `db_courses`.`modules`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_courses`.`modules` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`));


-- -----------------------------------------------------
-- Table `db_courses`.`teachers`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_courses`.`teachers` (
  `dni` INT NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  `lastname` VARCHAR(45) NOT NULL,
  `email` VARCHAR(45) NOT NULL,
  `celphone` INT NOT NULL,
  PRIMARY KEY (`dni`));


-- -----------------------------------------------------
-- Table `db_courses`.`courses`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_courses`.`courses` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `module_id` INT NOT NULL,
  `start_date` DATE NOT NULL,
  `end_date` DATE NOT NULL,
  `dni_teacher` INT NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `module_id`
    FOREIGN KEY (`module_id`)
    REFERENCES `db_courses`.`modules` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `dni_teacher`
    FOREIGN KEY (`dni_teacher`)
    REFERENCES `db_courses`.`teachers` (`dni`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


-- -----------------------------------------------------
-- Table `db_courses`.`students`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_courses`.`students` (
  `file` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `last_name` VARCHAR(45) NOT NULL,
  `birth_date` DATE NOT NULL,
  PRIMARY KEY (`file`));


-- -----------------------------------------------------
-- Table `db_courses`.`registration`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_courses`.`registrations` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `id_course` INT NOT NULL,
  `student_file` INT NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `id_course`
    FOREIGN KEY (`id_course`)
    REFERENCES `db_courses`.`courses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `student_file`
    FOREIGN KEY (`student_file`)
    REFERENCES `db_courses`.`students` (`file`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


INSERT INTO `db_courses`.`students` (name, last_name, birth_date) VALUES
('Juan', 'Pérez', '1990-01-15'),
('María', 'González', '1992-05-23'),
('Carlos', 'Rodríguez', '1988-11-30');


INSERT INTO `db_courses`.`teachers` (dni, name, lastname, email, celphone) VALUES
(20445566, 'Alberto','Prieto', 'albert_perez@hotmail.com',351333444),
(20557799, 'Claudio', 'Paz', 'claudio_paz@gmail.com', 354322244),
(20445344, 'Monica', 'Fuentes', 'monica_fuentes@yahoo.com', 322244444);

INSERT INTO `db_courses`.`modules` (name) VALUES
('Informatica 1'),
('Electronica Aplicada 1'),
('Analisis Matematico 1');

INSERT INTO `db_courses`.`courses` (module_id, start_date, end_date, dni_teacher) VALUES
(1, '2024-06-10', '2024-12-11', 20445566),
(2, '2024-06-10', '2024-12-11', 20557799),
(3, '2024-06-10', '2024-12-11', 20445344);

INSERT INTO `db_courses`.`registrations` (id_course, student_file) VALUES
(1, 1),
(2, 2),
(3, 3);