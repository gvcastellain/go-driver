CREATE TABLE folders (
    id SERIAL,
    parent_id INT, 
    name varchar(60) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp, 
    modified_at TIMESTAMP NOT NULL, 
    deleted BOOL NOT NULL DEFAULT false,
    PRIMARY KEY(id),
    CONSTRAINT fk_parent
    FOREIGN KEY(parent_id)
    references folders(id)
)