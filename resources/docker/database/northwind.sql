--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

SET default_tablespace = '';

SET default_with_oids = false;

---
--- drop tables
---

DROP TABLE IF EXISTS url;

CREATE TABLE url (
    url_id bigint NOT NULL PRIMARY KEY,
    url character varying(100) NOT NULL,
    created_at smallint
);

--
-- PostgreSQL database dump complete
--

