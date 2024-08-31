sireng, sinau bareng, buat tracking belajar, tp bisa add friend, chatting, ingetin tmen buat belajar dll

pake ap aja :
golang, clean code, ddd, mysql, redis caching, schema validation, aws s3 bucket, docker dll

server:
    port: 8080
database
    user:
    password:
    protocol:
    path:
    dbname:

jwt:
    secret:

tables:
CREATE TABLE users (
id SERIAL PRIMARY KEY,
username VARCHAR(50) UNIQUE NOT NULL,
password_hashed VARCHAR(60) NOT NULL,
role varchar(20) not null default 'user'
);


CREATE TABLE user_profile (
id SERIAL PRIMARY KEY,
user_id bigint(20) unsigned NOT NULL,
profile_photo_url VARCHAR(255),
first_name VARCHAR(50),
last_name VARCHAR(50),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
bio TEXT,
UNIQUE (user_id),
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE rooms (
id SERIAL PRIMARY KEY,
name VARCHAR(255),
description text,
);

CREATE TABLE room_users (
id SERIAL PRIMARY KEY,
room_id bigint(20) unsigned REFERENCES rooms(id) ON DELETE CASCADE,
user_id bigint(20) unsigned REFERENCES users(id) ON DELETE CASCADE,
UNIQUE (room_id, user_id)
);


CREATE TABLE study_sessions (
    id serial NOT NULL,
    user_id bigint(20) unsigned NOT Null references users(id) on delete cascade,
    name VARCHAR(30) DEFAULT NULL,
    session_start TIMESTAMP NOT NULL,
    session_end TIMESTAMP DEFAULT NULL,
    total_time INT(11) AS (TIMESTAMPDIFF(SECOND, session_start, session_end)) STORED,
    note TEXT DEFAULT NULL,
    PRIMARY KEY (id)
);

alter table study_sessions add constraint pk_study_sessions_users foreign key (user_id) references users(id);
