-- Создание таблицы positions
CREATE TABLE IF NOT EXISTS positions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    hourly_rate FLOAT NOT NULL
);

-- Создание таблицы employees
CREATE TABLE IF NOT EXISTS employees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    position_id INT,
    FOREIGN KEY (position_id) REFERENCES positions(id)
);

-- Создание таблицы tasks
CREATE TABLE IF NOT EXISTS tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

-- Создание таблицы timesheets
CREATE TABLE IF NOT EXISTS timesheets (
    id INT AUTO_INCREMENT PRIMARY KEY,
    employee_id INT,
    task_id INT,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES employees(id),
    FOREIGN KEY (task_id) REFERENCES tasks(id)
);

-- Создание триггера для проверки пересечения временных рядов при INSERT
DROP TRIGGER IF EXISTS before_timesheet_insert;
CREATE TRIGGER before_timesheet_insert
BEFORE INSERT ON timesheets
FOR EACH ROW
BEGIN
    DECLARE overlap_count INT;

    -- Проверка пересечения временных рядов
    SET overlap_count = (
        SELECT COUNT(*)
        FROM timesheets t
        WHERE t.employee_id = NEW.employee_id
          AND (
              (t.start_time BETWEEN NEW.start_time AND NEW.end_time) OR
              (t.end_time BETWEEN NEW.start_time AND NEW.end_time) OR
              (NEW.start_time BETWEEN t.start_time AND t.end_time) OR
              (NEW.end_time BETWEEN t.start_time AND t.end_time)
          )
    );

    -- Если найдены пересечения, вызываем ошибку
    IF overlap_count > 0 THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Timesheet overlap detected';
    END IF;
END;

-- Создание триггера для проверки пересечения временных рядов при UPDATE
DROP TRIGGER IF EXISTS before_timesheet_update;
CREATE TRIGGER before_timesheet_update
BEFORE UPDATE ON timesheets
FOR EACH ROW
BEGIN
    DECLARE overlap_count INT;

    -- Проверка пересечения временных рядов
    SET overlap_count = (
        SELECT COUNT(*)
        FROM timesheets t
        WHERE t.employee_id = NEW.employee_id
          AND t.id <> NEW.id
          AND (
              (t.start_time BETWEEN NEW.start_time AND NEW.end_time) OR
              (t.end_time BETWEEN NEW.start_time AND NEW.end_time) OR
              (NEW.start_time BETWEEN t.start_time AND t.end_time) OR
              (NEW.end_time BETWEEN t.start_time AND t.end_time)
          )
    );

    -- Если найдены пересечения, вызываем ошибку
    IF overlap_count > 0 THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Timesheet overlap detected';
    END IF;
END;
