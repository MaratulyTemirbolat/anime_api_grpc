-- TRIGGERS FOR USERS TABLE

-- DELETE TRIGGER
CREATE OR REPLACE FUNCTION delete_user_trigger()
    RETURNS TRIGGER
AS
$$
DECLARE
BEGIN
    UPDATE users
    SET deleted_at = current_timestamp
    WHERE id = OLD.id AND deleted_at IS NULL;
    RETURN NULL;
END;
$$ language plpgsql;

CREATE TRIGGER delete_user_trigger
    BEFORE DELETE
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE delete_user_trigger();

-- UPDATE TRIGGER
CREATE OR REPLACE FUNCTION update_function_trigger()
    RETURNS TRIGGER
AS
$$
DECLARE
BEGIN
    NEW.updated_at = current_timestamp;
    RETURN NEW;
END;
$$ language plpgsql;

CREATE TRIGGER update_user_trigger
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE update_function_trigger();

SELECT * FROM users;
SET datestyle = dmy;
INSERT INTO users(username, first_name, last_name, email, password, birthday) VALUES ('Temir','Temirbolat', 'Maratuly','t_maratuly@kbtu.kz','12345','31.01.2001');
INSERT INTO users(username, first_name, last_name, email, password, birthday) VALUES ('A','Temirbolat', 'Maratuly','t_maratuly@gmail.com','12345','31.01.2001');
INSERT INTO users(username, first_name, last_name, email, password, birthday) VALUES ('B','Temirbolat', 'Maratuly','t_maratuly@mail.ru','12345','31.01.2001');

-- TRIGGERS FOR ANIME TABLE

-- DELETE TRIGGER
CREATE OR REPLACE FUNCTION delete_anime_function_trigger()
    RETURNS TRIGGER
AS
$$
DECLARE
BEGIN
    UPDATE animes
    SET deleted_at = current_timestamp
    WHERE id = OLD.id AND deleted_at IS NULL;
    RETURN NULL;
END;
$$ language plpgsql;

CREATE TRIGGER delete_anime_trigger
    BEFORE DELETE
    ON animes
    FOR EACH ROW
EXECUTE PROCEDURE delete_anime_function_trigger();

-- UPDATE TRIGGER

CREATE TRIGGER update_anime_trigger
    BEFORE UPDATE
    ON animes
    FOR EACH ROW
EXECUTE PROCEDURE update_function_trigger();

-- TRIGGERS ON TYPES TABLE

-- DELETE TRIGGER
CREATE OR REPLACE FUNCTION delete_type_function_trigger()
    RETURNS TRIGGER
AS
$$
DECLARE
BEGIN
    UPDATE types
    SET deleted_at = current_timestamp
    WHERE id = OLD.id AND deleted_at IS NULL;
    RETURN NULL;
END;
$$ language plpgsql;

CREATE TRIGGER delete_anime_trigger
    BEFORE DELETE
    ON types
    FOR EACH ROW
EXECUTE PROCEDURE delete_type_function_trigger();

-- UPDATE TRIGGER

CREATE TRIGGER update_type_trigger
    BEFORE UPDATE
    ON types
    FOR EACH ROW
EXECUTE PROCEDURE update_function_trigger();


-- TRIGGERS FOR ANIME_GROUPS TABLE
-- DELETE TRIGGER
CREATE OR REPLACE FUNCTION delete_anime_group_function_trigger()
    RETURNS TRIGGER
AS
$$
DECLARE
BEGIN
    UPDATE anime_groups
    SET deleted_at = current_timestamp
    WHERE id = OLD.id AND deleted_at IS NULL;
    RETURN NULL;
END;
$$ language plpgsql;

CREATE TRIGGER delete_anime_group_trigger
    BEFORE DELETE
    ON anime_groups
    FOR EACH ROW
EXECUTE PROCEDURE delete_anime_group_function_trigger();

-- UPDATE TRIGGER

CREATE TRIGGER update_anime_group_trigger
    BEFORE UPDATE
    ON anime_groups
    FOR EACH ROW
EXECUTE PROCEDURE update_function_trigger();

-- TRIGGERS FOR STUDIOS TABLE
-- DELETE TRIGGER
CREATE OR REPLACE FUNCTION delete_studios_function_trigger()
    RETURNS TRIGGER
AS
$$
DECLARE
BEGIN
    UPDATE studios
    SET deleted_at = current_timestamp
    WHERE id = OLD.id AND deleted_at IS NULL;
    RETURN NULL;
END;
$$ language plpgsql;

CREATE TRIGGER delete_studios_trigger
    BEFORE DELETE
    ON studios
    FOR EACH ROW
EXECUTE PROCEDURE delete_studios_function_trigger();

-- UPDATE TRIGGER

CREATE TRIGGER update_studios_trigger
    BEFORE UPDATE
    ON studios
    FOR EACH ROW
EXECUTE PROCEDURE update_function_trigger();

-- TRIGGERS FOR ACTIONS TABLE
-- DELETE TRIGGER
CREATE OR REPLACE FUNCTION delete_actions_function_trigger()
    RETURNS TRIGGER
AS
$$
DECLARE
BEGIN
    UPDATE actions
    SET deleted_at = current_timestamp
    WHERE id = OLD.id AND deleted_at IS NULL;
    RETURN NULL;
END;
$$ language plpgsql;

CREATE TRIGGER delete_actions_trigger
    BEFORE DELETE
    ON actions
    FOR EACH ROW
EXECUTE PROCEDURE delete_actions_function_trigger();

-- UPDATE TRIGGER

CREATE TRIGGER update_actions_trigger
    BEFORE UPDATE
    ON actions
    FOR EACH ROW
EXECUTE PROCEDURE update_function_trigger();


-- TRIGGERS FOR TAGS TABLE
-- DELETE TRIGGER
CREATE OR REPLACE FUNCTION delete_tags_function_trigger()
    RETURNS TRIGGER
AS
$$
DECLARE
BEGIN
    UPDATE tags
    SET deleted_at = current_timestamp
    WHERE id = OLD.id AND deleted_at IS NULL;
    RETURN NULL;
END;
$$ language plpgsql;

CREATE TRIGGER delete_tags_trigger
    BEFORE DELETE
    ON tags
    FOR EACH ROW
EXECUTE PROCEDURE delete_tags_function_trigger();

-- UPDATE TRIGGER

CREATE TRIGGER update_tags_trigger
    BEFORE UPDATE
    ON tags
    FOR EACH ROW
EXECUTE PROCEDURE update_function_trigger();


-- TRIGGERS FOR GENRE TABLE
-- DELETE TRIGGER
CREATE OR REPLACE FUNCTION delete_genre_function_trigger()
    RETURNS TRIGGER
AS
$$
DECLARE
BEGIN
    UPDATE genres
    SET deleted_at = current_timestamp
    WHERE id = OLD.id AND deleted_at IS NULL;
    RETURN NULL;
END;
$$ language plpgsql;

CREATE TRIGGER delete_genres_trigger
    BEFORE DELETE
    ON genres
    FOR EACH ROW
EXECUTE PROCEDURE delete_genre_function_trigger();

-- UPDATE TRIGGER

CREATE TRIGGER update_genres_trigger
    BEFORE UPDATE
    ON genres
    FOR EACH ROW
EXECUTE PROCEDURE update_function_trigger();


-- TRIGGERS FOR PHONES TABLE

-- DELETE TRIGGER
CREATE OR REPLACE FUNCTION delete_phone_function_trigger()
    RETURNS TRIGGER
AS
$$
DECLARE
BEGIN
    UPDATE phones
    SET deleted_at = current_timestamp
    WHERE id = OLD.id AND deleted_at IS NULL;
    RETURN NULL;
END;
$$ language plpgsql;

CREATE TRIGGER delete_phones_trigger
    BEFORE DELETE
    ON phones
    FOR EACH ROW
EXECUTE PROCEDURE delete_phone_function_trigger();

-- UPDATE TRIGGER

CREATE TRIGGER update_phones_trigger
    BEFORE UPDATE
    ON phones
    FOR EACH ROW
EXECUTE PROCEDURE update_function_trigger();

-- TRIGGERS FOR COMMENTS TABLE

-- DELETE TRIGGER
CREATE OR REPLACE FUNCTION delete_comments_function_trigger()
    RETURNS TRIGGER
AS
$$
DECLARE
BEGIN
    UPDATE comments
    SET deleted_at = current_timestamp
    WHERE id = OLD.id AND deleted_at IS NULL;
    RETURN NULL;
END;
$$ language plpgsql;

CREATE TRIGGER delete_comments_trigger
    BEFORE DELETE
    ON comments
    FOR EACH ROW
EXECUTE PROCEDURE delete_comments_function_trigger();

-- UPDATE TRIGGER

CREATE TRIGGER update_comments_trigger
    BEFORE UPDATE
    ON comments
    FOR EACH ROW
EXECUTE PROCEDURE update_function_trigger();

-- TRIGGERS FOR SEASONS TABLE

-- DELETE TRIGGER
CREATE OR REPLACE FUNCTION delete_seasons_function_trigger()
    RETURNS TRIGGER
AS
$$
DECLARE
BEGIN
    UPDATE seasons
    SET deleted_at = current_timestamp
    WHERE id = OLD.id AND deleted_at IS NULL;
    RETURN NULL;
END;
$$ language plpgsql;

CREATE TRIGGER delete_seasons_trigger
    BEFORE DELETE
    ON seasons
    FOR EACH ROW
EXECUTE PROCEDURE delete_seasons_function_trigger();

-- UPDATE TRIGGER

CREATE TRIGGER update_seasons_trigger
    BEFORE UPDATE
    ON seasons
    FOR EACH ROW
EXECUTE PROCEDURE update_function_trigger();