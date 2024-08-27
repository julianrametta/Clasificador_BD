-- -----------------------------------------------------
-- Schema db_clasificador
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `db_clasificador` DEFAULT CHARACTER SET utf8 ;
USE `db_clasificador` ;

-- -----------------------------------------------------
-- Table `db_clasificador`.`bd_configs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_clasificador`.`db_configs` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `connection_string` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE);


-- -----------------------------------------------------
-- Table `db_clasificador`.`info_type_col_names`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_clasificador`.`info_type_col_names` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `description` VARCHAR(45) NOT NULL,
  `regex` VARCHAR(70) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE);


-- -----------------------------------------------------
-- Table `db_clasificador`.`info_type_table_names`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_clasificador`.`info_type_table_names` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `description` VARCHAR(45) NOT NULL,
  `regex` VARCHAR(70) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE);

-- -----------------------------------------------------
-- Table `db_clasificador`.`schemas`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_clasificador`.`schemas` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `db_config_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `schemas_bd_config_id`
    FOREIGN KEY (`db_config_id`)
    REFERENCES `db_clasificador`.`db_configs` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION);

-- -----------------------------------------------------
-- Table `db_clasificador`.`tables`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_clasificador`.`tables` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `info_type_table_name_id` INT NOT NULL,
  `schema_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `schema_id`
    FOREIGN KEY (`schema_id`)
    REFERENCES `db_clasificador`.`schemas` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `info_type_table_name_id`
    FOREIGN KEY (`info_type_table_name_id`)
    REFERENCES `db_clasificador`.`info_type_table_names` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


-- -----------------------------------------------------
-- Table `db_clasificador`.`columns`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_clasificador`.`columns` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `info_type_col_name_id` INT NOT NULL,
  `table_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `table_id`
    FOREIGN KEY (`table_id`)
    REFERENCES `db_clasificador`.`tables` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `info_type_col_name_id`
    FOREIGN KEY (`info_type_col_name_id`)
    REFERENCES `db_clasificador`.`info_type_col_names` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);

-- -----------------------------------------------------
-- Table `db_clasificador`.`scans`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_clasificador`.`scans` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `date` DATETIME NOT NULL,
  `username` VARCHAR(45) NOT NULL,
  `db_config_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `scans_bd_config_id`
    FOREIGN KEY (`db_config_id`)
    REFERENCES `db_clasificador`.`db_configs` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);

INSERT INTO `db_clasificador`.`info_type_col_names` (name, description, regex) VALUES
('N/A', 'Tipo de info de col default',''),
('NAME', 'Tipo de info de col para nombres','\\b(NAME)\\b|\\b(Name)\\b|\\b(name)\\b|\\b(lastname)\\b'),
('NAME', 'Tipo de info de col para fechas','\\b(start_date)\\b|\\b(end_date)\\b|\\b(Date)\\b|\\b(birth_date)\\b');


INSERT INTO `db_clasificador`.`info_type_table_names` (name, description, regex) VALUES
('N/A', 'Tipo de info de tabla default',''),
('PERSON', 'Tipo de info de tabla para personas','\\b(teachers)\\b|\\b(Teachers)\\b|\\b(students)\\b|\\b(Students)\\b'),
('COURSE', 'Tipo de info de tabla para cursos','\\b(modules)\\b|\\b(courses)\\b|\\b(Modules)\\b|\\b(Courses)\\b');