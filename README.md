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
    id INT(11) NOT NULL AUTO_INCREMENT,
    username VARCHAR(20) NOT NULL UNIQUE,
    password_hashed VARCHAR(60) NOT NULL,
    profile_photo_url VARCHAR(255) DEFAULT NULL,
    role varchar(20) NOT NULL DEFAULT "user",
    first_name VARCHAR(50) DEFAULT NULL,
    last_name VARCHAR(50) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    bio TEXT DEFAULT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE study_sessions (
    id INT(11) NOT NULL AUTO_INCREMENT,
    user_id INT(11) NOT NULL,
    name VARCHAR(30) DEFAULT NULL,
    session_start TIMESTAMP NOT NULL,
    session_end TIMESTAMP DEFAULT NULL,
    total_time INT(11) AS (TIMESTAMPDIFF(SECOND, session_start, session_end)) STORED,
    note TEXT DEFAULT NULL,
    PRIMARY KEY (id),
);

alter table study_sessions add constraint pk_study_sessions_users foreign key (user_id) references users(id);
