USE mysql;

SET GLOBAL max_allowed_packet = 1024 * 1024 * 256;

CREATE TABLE ROUTINES (
    id INT PRIMARY KEY,
    timestamp VARCHAR(100) NOT NULL,
    task_name VARCHAR(100) NOT NULL,
    iscompleted BOOLEAN DEFAULT false,
    wasinterrupted BOOLEAN DEFAULT false
);