--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Drop databases (except postgres and template1)
--

DROP DATABASE pokemon_jackpot_db;




--
-- Drop roles
--

DROP ROLE postgres;


--
-- Roles
--

--
-- User Configurations
--








--
-- Databases
--

--
-- Database "template1" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Debian 16.4-1.pgdg120+1)
-- Dumped by pg_dump version 16.4 (Debian 16.4-1.pgdg120+1)

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

UPDATE pg_catalog.pg_database SET datistemplate = false WHERE datname = 'template1';
DROP DATABASE template1;
--
-- Name: template1; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE template1 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE template1 OWNER TO postgres;

\connect template1

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

--
-- Name: DATABASE template1; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE template1 IS 'default template for new databases';


--
-- Name: template1; Type: DATABASE PROPERTIES; Schema: -; Owner: postgres
--

ALTER DATABASE template1 IS_TEMPLATE = true;


\connect template1

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

--
-- Name: DATABASE template1; Type: ACL; Schema: -; Owner: postgres
--

REVOKE CONNECT,TEMPORARY ON DATABASE template1 FROM PUBLIC;
GRANT CONNECT ON DATABASE template1 TO PUBLIC;


--
-- PostgreSQL database dump complete
--

--
-- Database "pokemon_jackpot_db" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Debian 16.4-1.pgdg120+1)
-- Dumped by pg_dump version 16.4 (Debian 16.4-1.pgdg120+1)

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

--
-- Name: pokemon_jackpot_db; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE pokemon_jackpot_db WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE pokemon_jackpot_db OWNER TO postgres;

\connect pokemon_jackpot_db

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

--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Debian 16.4-1.pgdg120+1)
-- Dumped by pg_dump version 16.4 (Debian 16.4-1.pgdg120+1)

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

DROP DATABASE postgres;
--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE postgres OWNER TO postgres;

\connect postgres

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

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: pokemons; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pokemons (
    id bigint NOT NULL,
    image_url text,
    name text,
    type text
);


ALTER TABLE public.pokemons OWNER TO postgres;

--
-- Name: pokemons_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.pokemons_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pokemons_id_seq OWNER TO postgres;

--
-- Name: pokemons_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.pokemons_id_seq OWNED BY public.pokemons.id;


--
-- Name: user_pokemons; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_pokemons (
    user_id bigint,
    pokemon_id bigint
);


ALTER TABLE public.user_pokemons OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    username text,
    email text NOT NULL,
    password text
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: pokemons id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pokemons ALTER COLUMN id SET DEFAULT nextval('public.pokemons_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: pokemons; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.pokemons (id, image_url, name, type) FROM stdin;
1	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png	bulbasaur	grass
2	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/2.png	ivysaur	grass
3	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/3.png	venusaur	grass
4	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/4.png	charmander	fire
5	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/5.png	charmeleon	fire
6	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/6.png	charizard	fire
7	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/7.png	squirtle	water
8	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/8.png	wartortle	water
9	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/9.png	blastoise	water
10	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/10.png	caterpie	bug
11	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/11.png	metapod	bug
12	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/12.png	butterfree	bug
13	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/13.png	weedle	bug
14	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/14.png	kakuna	bug
15	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/15.png	beedrill	bug
16	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/16.png	pidgey	normal
17	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/17.png	pidgeotto	normal
18	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/18.png	pidgeot	normal
19	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/19.png	rattata	normal
20	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/20.png	raticate	normal
21	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/21.png	spearow	normal
22	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/22.png	fearow	normal
23	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/23.png	ekans	poison
24	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/24.png	arbok	poison
25	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/25.png	pikachu	electric
26	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/26.png	raichu	electric
27	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/27.png	sandshrew	ground
28	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/28.png	sandslash	ground
29	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/29.png	nidoran-f	poison
30	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/30.png	nidorina	poison
31	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/31.png	nidoqueen	poison
32	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/32.png	nidoran-m	poison
33	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/33.png	nidorino	poison
34	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/34.png	nidoking	poison
35	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/35.png	clefairy	fairy
36	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/36.png	clefable	fairy
37	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/37.png	vulpix	fire
38	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/38.png	ninetales	fire
39	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/39.png	jigglypuff	normal
40	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/40.png	wigglytuff	normal
41	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/41.png	zubat	poison
42	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/42.png	golbat	poison
43	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/43.png	oddish	grass
44	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/44.png	gloom	grass
45	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/45.png	vileplume	grass
46	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/46.png	paras	bug
47	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/47.png	parasect	bug
48	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/48.png	venonat	bug
49	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/49.png	venomoth	bug
50	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/50.png	diglett	ground
51	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/51.png	dugtrio	ground
52	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/52.png	meowth	normal
53	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/53.png	persian	normal
54	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/54.png	psyduck	water
55	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/55.png	golduck	water
56	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/56.png	mankey	fighting
57	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/57.png	primeape	fighting
58	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/58.png	growlithe	fire
59	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/59.png	arcanine	fire
60	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/60.png	poliwag	water
61	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/61.png	poliwhirl	water
62	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/62.png	poliwrath	water
63	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/63.png	abra	psychic
64	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/64.png	kadabra	psychic
65	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/65.png	alakazam	psychic
66	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/66.png	machop	fighting
67	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/67.png	machoke	fighting
68	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/68.png	machamp	fighting
69	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/69.png	bellsprout	grass
70	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/70.png	weepinbell	grass
71	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/71.png	victreebel	grass
72	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/72.png	tentacool	water
73	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/73.png	tentacruel	water
74	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/74.png	geodude	rock
75	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/75.png	graveler	rock
76	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/76.png	golem	rock
77	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/77.png	ponyta	fire
78	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/78.png	rapidash	fire
79	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/79.png	slowpoke	water
80	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/80.png	slowbro	water
81	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/81.png	magnemite	electric
82	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/82.png	magneton	electric
83	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/83.png	farfetchd	normal
84	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/84.png	doduo	normal
85	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/85.png	dodrio	normal
86	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/86.png	seel	water
87	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/87.png	dewgong	water
88	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/88.png	grimer	poison
89	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/89.png	muk	poison
90	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/90.png	shellder	water
91	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/91.png	cloyster	water
92	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/92.png	gastly	ghost
93	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/93.png	haunter	ghost
94	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/94.png	gengar	ghost
95	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/95.png	onix	rock
96	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/96.png	drowzee	psychic
97	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/97.png	hypno	psychic
98	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/98.png	krabby	water
99	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/99.png	kingler	water
100	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/100.png	voltorb	electric
101	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/101.png	electrode	electric
102	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/102.png	exeggcute	grass
103	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/103.png	exeggutor	grass
104	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/104.png	cubone	ground
105	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/105.png	marowak	ground
106	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/106.png	hitmonlee	fighting
107	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/107.png	hitmonchan	fighting
108	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/108.png	lickitung	normal
109	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/109.png	koffing	poison
110	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/110.png	weezing	poison
111	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/111.png	rhyhorn	ground
112	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/112.png	rhydon	ground
113	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/113.png	chansey	normal
114	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/114.png	tangela	grass
115	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/115.png	kangaskhan	normal
116	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/116.png	horsea	water
117	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/117.png	seadra	water
118	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/118.png	goldeen	water
119	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/119.png	seaking	water
120	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/120.png	staryu	water
121	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/121.png	starmie	water
122	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/122.png	mr-mime	psychic
123	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/123.png	scyther	bug
124	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/124.png	jynx	ice
125	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/125.png	electabuzz	electric
126	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/126.png	magmar	fire
127	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/127.png	pinsir	bug
128	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/128.png	tauros	normal
129	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/129.png	magikarp	water
130	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/130.png	gyarados	water
131	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/131.png	lapras	water
132	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/132.png	ditto	normal
133	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/133.png	eevee	normal
134	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/134.png	vaporeon	water
135	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/135.png	jolteon	electric
136	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/136.png	flareon	fire
137	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/137.png	porygon	normal
138	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/138.png	omanyte	rock
139	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/139.png	omastar	rock
140	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/140.png	kabuto	rock
141	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/141.png	kabutops	rock
142	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/142.png	aerodactyl	rock
143	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/143.png	snorlax	normal
144	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/144.png	articuno	ice
145	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/145.png	zapdos	electric
146	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/146.png	moltres	fire
147	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/147.png	dratini	dragon
148	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/148.png	dragonair	dragon
149	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/149.png	dragonite	dragon
150	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/150.png	mewtwo	psychic
151	https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/151.png	mew	psychic
\.


--
-- Data for Name: user_pokemons; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_pokemons (user_id, pokemon_id) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, email, password) FROM stdin;
\.


--
-- Name: pokemons_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.pokemons_id_seq', 151, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: pokemons pokemons_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pokemons
    ADD CONSTRAINT pokemons_pkey PRIMARY KEY (id);


--
-- Name: users uni_users_email; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_email UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database cluster dump complete
--

