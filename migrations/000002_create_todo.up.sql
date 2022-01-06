CREATE TABLE lists
(
    id INT auto_increment,
    title VARCHAR (255) NOT NULL,
    description TEXT,

    primary key (id)
);

CREATE TABLE items
(
    id INT auto_increment,
    title VARCHAR (255) NOT NULL,
    text TEXT,
    due_date TIMESTAMP,
    checked TINYINT(1) NOT NULL DEFAULT (0),

    primary key (id)
);

CREATE TABLE users_lists
(
    id INT auto_increment primary key,
    user_id INT NOT NULL,
    list_id INT NOT NULL,

    foreign key (user_id) references users (id) ON DELETE CASCADE,
    foreign key (list_id) references lists (id) ON DELETE CASCADE
);

CREATE TABLE lists_items
(
    id INT auto_increment primary key,
    list_id INT NOT NULL,
    item_id INT NOT NULL,

    foreign key (list_id) references lists (id) ON DELETE CASCADE,
    foreign key (item_id) references items (id) ON DELETE CASCADE
);
