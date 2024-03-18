CREATE TYPE gender AS ENUM ('Male', 'Female');

CREATE TABLE Actors (
    actor_id SERIAL PRIMARY KEY,
    actor_name VARCHAR(100) NOT NULL,
    gender gender,
    date_of_birth DATE
);

CREATE TABLE Movies (
    movie_id SERIAL PRIMARY KEY,
    movie_title VARCHAR(150) NOT NULL,
    description VARCHAR(1000),
    release_date DATE,
    rating INTEGER CHECK (rating BETWEEN 1 AND 10)
);

CREATE TABLE MovieActors (
    movie_id INT,
    actor_id INT,
    PRIMARY KEY (movie_id, actor_id),
    FOREIGN KEY (movie_id) REFERENCES Movies(movie_id) ON DELETE CASCADE,
    FOREIGN KEY (actor_id) REFERENCES Actors(actor_id) ON DELETE CASCADE
);

CREATE TABLE Users (
    username VARCHAR(256) UNIQUE NOT NULL,
    password VARCHAR(256) NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO Users VALUES ('admin', '73656372657421232f297a57a5a743894a0e4a801fc3', TRUE);
