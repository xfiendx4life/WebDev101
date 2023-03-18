CREATE EXTENSION pgcrypto;
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(256) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    bio TEXT
);
CREATE TABLE places (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(256) NOT NULL UNIQUE,
    location POINT,
    opens TIME,
    closes TIME,
    additional_info TEXT
);
-- SELECT * FROM places;
CREATE TABLE teams (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(256) NOT NULL UNIQUE,
    logo VARCHAR(256),
    additional_info TEXT
);
CREATE TABLE meetings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(256) NOT NULL UNIQUE,
    time TIMESTAMP,
    additional_info TEXT,
    place_id UUID,
    CONSTRAINT fk_place FOREIGN KEY(place_id) REFERENCES places(id)
);
CREATE TABLE users_meetings (
    user_id UUID,
    meetings_id UUID,
    PRIMARY KEY (user_id, meetings_id),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_meeting FOREIGN KEY(meetings_id) REFERENCES meetings(id)
);
CREATE TABLE users_teams (
    user_id UUID,
    team_id UUID,
    status_id VARCHAR(256) [],
    PRIMARY KEY (user_id, team_id),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_team FOREIGN KEY(team_id) REFERENCES teams(id)
);
INSERT INTO users (name, password, bio)
VALUES (
        'first@m.ru',
        '123',
        'The very first user of the app'
    );
INSERT INTO places (name, location, opens, closes, additional_info)
VALUES (
        'stadium near the house',
        point(55.751396, 37.613341),
        '10:00:00',
        '22:00:00',
        'The best place to rest in peace'
    );
INSERT INTO teams (name, logo, additional_info)
VALUES (
        'kids from the streets',
        'https://upload.wikimedia.org/wikipedia/en/1/1f/FC_Chertanovo_Moscow_logo.png',
        'First team'
    );
INSERT INTO meetings (name, place_id, time, additional_info)
VALUES (
        'basketball game',
        (
            SELECT id
            FROM places
            WHERE name = 'stadium near the house'
        ),
        NOW(),
        'first game ever'
    );

INSERT INTO users_teams (user_id, team_id, status_id)
VALUES (
        (
            SELECT id
            FROM users
            WHERE name = 'first@m.ru'
        ),
        (
            SELECT id
            FROM teams
            WHERE name = 'kids from the streets'
        ),
        ARRAY ['creator']
    );

INSERT INTO users_meetings (user_id, meetings_id)
VALUES (
        (
            SELECT id
            FROM users
            WHERE name = 'first@m.ru'
        ),
        (
            SELECT id
            FROM meetings
            WHERE name = 'basketball game'
        )
    );

SELECT *
FROM users;
SELECT *
FROM places;
SELECT *
FROM teams;
SELECT id
FROM places
WHERE name = 'stadium near the house';

SELECT users.name AS user, meetings.name AS meeting_name, meetings.time, places.name AS place FROM users
    INNER JOIN users_meetings ON users.id = user_id 
    INNER JOIN meetings ON meetings_id = meetings.id
    INNER JOIN places ON meetings.place_id = places.id ;