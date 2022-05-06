SET
    datestyle = dmy;

CREATE TABLE users
(
    id         serial PRIMARY KEY,
    username   varchar(60) UNIQUE                  NOT NULL,
    first_name varchar(70)                         NOT NULL,
    last_name  varchar(70)                         NOT NULL,
    email      varchar(70) UNIQUE                  NOT NULL,
    password   varchar(60)                         NOT NULL,
    birthday   date                                NOT NULL,
    created_at timestamp DEFAULT current_timestamp NOT NULL,
    updated_at timestamp DEFAULT current_timestamp NOT NULL,
    deleted_at timestamp DEFAULT NULL
);

CREATE INDEX users_username_index
    ON users (username);

CREATE TABLE seasons
(
    id         serial PRIMARY KEY,
    name       varchar(20) UNIQUE                  NOT NULL,
    created_at timestamp DEFAULT current_timestamp NOT NULL,
    updated_at timestamp DEFAULT current_timestamp NOT NULL,
    deleted_at timestamp DEFAULT NULL
);

CREATE INDEX season_name_index
    ON seasons (name);

CREATE TABLE phones
(
    id         serial PRIMARY KEY,
    phone      varchar(15) UNIQUE                  NOT NULL,
    owner_id   int                                 NOT NULL,
    created_at timestamp DEFAULT current_timestamp NOT NULL,
    updated_at timestamp DEFAULT current_timestamp NOT NULL,
    deleted_at timestamp DEFAULT NULL,
    CONSTRAINT user_phones FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE RESTRICT
);

CREATE INDEX users_phones_index
    ON phones (phone);

CREATE TABLE friends
(
    id         serial PRIMARY KEY,
    user_a_id  int NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
    user_b_id  int NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
    is_blocked bool DEFAULT False,
    CONSTRAINT unique_friends UNIQUE (user_a_id, user_b_id)
);

CREATE TABLE actions
(
    id         serial PRIMARY KEY,
    name       varchar(70) UNIQUE                  NOT NULL,
    created_at timestamp DEFAULT current_timestamp NOT NULL,
    updated_at timestamp DEFAULT current_timestamp NOT NULL,
    deleted_at timestamp DEFAULT NULL
);

CREATE TABLE genres
(
    id         serial PRIMARY KEY,
    name       varchar(60) UNIQUE                  NOT NULL,
    created_at timestamp DEFAULT current_timestamp NOT NULL,
    updated_at timestamp DEFAULT current_timestamp NOT NULL,
    deleted_at timestamp DEFAULT NULL
);

CREATE INDEX genre_index
    ON genres (name);

CREATE TABLE tags
(
    id         serial PRIMARY KEY,
    name       varchar(60) UNIQUE                  NOT NULL,
    created_at timestamp DEFAULT current_timestamp NOT NULL,
    updated_at timestamp DEFAULT current_timestamp NOT NULL,
    deleted_at timestamp DEFAULT NULL
);

CREATE INDEX tag_index
    ON tags (name);

CREATE TABLE types
(
    id         serial PRIMARY KEY,
    name       varchar(50) UNIQUE                  NOT NULL,
    created_at timestamp DEFAULT current_timestamp NOT NULL,
    updated_at timestamp DEFAULT current_timestamp NOT NULL,
    deleted_at timestamp DEFAULT NULL
);

CREATE INDEX types_index
    ON types (name);

CREATE TABLE studios
(
    id         serial PRIMARY KEY,
    name       varchar(50) UNIQUE                  NOT NULL,
    created_at timestamp DEFAULT current_timestamp NOT NULL,
    updated_at timestamp DEFAULT current_timestamp NOT NULL,
    deleted_at timestamp DEFAULT NULL
);

CREATE INDEX studio_name_index
    ON studios (name);

CREATE TABLE anime_groups
(
    id         serial PRIMARY KEY,
    name       varchar(200) UNIQUE                 NOT NULL,
    created_at timestamp DEFAULT current_timestamp NOT NULL,
    updated_at timestamp DEFAULT current_timestamp NOT NULL,
    deleted_at timestamp DEFAULT NULL
);

CREATE INDEX anime_groups_name_index
    ON anime_groups (name);

CREATE TABLE animes
(
    id           serial PRIMARY KEY,
    name         varchar(200) UNIQUE                 NOT NULL,
    description  text                                NOT NULL,
    release_year int                                 NOT NULL,
    group_id     int REFERENCES anime_groups (id),
    rating       double precision                    NOT NULL CHECK ( rating >= 0.0 AND rating <= 10.0),
    views_number bigint                              NOT NULL DEFAULT 0 CHECK ( views_number >= 0 ),
    type_id      int                                 NOT NULL REFERENCES types (id) ON DELETE RESTRICT,
    studio_id    int                                 NOT NULL REFERENCES studios (id) ON DELETE RESTRICT,
    season_id    int                                 NOT NULL REFERENCES seasons (id) ON DELETE RESTRICT,
    created_at   timestamp DEFAULT current_timestamp NOT NULL,
    updated_at   timestamp DEFAULT current_timestamp NOT NULL,
    deleted_at   timestamp DEFAULT NULL
);

CREATE INDEX anime_name_index
    ON animes (name);

CREATE TABLE user_anime_actions
(
    id           serial PRIMARY KEY,
    user_id      int  NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
    anime_id     int  NOT NULL REFERENCES animes (id) ON DELETE RESTRICT,
    action_id    int  NOT NULL REFERENCES actions (id) ON DELETE RESTRICT,
    is_favourite bool NOT NULL    DEFAULT False,
    rating       double precision DEFAULT NULL CHECK ( rating >= 0.0 AND rating <= 10.0)
);

CREATE TABLE anime_genres
(
    id       serial PRIMARY KEY,
    anime_id int NOT NULL REFERENCES animes (id) ON DELETE RESTRICT,
    genre_id int NOT NULL REFERENCES genres (id) ON DELETE RESTRICT,
    CONSTRAINT unique_anime_genre UNIQUE (anime_id, genre_id)
);

CREATE TABLE anime_tags
(
    id       serial PRIMARY KEY,
    anime_id int NOT NULL REFERENCES animes (id) ON DELETE RESTRICT,
    tag_id   int NOT NULL REFERENCES tags (id) ON DELETE RESTRICT,
    CONSTRAINT unique_anime_tags UNIQUE (anime_id, tag_id)
);

CREATE TABLE comments
(
    id                 serial PRIMARY KEY,
    content            text                                NOT NULL,
    replied_comment_id int       DEFAULT NULL REFERENCES comments (id) ON DELETE RESTRICT,
    owner_id           int                                 NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
    anime_id           int                                 NOT NULL REFERENCES animes (id) ON DELETE RESTRICT,
    created_at         timestamp DEFAULT current_timestamp NOT NULL,
    updated_at         timestamp DEFAULT current_timestamp NOT NULL,
    deleted_at         timestamp DEFAULT NULL
);

CREATE INDEX comment_owner_anime_index
    ON comments (owner_id, anime_id);
