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
    meeting_id UUID,
    user_status VARCHAR(256)[],
    PRIMARY KEY (user_id, meeting_id),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_meeting FOREIGN KEY(meeting_id) REFERENCES meetings(id)
);

CREATE TABLE users_teams (
    user_id UUID,
    team_id UUID,
    status_id VARCHAR(256)[],
    PRIMARY KEY (user_id, team_id),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_team FOREIGN KEY(team_id) REFERENCES teams(id)
);

