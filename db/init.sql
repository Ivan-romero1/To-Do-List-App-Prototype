
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE OR REPLACE PROCEDURE add_task(
    IN task_description TEXT
)
LANGUAGE SQL
AS $$
    INSERT INTO tasks (description)
    VALUES (task_description);
$$;
