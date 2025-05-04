--
-- PostgreSQL database dump
--

-- Dumped from database version 13.20 (Postgres.app)
-- Dumped by pg_dump version 14.17 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: reservations; Type: TABLE; Schema: public; Owner: jim
--

CREATE TABLE public.reservations (
    id uuid NOT NULL,
    first_name character varying(255) DEFAULT ''::character varying NOT NULL,
    last_name character varying(255) DEFAULT ''::character varying NOT NULL,
    email character varying(255) NOT NULL,
    phone character varying(255) DEFAULT ''::character varying NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    room_id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.reservations OWNER TO jim;

--
-- Name: restrictions; Type: TABLE; Schema: public; Owner: jim
--

CREATE TABLE public.restrictions (
    id uuid NOT NULL,
    restriction_name character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.restrictions OWNER TO jim;

--
-- Name: room_restrictions; Type: TABLE; Schema: public; Owner: jim
--

CREATE TABLE public.room_restrictions (
    id uuid NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    room_id integer NOT NULL,
    reservation_id integer NOT NULL,
    restriction_id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.room_restrictions OWNER TO jim;

--
-- Name: rooms; Type: TABLE; Schema: public; Owner: jim
--

CREATE TABLE public.rooms (
    id uuid NOT NULL,
    room_name character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.rooms OWNER TO jim;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: jim
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO jim;

--
-- Name: users; Type: TABLE; Schema: public; Owner: jim
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    first_name character varying(255) DEFAULT ''::character varying NOT NULL,
    last_name character varying(255) DEFAULT ''::character varying NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(60) NOT NULL,
    access_level integer DEFAULT 1 NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO jim;

--
-- Name: reservations reservations_pkey; Type: CONSTRAINT; Schema: public; Owner: jim
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservations_pkey PRIMARY KEY (id);


--
-- Name: restrictions restrictions_pkey; Type: CONSTRAINT; Schema: public; Owner: jim
--

ALTER TABLE ONLY public.restrictions
    ADD CONSTRAINT restrictions_pkey PRIMARY KEY (id);


--
-- Name: room_restrictions room_restrictions_pkey; Type: CONSTRAINT; Schema: public; Owner: jim
--

ALTER TABLE ONLY public.room_restrictions
    ADD CONSTRAINT room_restrictions_pkey PRIMARY KEY (id);


--
-- Name: rooms rooms_pkey; Type: CONSTRAINT; Schema: public; Owner: jim
--

ALTER TABLE ONLY public.rooms
    ADD CONSTRAINT rooms_pkey PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: jim
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: jim
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: jim
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

