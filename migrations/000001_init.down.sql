-- Удаление триггера
DROP TRIGGER IF EXISTS update_messages_updated_at ON messages;

-- Удаление функции
DROP FUNCTION IF EXISTS update_updated_at_column;

-- Удаление таблицы
DROP TABLE IF EXISTS messages;
