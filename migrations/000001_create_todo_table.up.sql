CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE todo_lists (
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    public BOOLEAN NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE todo_items (
    id SERIAL NOT NULL UNIQUE,
    description VARCHAR(255),
    done BOOLEAN NOT NULL DEFAULT FALSE,
    due_date TIMESTAMP NOT NULL DEFAULT NOW(),
    priority INT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE user_lists (
    user_id UUID NOT NULL,
    list_id INT NOT NULL,
    PRIMARY KEY (user_id, list_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists(id) ON DELETE CASCADE
);

CREATE TABLE list_items (
    todo_id INT NOT NULL,
    item_id INT NOT NULL,
    PRIMARY KEY (todo_id, item_id),
    FOREIGN KEY (todo_id) REFERENCES todo_lists(id) ON DELETE CASCADE,
    FOREIGN KEY (item_id) REFERENCES todo_items(id) ON DELETE CASCADE
);
