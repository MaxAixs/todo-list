ALTER TABLE todo_items
DROP COLUMN IF EXISTS sent_notify,
DROP COLUMN IF EXISTS sent_analys,
DROP COLUMN IF EXISTS completed_at;
