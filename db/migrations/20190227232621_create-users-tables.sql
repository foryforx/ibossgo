-- migrate:up
CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp with time zone,
    sub     CHARACTER VARYING(50) NOT NULL ,     
    name     CHARACTER VARYING(250) NOT NULL  ,   
    given_name CHARACTER VARYING(250) NOT NULL  ,  
    family_name CHARACTER VARYING(250) NULL  , 
    profile    CHARACTER VARYING(550)  NULL  , 
    picture    CHARACTER VARYING(550)  NULL  , 
    email   CHARACTER VARYING(350) NOT NULL   ,   
    email_verified BOOLEAN NOT NULL DEFAULT false,
    gender   CHARACTER VARYING(550)  NULL   ,
    CONSTRAINT users_email_uniq UNIQUE (email)   
);

CREATE OR REPLACE FUNCTION trigger_set_update_at_timestamp()
  RETURNS trigger AS $$
BEGIN
	NEW.updated_at = NOW();
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- trigger for updated_at
CREATE TRIGGER trigger_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE PROCEDURE trigger_set_update_at_timestamp();



-- migrate:down
DROP TRIGGER trigger_users_updated_at ON users;
DROP TABLE users;
DROP FUNCTION trigger_set_update_at_timestamp();
