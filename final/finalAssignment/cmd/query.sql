CREATE TABLE if NOT EXISTS list (
id INTEGER PRIMARY KEY AUTOINCREMENT,
name TEXT(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS task (
id INTEGER PRIMARY KEY AUTOINCREMENT,
text TEXT NOT NULL,
listId INTEGER NOT NULL,
completed INTEGER NOT NULL,
FOREIGN KEY(listId) REFERENCES list(id)
ON DELETE CASCADE 
);

INSERT INTO list (name)
VALUES 
("list1"),
("list2"),
("list3")

INSERT INTO task (text, listId, completed)
VALUES 
("task11", 1, 0),
("task12", 1, 0),
("task13", 1, 0),
("task21", 2, 0),
("task22", 2, 0),
("task23", 2, 0)


SELECT * FROM list

SELECT * FROM task




