USE mysql;

SET GLOBAL max_allowed_packet = 1024 * 1024 * 256;

CREATE TABLE ROUTINES (
    id INT PRIMARY KEY,
    timestamp VARCHAR(100) NOT NULL,
    task_name VARCHAR(100) NOT NULL,
    iscompleted BOOLEAN DEFAULT false,
    wasinterrupted BOOLEAN DEFAULT false
);


-- for long running task
CREATE TABLE BIGTASK (
    field_1 VARCHAR(20), 
    field_2 VARCHAR(20), 
    field_3 VARCHAR(20), 
    field_4 VARCHAR(20), 
    field_5 VARCHAR(20), 
    field_6 VARCHAR(20), 
    field_7 VARCHAR(20), 
    field_8 VARCHAR(20), 
    field_9 VARCHAR(20), 
    field_10 VARCHAR(20), 
    field_11 VARCHAR(20), 
    field_12 VARCHAR(20), 
    field_13 VARCHAR(20), 
    field_14 VARCHAR(20), 
    field_15 VARCHAR(20), 
    field_16 VARCHAR(20), 
    field_17 VARCHAR(20), 
    field_18 VARCHAR(20), 
    field_19 VARCHAR(20), 
    field_20 VARCHAR(20), 
    field_21 VARCHAR(20), 
    field_22 VARCHAR(20), 
    field_23 VARCHAR(20), 
    field_24 VARCHAR(20), 
    field_25 VARCHAR(20)
);