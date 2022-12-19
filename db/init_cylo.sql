--
-- PostgreSQL database dump
--

-- Dumped from database version 11.5
-- Dumped by pg_dump version 14.5 (Ubuntu 14.5-0ubuntu0.22.04.1)

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

--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgreuser
--

CREATE TABLE public.categories (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    parent_id bigint
);


ALTER TABLE public.categories OWNER TO postgreuser;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgreuser
--

CREATE SEQUENCE public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.categories_id_seq OWNER TO postgreuser;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgreuser
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: postgreuser
--

CREATE TABLE public.products (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    price numeric,
    images text[],
    description text,
    category_id bigint
);


ALTER TABLE public.products OWNER TO postgreuser;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: postgreuser
--

CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.products_id_seq OWNER TO postgreuser;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgreuser
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: postgreuser
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: postgreuser
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgreuser
--

COPY public.categories (id, created_at, updated_at, deleted_at, name, parent_id) FROM stdin;
1	2022-12-18 12:19:38.367562+07	2022-12-18 12:19:38.367562+07	\N	ROOT	\N
2	2022-12-18 12:19:38.367562+07	2022-12-18 12:19:38.367562+07	\N	Fashion	1
3	2022-12-18 12:19:38.367562+07	2022-12-18 12:19:38.367562+07	\N	Shirt	2
4	2022-12-18 12:19:38.367562+07	2022-12-18 12:19:38.367562+07	\N	Jacket	2
5	2022-12-18 12:19:38.367562+07	2022-12-18 12:19:38.367562+07	\N	Dress	2
6	2022-12-18 12:19:38.367562+07	2022-12-18 12:19:38.367562+07	\N	Unisex	2
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgreuser
--

COPY public.products (id, created_at, updated_at, deleted_at, name, price, images, description, category_id) FROM stdin;
1	2022-12-18 12:37:45.423666+07	2022-12-18 12:37:45.423666+07	\N	Love and Thunder Blazer	100000	{https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F0%2F0.jpg?alt=media&token=7b40f3cc-37a2-414a-8550-cc8b2389dda3,https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F0%2F1.jpg?alt=media&token=8faa9613-e27d-4193-91ec-615d2d53dfe5}	The pink blazer from Love and Thunder.	4
2	2022-12-18 12:37:45.423666+07	2022-12-18 12:37:45.423666+07	\N	Black Window Dress	300000	{https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F1%2F0.jpg?alt=media&token=cfb5ea77-7215-45a1-946c-78f281fdd4b6,https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F1%2F1.jpg?alt=media&token=7ed979aa-2992-4d97-a862-c98007312ad6}	Black Window wears it to protect the Earth from Thanos.	5
3	2022-12-18 12:37:45.423666+07	2022-12-18 12:37:45.423666+07	\N	Sport Suit IronHeart	250000	{https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F10%2F0.jpg?alt=media&token=9baece8e-5bb2-458f-b454-9308c6afebf2,https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F10%2F1.jpg?alt=media&token=fb028a35-6ddf-4c3c-9634-2db90d202756}	IronHeart x Shuri in Wakanda forever.	5
4	2022-12-18 12:37:45.423666+07	2022-12-18 12:37:45.423666+07	\N	Rubber suit	500000	{https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F11%2F0.jpg?alt=media&token=909e31c1-38bc-4de1-ad08-4f5b52f9bad1,https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F11%2F1.jpg?alt=media&token=9f2a9325-4c9e-4d04-85f2-a455add0b901}	This dress trades on people's fears to make you special.	6
5	2022-12-18 12:37:45.423666+07	2022-12-18 12:37:45.423666+07	\N	The jacket of Tom Holland	900000	{https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F12%2F0.jpg?alt=media&token=a4379223-24db-4184-89de-dcf4cb089a15,https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F12%2F1.jpg?alt=media&token=9d705a38-8bcf-4daa-9043-5e073e4e0d5a}	Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.	3
6	2022-12-18 12:37:45.423666+07	2022-12-18 12:37:45.423666+07	\N	Calvin Klein Sheath Chiffon Bell Sleeves	1000000	{https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F9%2F0.jpg?alt=media&token=c1a463a8-60f9-4cea-a5de-0041d8a04eff,https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F9%2F1.jpg?alt=media&token=8e7b175c-04fe-4f67-b0db-5fcd834bef42}	Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.	6
7	2022-12-18 12:37:45.423666+07	2022-12-18 12:37:45.423666+07	\N	Dress the Population Women	230000	{https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F8%2F0.jpg?alt=media&token=cd734ffb-6f6a-404a-bc7c-92082ae69d7b,https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F8%2F1.jpg?alt=media&token=b7731bb4-f161-4052-94b0-168749014fbe}	Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.	3
\.


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgreuser
--

SELECT pg_catalog.setval('public.categories_id_seq', 8, true);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgreuser
--

SELECT pg_catalog.setval('public.products_id_seq', 5, true);


--
-- Name: categories categories_name_key; Type: CONSTRAINT; Schema: public; Owner: postgreuser
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_name_key UNIQUE (name);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgreuser
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgreuser
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: idx_categories_deleted_at; Type: INDEX; Schema: public; Owner: postgreuser
--

CREATE INDEX idx_categories_deleted_at ON public.categories USING btree (deleted_at);


--
-- Name: idx_products_deleted_at; Type: INDEX; Schema: public; Owner: postgreuser
--

CREATE INDEX idx_products_deleted_at ON public.products USING btree (deleted_at);


--
-- Name: categories fk_categories_parent; Type: FK CONSTRAINT; Schema: public; Owner: postgreuser
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT fk_categories_parent FOREIGN KEY (parent_id) REFERENCES public.categories(id);


--
-- Name: products fk_products_category; Type: FK CONSTRAINT; Schema: public; Owner: postgreuser
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_category FOREIGN KEY (category_id) REFERENCES public.categories(id);


--
-- PostgreSQL database dump complete
--

