CREATE TABLE files (
    id serial,
    fodler_id INT, 
    owner_id INT NOT NULL,
    name VARCHAR(200) NOT NULL, 
    type VARCHAR(50) NOT NULL,
    path VARCHAR(250) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    modified_at TIMESTAMP NOT NULL,
    deleted BOOL NOT NULL DEFAULT false,
    PRIMARY KEY (id),
    CONSTRAINT fk_folders
        FOREIGN KEY(fodler_id)
            references folders(id),
    CONSTRAINT fk_owner
        FOREIGN KEY(owner_id)
            REFERENCES users(id)
)