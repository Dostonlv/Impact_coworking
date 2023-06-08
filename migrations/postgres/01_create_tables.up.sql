CREATE TYPE room_type AS ENUM ('focus', 'team', 'conference');

CREATE TABLE IF NOT EXISTS room(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    type room_type NOT NULL,
    capacity INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS booking(
    id SERIAL PRIMARY KEY,
    resident VARCHAR(255) NOT NULL,
    room_id INTEGER NOT NULL,
    period tsrange NOT NULL,
    FOREIGN KEY (room_id) REFERENCES room(id)
);


INSERT INTO room (name, type, capacity) VALUES ('mytaxi', 'focus', 1);
INSERT INTO room (name, type, capacity) VALUES ('workly', 'team', 5);
INSERT INTO room (name, type, capacity) VALUES ('express24', 'conference', 15);
