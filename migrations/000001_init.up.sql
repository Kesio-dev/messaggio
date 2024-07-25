-- Создание таблицы messages
CREATE TABLE messages
(
    id         SERIAL PRIMARY KEY,
    message    TEXT        NOT NULL,
    status     VARCHAR(20) NOT NULL,
    created_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Функция для обновления поля updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Триггер для автоматического обновления поля updated_at
CREATE TRIGGER update_messages_updated_at
    BEFORE UPDATE
    ON messages
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
