-- Started on 2026-01-11 20:49:14

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 239 (class 1255 OID 18744)
-- Name: update_updated_at_column(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_updated_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_updated_at_column() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 222 (class 1259 OID 18550)
-- Name: auth_tokens; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.auth_tokens (
    id integer NOT NULL,
    user_id integer NOT NULL,
    token character varying(255) NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.auth_tokens OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 18549)
-- Name: auth_tokens_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.auth_tokens_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.auth_tokens_id_seq OWNER TO postgres;

--
-- TOC entry 5162 (class 0 OID 0)
-- Dependencies: 221
-- Name: auth_tokens_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.auth_tokens_id_seq OWNED BY public.auth_tokens.id;


--
-- TOC entry 236 (class 1259 OID 18673)
-- Name: bookings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bookings (
    id integer NOT NULL,
    user_id integer NOT NULL,
    showtime_id integer NOT NULL,
    seat_id integer NOT NULL,
    booking_code character varying(20) NOT NULL,
    status character varying(20) DEFAULT 'pending'::character varying,
    total_price numeric(10,2) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.bookings OWNER TO postgres;

--
-- TOC entry 235 (class 1259 OID 18672)
-- Name: bookings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bookings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.bookings_id_seq OWNER TO postgres;

--
-- TOC entry 5163 (class 0 OID 0)
-- Dependencies: 235
-- Name: bookings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bookings_id_seq OWNED BY public.bookings.id;


--
-- TOC entry 226 (class 1259 OID 18587)
-- Name: cinemas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cinemas (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    location character varying(255) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.cinemas OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 18586)
-- Name: cinemas_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cinemas_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cinemas_id_seq OWNER TO postgres;

--
-- TOC entry 5164 (class 0 OID 0)
-- Dependencies: 225
-- Name: cinemas_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cinemas_id_seq OWNED BY public.cinemas.id;


--
-- TOC entry 228 (class 1259 OID 18600)
-- Name: movies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.movies (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    duration integer NOT NULL,
    genre character varying(100),
    poster_url character varying(500),
    rating character varying(10),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.movies OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 18599)
-- Name: movies_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.movies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.movies_id_seq OWNER TO postgres;

--
-- TOC entry 5165 (class 0 OID 0)
-- Dependencies: 227
-- Name: movies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.movies_id_seq OWNED BY public.movies.id;


--
-- TOC entry 224 (class 1259 OID 18569)
-- Name: otp_codes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.otp_codes (
    id integer NOT NULL,
    user_id integer NOT NULL,
    code character varying(6) NOT NULL,
    is_used boolean DEFAULT false,
    expires_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.otp_codes OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 18568)
-- Name: otp_codes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.otp_codes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.otp_codes_id_seq OWNER TO postgres;

--
-- TOC entry 5166 (class 0 OID 0)
-- Dependencies: 223
-- Name: otp_codes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.otp_codes_id_seq OWNED BY public.otp_codes.id;


--
-- TOC entry 234 (class 1259 OID 18659)
-- Name: payment_methods; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payment_methods (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    code character varying(20) NOT NULL,
    is_active boolean DEFAULT true,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.payment_methods OWNER TO postgres;

--
-- TOC entry 233 (class 1259 OID 18658)
-- Name: payment_methods_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payment_methods_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payment_methods_id_seq OWNER TO postgres;

--
-- TOC entry 5167 (class 0 OID 0)
-- Dependencies: 233
-- Name: payment_methods_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payment_methods_id_seq OWNED BY public.payment_methods.id;


--
-- TOC entry 238 (class 1259 OID 18708)
-- Name: payments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payments (
    id integer NOT NULL,
    booking_id integer NOT NULL,
    payment_method_id integer NOT NULL,
    amount numeric(10,2) NOT NULL,
    status character varying(20) DEFAULT 'pending'::character varying,
    payment_details jsonb,
    paid_at timestamp without time zone,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.payments OWNER TO postgres;

--
-- TOC entry 237 (class 1259 OID 18707)
-- Name: payments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payments_id_seq OWNER TO postgres;

--
-- TOC entry 5168 (class 0 OID 0)
-- Dependencies: 237
-- Name: payments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payments_id_seq OWNED BY public.payments.id;


--
-- TOC entry 232 (class 1259 OID 18639)
-- Name: seats; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.seats (
    id integer NOT NULL,
    cinema_id integer NOT NULL,
    seat_row character varying(2) NOT NULL,
    seat_number integer NOT NULL,
    seat_type character varying(20) DEFAULT 'regular'::character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.seats OWNER TO postgres;

--
-- TOC entry 231 (class 1259 OID 18638)
-- Name: seats_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.seats_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.seats_id_seq OWNER TO postgres;

--
-- TOC entry 5169 (class 0 OID 0)
-- Dependencies: 231
-- Name: seats_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.seats_id_seq OWNED BY public.seats.id;


--
-- TOC entry 230 (class 1259 OID 18613)
-- Name: showtimes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.showtimes (
    id integer NOT NULL,
    cinema_id integer NOT NULL,
    movie_id integer NOT NULL,
    show_date date NOT NULL,
    show_time time without time zone NOT NULL,
    price numeric(10,2) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.showtimes OWNER TO postgres;

--
-- TOC entry 229 (class 1259 OID 18612)
-- Name: showtimes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.showtimes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.showtimes_id_seq OWNER TO postgres;

--
-- TOC entry 5170 (class 0 OID 0)
-- Dependencies: 229
-- Name: showtimes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.showtimes_id_seq OWNED BY public.showtimes.id;


--
-- TOC entry 220 (class 1259 OID 18532)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    email character varying(100) NOT NULL,
    password_hash character varying(255) NOT NULL,
    is_verified boolean DEFAULT false,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 18531)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 5171 (class 0 OID 0)
-- Dependencies: 219
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 4906 (class 2604 OID 18553)
-- Name: auth_tokens id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_tokens ALTER COLUMN id SET DEFAULT nextval('public.auth_tokens_id_seq'::regclass);


--
-- TOC entry 4923 (class 2604 OID 18676)
-- Name: bookings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings ALTER COLUMN id SET DEFAULT nextval('public.bookings_id_seq'::regclass);


--
-- TOC entry 4911 (class 2604 OID 18590)
-- Name: cinemas id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas ALTER COLUMN id SET DEFAULT nextval('public.cinemas_id_seq'::regclass);


--
-- TOC entry 4913 (class 2604 OID 18603)
-- Name: movies id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);


--
-- TOC entry 4908 (class 2604 OID 18572)
-- Name: otp_codes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.otp_codes ALTER COLUMN id SET DEFAULT nextval('public.otp_codes_id_seq'::regclass);


--
-- TOC entry 4920 (class 2604 OID 18662)
-- Name: payment_methods id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods ALTER COLUMN id SET DEFAULT nextval('public.payment_methods_id_seq'::regclass);


--
-- TOC entry 4927 (class 2604 OID 18711)
-- Name: payments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments ALTER COLUMN id SET DEFAULT nextval('public.payments_id_seq'::regclass);


--
-- TOC entry 4917 (class 2604 OID 18642)
-- Name: seats id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats ALTER COLUMN id SET DEFAULT nextval('public.seats_id_seq'::regclass);


--
-- TOC entry 4915 (class 2604 OID 18616)
-- Name: showtimes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes ALTER COLUMN id SET DEFAULT nextval('public.showtimes_id_seq'::regclass);


--
-- TOC entry 4902 (class 2604 OID 18535)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 5140 (class 0 OID 18550)
-- Dependencies: 222
-- Data for Name: auth_tokens; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.auth_tokens VALUES (1, 1, 'edd3ed98c8f9066d87078a9ddf721bd94245d351860627321c9a642e77f38fd6', '2026-01-12 14:52:59.079102', '2026-01-11 14:52:59.079102');
INSERT INTO public.auth_tokens VALUES (4, 1, '5c5e4f33733eb649236255a99aa547185764313f86abaf64191f37d05f18a4cc', '2026-01-12 15:12:08.25453', '2026-01-11 15:12:08.25453');
INSERT INTO public.auth_tokens VALUES (5, 1, 'e65e23f462761bf667e304f5ae4ff8ed99297743a5908cb18d75bc0b546cc5a6', '2026-01-12 15:53:49.717154', '2026-01-11 15:53:49.717154');
INSERT INTO public.auth_tokens VALUES (6, 1, 'dded2473001d6291caa4a5e78d87001afc1afc872b59ab03e5fbf5d714cce766', '2026-01-12 18:57:53.385697', '2026-01-11 18:57:53.385697');
INSERT INTO public.auth_tokens VALUES (8, 3, '7c30f8394b273b3459e3040da0b8eb1ed648acfc99ece4f4562600385c4e46f6', '2026-01-12 19:31:37.194138', '2026-01-11 19:31:37.194138');
INSERT INTO public.auth_tokens VALUES (9, 4, '9bea8a85805f11ae04ec79911546700fc724f6560d651a5450406dfdbfd77c95', '2026-01-12 19:37:20.553049', '2026-01-11 19:37:20.553049');


--
-- TOC entry 5154 (class 0 OID 18673)
-- Dependencies: 236
-- Data for Name: bookings; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.bookings VALUES (1, 1, 1, 5, 'BK4d3f4a5c1df4', 'confirmed', 50000.00, '2026-01-11 15:55:16.590239', '2026-01-11 16:06:55.496969');
INSERT INTO public.bookings VALUES (2, 1, 1, 10, 'BK3d664b62cfdd', 'confirmed', 50000.00, '2026-01-11 16:13:23.493926', '2026-01-11 16:19:39.721201');
INSERT INTO public.bookings VALUES (3, 1, 1, 9, 'BK6f981abc5793', 'confirmed', 50000.00, '2026-01-11 16:20:35.513537', '2026-01-11 16:20:46.330849');
INSERT INTO public.bookings VALUES (4, 1, 1, 8, 'BK34246cc2e5dd', 'confirmed', 50000.00, '2026-01-11 16:28:27.041895', '2026-01-11 16:28:59.51716');
INSERT INTO public.bookings VALUES (5, 1, 1, 7, 'BKfe338665bd04', 'confirmed', 50000.00, '2026-01-11 16:41:27.421688', '2026-01-11 16:42:21.372575');
INSERT INTO public.bookings VALUES (6, 1, 1, 6, 'BK13f5d02dd3c3', 'confirmed', 50000.00, '2026-01-11 18:58:11.596526', '2026-01-11 18:58:55.390887');


--
-- TOC entry 5144 (class 0 OID 18587)
-- Dependencies: 226
-- Data for Name: cinemas; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.cinemas VALUES (1, 'CGV Grand Indonesia', 'Jakarta Pusat', 'Cinema dengan fasilitas premium di pusat kota Jakarta', '2026-01-11 14:24:20.203825');
INSERT INTO public.cinemas VALUES (2, 'XXI Plaza Senayan', 'Jakarta Selatan', 'Bioskop modern dengan teknologi terkini', '2026-01-11 14:24:20.203825');
INSERT INTO public.cinemas VALUES (3, 'Cinepolis Lippo Mall Puri', 'Jakarta Barat', 'Cinema dengan konsep luxury dan VIP lounge', '2026-01-11 14:24:20.203825');
INSERT INTO public.cinemas VALUES (4, 'CGV Blitz Megaplex', 'Jakarta Timur', 'Bioskop dengan berbagai pilihan film dan seat type', '2026-01-11 14:24:20.203825');
INSERT INTO public.cinemas VALUES (5, 'XXI Pondok Indah Mall', 'Jakarta Selatan', 'Cinema dengan sound system Dolby Atmos', '2026-01-11 14:24:20.203825');


--
-- TOC entry 5146 (class 0 OID 18600)
-- Dependencies: 228
-- Data for Name: movies; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.movies VALUES (1, 'The Dark Knight', 'Batman melawan Joker dalam pertarungan epik untuk menyelamatkan Gotham City', 152, 'Action, Crime, Drama', 'https://example.com/dark-knight.jpg', '13+', '2026-01-11 14:24:20.203825');
INSERT INTO public.movies VALUES (2, 'Inception', 'Seorang pencuri yang mencuri rahasia korporat melalui dream-sharing technology', 148, 'Action, Sci-Fi, Thriller', 'https://example.com/inception.jpg', '13+', '2026-01-11 14:24:20.203825');
INSERT INTO public.movies VALUES (3, 'Interstellar', 'Sekelompok penjelajah menggunakan lubang cacing untuk melintasi dimensi ruang', 169, 'Adventure, Drama, Sci-Fi', 'https://example.com/interstellar.jpg', '13+', '2026-01-11 14:24:20.203825');
INSERT INTO public.movies VALUES (4, 'Parasite', 'Keluarga miskin yang menyusup ke kehidupan keluarga kaya', 132, 'Comedy, Drama, Thriller', 'https://example.com/parasite.jpg', '17+', '2026-01-11 14:24:20.203825');
INSERT INTO public.movies VALUES (5, 'Avengers: Endgame', 'Para Avengers berkumpul untuk mengalahkan Thanos sekali dan untuk selamanya', 181, 'Action, Adventure, Sci-Fi', 'https://example.com/endgame.jpg', '13+', '2026-01-11 14:24:20.203825');


--
-- TOC entry 5142 (class 0 OID 18569)
-- Dependencies: 224
-- Data for Name: otp_codes; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.otp_codes VALUES (2, 3, '868967', false, '2026-01-11 19:41:37.192631', '2026-01-11 19:31:37.192631');
INSERT INTO public.otp_codes VALUES (3, 4, '118454', true, '2026-01-11 19:47:20.551536', '2026-01-11 19:37:20.551536');


--
-- TOC entry 5152 (class 0 OID 18659)
-- Dependencies: 234
-- Data for Name: payment_methods; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.payment_methods VALUES (1, 'Credit Card', 'CREDIT_CARD', true, '2026-01-11 14:24:20.203825');
INSERT INTO public.payment_methods VALUES (2, 'Debit Card', 'DEBIT_CARD', true, '2026-01-11 14:24:20.203825');
INSERT INTO public.payment_methods VALUES (3, 'GoPay', 'GOPAY', true, '2026-01-11 14:24:20.203825');
INSERT INTO public.payment_methods VALUES (4, 'OVO', 'OVO', true, '2026-01-11 14:24:20.203825');
INSERT INTO public.payment_methods VALUES (5, 'DANA', 'DANA', true, '2026-01-11 14:24:20.203825');
INSERT INTO public.payment_methods VALUES (6, 'ShopeePay', 'SHOPEEPAY', true, '2026-01-11 14:24:20.203825');
INSERT INTO public.payment_methods VALUES (7, 'Bank Transfer', 'BANK_TRANSFER', true, '2026-01-11 14:24:20.203825');


--
-- TOC entry 5156 (class 0 OID 18708)
-- Dependencies: 238
-- Data for Name: payments; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.payments VALUES (2, 2, 3, 50000.00, 'success', '{"phone": "081234567890", "provider": "GoPay", "transaction_id": "GP123456789"}', '2026-01-11 16:19:39.718247', '2026-01-11 16:19:39.718247');
INSERT INTO public.payments VALUES (3, 3, 1, 50000.00, 'success', '{"bank": "BCA", "last4": "1234", "card_type": "Visa"}', '2026-01-11 16:20:46.328229', '2026-01-11 16:20:46.328229');
INSERT INTO public.payments VALUES (4, 4, 1, 50000.00, 'success', '{"type": "e_wallet", "phone": "081234567890", "provider": "GoPay", "timestamp": "2026-01-11T16:28:59+07:00", "payment_method": "Credit Card", "transaction_id": "GP123456789"}', '2026-01-11 16:28:59.51423', '2026-01-11 16:28:59.51423');
INSERT INTO public.payments VALUES (5, 5, 3, 50000.00, 'success', '{"phone": "081234567890", "provider": "GoPay", "timestamp": "2026-01-11T16:32:15Z", "payment_method": "GoPay", "transaction_id": "GP2026011116321"}', '2026-01-11 16:42:21.370023', '2026-01-11 16:42:21.370023');
INSERT INTO public.payments VALUES (6, 6, 3, 50000.00, 'success', '{"timestamp": "2026-01-11T18:58:55+07:00", "payment_method": "GoPay"}', '2026-01-11 18:58:55.389311', '2026-01-11 18:58:55.389311');


--
-- TOC entry 5150 (class 0 OID 18639)
-- Dependencies: 232
-- Data for Name: seats; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.seats VALUES (1, 1, 'A', 1, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (2, 1, 'A', 2, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (3, 1, 'A', 3, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (4, 1, 'A', 4, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (5, 1, 'A', 5, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (6, 1, 'A', 6, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (7, 1, 'A', 7, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (8, 1, 'A', 8, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (9, 1, 'A', 9, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (10, 1, 'A', 10, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (11, 1, 'B', 1, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (12, 1, 'B', 2, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (13, 1, 'B', 3, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (14, 1, 'B', 4, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (15, 1, 'B', 5, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (16, 1, 'B', 6, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (17, 1, 'B', 7, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (18, 1, 'B', 8, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (19, 1, 'B', 9, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (20, 1, 'B', 10, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (21, 1, 'C', 1, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (22, 1, 'C', 2, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (23, 1, 'C', 3, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (24, 1, 'C', 4, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (25, 1, 'C', 5, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (26, 1, 'C', 6, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (27, 1, 'C', 7, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (28, 1, 'C', 8, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (29, 1, 'C', 9, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (30, 1, 'C', 10, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (31, 1, 'D', 1, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (32, 1, 'D', 2, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (33, 1, 'D', 3, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (34, 1, 'D', 4, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (35, 1, 'D', 5, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (36, 1, 'D', 6, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (37, 1, 'D', 7, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (38, 1, 'D', 8, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (39, 1, 'D', 9, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (40, 1, 'D', 10, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (41, 1, 'E', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (42, 1, 'E', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (43, 1, 'E', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (44, 1, 'E', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (45, 1, 'E', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (46, 1, 'E', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (47, 1, 'E', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (48, 1, 'E', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (49, 1, 'E', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (50, 1, 'E', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (51, 1, 'F', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (52, 1, 'F', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (53, 1, 'F', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (54, 1, 'F', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (55, 1, 'F', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (56, 1, 'F', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (57, 1, 'F', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (58, 1, 'F', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (59, 1, 'F', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (60, 1, 'F', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (61, 1, 'G', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (62, 1, 'G', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (63, 1, 'G', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (64, 1, 'G', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (65, 1, 'G', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (66, 1, 'G', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (67, 1, 'G', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (68, 1, 'G', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (69, 1, 'G', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (70, 1, 'G', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (71, 1, 'H', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (72, 1, 'H', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (73, 1, 'H', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (74, 1, 'H', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (75, 1, 'H', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (76, 1, 'H', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (77, 1, 'H', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (78, 1, 'H', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (79, 1, 'H', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (80, 1, 'H', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (81, 2, 'A', 1, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (82, 2, 'A', 2, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (83, 2, 'A', 3, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (84, 2, 'A', 4, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (85, 2, 'A', 5, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (86, 2, 'A', 6, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (87, 2, 'A', 7, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (88, 2, 'A', 8, 'vip', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (89, 2, 'B', 1, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (90, 2, 'B', 2, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (91, 2, 'B', 3, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (92, 2, 'B', 4, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (93, 2, 'B', 5, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (94, 2, 'B', 6, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (95, 2, 'B', 7, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (96, 2, 'B', 8, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (97, 2, 'C', 1, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (98, 2, 'C', 2, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (99, 2, 'C', 3, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (100, 2, 'C', 4, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (101, 2, 'C', 5, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (102, 2, 'C', 6, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (103, 2, 'C', 7, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (104, 2, 'C', 8, 'premium', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (105, 2, 'D', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (106, 2, 'D', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (107, 2, 'D', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (108, 2, 'D', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (109, 2, 'D', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (110, 2, 'D', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (111, 2, 'D', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (112, 2, 'D', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (113, 2, 'E', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (114, 2, 'E', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (115, 2, 'E', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (116, 2, 'E', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (117, 2, 'E', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (118, 2, 'E', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (119, 2, 'E', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (120, 2, 'E', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (121, 2, 'F', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (122, 2, 'F', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (123, 2, 'F', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (124, 2, 'F', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (125, 2, 'F', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (126, 2, 'F', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (127, 2, 'F', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (128, 2, 'F', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (129, 3, 'A', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (130, 3, 'A', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (131, 3, 'A', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (132, 3, 'A', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (133, 3, 'A', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (134, 3, 'A', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (135, 3, 'A', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (136, 3, 'A', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (137, 3, 'A', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (138, 3, 'A', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (139, 3, 'B', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (140, 3, 'B', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (141, 3, 'B', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (142, 3, 'B', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (143, 3, 'B', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (144, 3, 'B', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (145, 3, 'B', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (146, 3, 'B', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (147, 3, 'B', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (148, 3, 'B', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (149, 3, 'C', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (150, 3, 'C', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (151, 3, 'C', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (152, 3, 'C', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (153, 3, 'C', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (154, 3, 'C', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (155, 3, 'C', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (156, 3, 'C', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (157, 3, 'C', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (158, 3, 'C', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (159, 3, 'D', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (160, 3, 'D', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (161, 3, 'D', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (162, 3, 'D', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (163, 3, 'D', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (164, 3, 'D', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (165, 3, 'D', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (166, 3, 'D', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (167, 3, 'D', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (168, 3, 'D', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (169, 3, 'E', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (170, 3, 'E', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (171, 3, 'E', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (172, 3, 'E', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (173, 3, 'E', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (174, 3, 'E', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (175, 3, 'E', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (176, 3, 'E', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (177, 3, 'E', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (178, 3, 'E', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (179, 3, 'F', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (180, 3, 'F', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (181, 3, 'F', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (182, 3, 'F', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (183, 3, 'F', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (184, 3, 'F', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (185, 3, 'F', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (186, 3, 'F', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (187, 3, 'F', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (188, 3, 'F', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (189, 4, 'A', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (190, 4, 'A', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (191, 4, 'A', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (192, 4, 'A', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (193, 4, 'A', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (194, 4, 'A', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (195, 4, 'A', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (196, 4, 'A', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (197, 4, 'A', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (198, 4, 'A', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (199, 4, 'B', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (200, 4, 'B', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (201, 4, 'B', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (202, 4, 'B', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (203, 4, 'B', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (204, 4, 'B', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (205, 4, 'B', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (206, 4, 'B', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (207, 4, 'B', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (208, 4, 'B', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (209, 4, 'C', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (210, 4, 'C', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (211, 4, 'C', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (212, 4, 'C', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (213, 4, 'C', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (214, 4, 'C', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (215, 4, 'C', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (216, 4, 'C', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (217, 4, 'C', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (218, 4, 'C', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (219, 4, 'D', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (220, 4, 'D', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (221, 4, 'D', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (222, 4, 'D', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (223, 4, 'D', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (224, 4, 'D', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (225, 4, 'D', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (226, 4, 'D', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (227, 4, 'D', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (228, 4, 'D', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (229, 4, 'E', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (230, 4, 'E', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (231, 4, 'E', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (232, 4, 'E', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (233, 4, 'E', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (234, 4, 'E', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (235, 4, 'E', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (236, 4, 'E', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (237, 4, 'E', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (238, 4, 'E', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (239, 4, 'F', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (240, 4, 'F', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (241, 4, 'F', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (242, 4, 'F', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (243, 4, 'F', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (244, 4, 'F', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (245, 4, 'F', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (246, 4, 'F', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (247, 4, 'F', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (248, 4, 'F', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (249, 5, 'A', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (250, 5, 'A', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (251, 5, 'A', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (252, 5, 'A', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (253, 5, 'A', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (254, 5, 'A', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (255, 5, 'A', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (256, 5, 'A', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (257, 5, 'A', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (258, 5, 'A', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (259, 5, 'B', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (260, 5, 'B', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (261, 5, 'B', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (262, 5, 'B', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (263, 5, 'B', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (264, 5, 'B', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (265, 5, 'B', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (266, 5, 'B', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (267, 5, 'B', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (268, 5, 'B', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (269, 5, 'C', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (270, 5, 'C', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (271, 5, 'C', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (272, 5, 'C', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (273, 5, 'C', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (274, 5, 'C', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (275, 5, 'C', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (276, 5, 'C', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (277, 5, 'C', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (278, 5, 'C', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (279, 5, 'D', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (280, 5, 'D', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (281, 5, 'D', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (282, 5, 'D', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (283, 5, 'D', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (284, 5, 'D', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (285, 5, 'D', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (286, 5, 'D', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (287, 5, 'D', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (288, 5, 'D', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (289, 5, 'E', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (290, 5, 'E', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (291, 5, 'E', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (292, 5, 'E', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (293, 5, 'E', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (294, 5, 'E', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (295, 5, 'E', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (296, 5, 'E', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (297, 5, 'E', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (298, 5, 'E', 10, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (299, 5, 'F', 1, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (300, 5, 'F', 2, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (301, 5, 'F', 3, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (302, 5, 'F', 4, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (303, 5, 'F', 5, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (304, 5, 'F', 6, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (305, 5, 'F', 7, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (306, 5, 'F', 8, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (307, 5, 'F', 9, 'regular', '2026-01-11 14:24:20.203825');
INSERT INTO public.seats VALUES (308, 5, 'F', 10, 'regular', '2026-01-11 14:24:20.203825');


--
-- TOC entry 5148 (class 0 OID 18613)
-- Dependencies: 230
-- Data for Name: showtimes; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.showtimes VALUES (1, 1, 1, '2026-01-11', '10:00:00', 50000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (2, 1, 1, '2026-01-11', '14:00:00', 50000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (3, 1, 1, '2026-01-11', '19:00:00', 75000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (4, 1, 2, '2026-01-11', '11:30:00', 50000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (5, 1, 2, '2026-01-11', '16:00:00', 50000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (6, 1, 3, '2026-01-11', '13:00:00', 50000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (7, 1, 3, '2026-01-11', '20:30:00', 75000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (8, 2, 4, '2026-01-11', '10:30:00', 55000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (9, 2, 4, '2026-01-11', '15:00:00', 55000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (10, 2, 5, '2026-01-11', '12:00:00', 60000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (11, 2, 5, '2026-01-11', '17:30:00', 60000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (12, 2, 5, '2026-01-11', '21:00:00', 80000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (13, 1, 1, '2026-01-12', '10:00:00', 50000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (14, 1, 2, '2026-01-12', '13:00:00', 50000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (15, 1, 3, '2026-01-12', '16:00:00', 50000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (16, 2, 4, '2026-01-12', '11:00:00', 55000.00, '2026-01-11 14:24:20.203825');
INSERT INTO public.showtimes VALUES (17, 2, 5, '2026-01-12', '14:30:00', 60000.00, '2026-01-11 14:24:20.203825');


--
-- TOC entry 5138 (class 0 OID 18532)
-- Dependencies: 220
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES (1, 'johndoe', 'john@example.com', '$2a$10$cmJgOK/j48e.btEMdB81SOdWkMVhZOfQtsNY2iJOGX/btA8VKzVX.', false, '2026-01-11 14:52:59.074591', '2026-01-11 14:52:59.074591');
INSERT INTO public.users VALUES (3, 'anas', '202210715282@mhs.ubharajaya.ac.id', '$2a$10$04Qxn33Dx6CjvngIxOCMmuhv6dNsw97sTW6e5QHSpJ9TjMPNihzxW', false, '2026-01-11 19:31:37.189389', '2026-01-11 19:31:37.189389');
INSERT INTO public.users VALUES (4, 'experiment', 'Saqib.Pigott@AllWebEmails.com', '$2a$10$Cj/FNudoC1.4JRydxD/Nl.bXByRri56n6stzWezRFhCpVMyWrPKlS', true, '2026-01-11 19:37:20.548533', '2026-01-11 19:38:21.732632');


--
-- TOC entry 5172 (class 0 OID 0)
-- Dependencies: 221
-- Name: auth_tokens_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.auth_tokens_id_seq', 9, true);


--
-- TOC entry 5173 (class 0 OID 0)
-- Dependencies: 235
-- Name: bookings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bookings_id_seq', 6, true);


--
-- TOC entry 5174 (class 0 OID 0)
-- Dependencies: 225
-- Name: cinemas_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cinemas_id_seq', 5, true);


--
-- TOC entry 5175 (class 0 OID 0)
-- Dependencies: 227
-- Name: movies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.movies_id_seq', 5, true);


--
-- TOC entry 5176 (class 0 OID 0)
-- Dependencies: 223
-- Name: otp_codes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.otp_codes_id_seq', 3, true);


--
-- TOC entry 5177 (class 0 OID 0)
-- Dependencies: 233
-- Name: payment_methods_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.payment_methods_id_seq', 7, true);


--
-- TOC entry 5178 (class 0 OID 0)
-- Dependencies: 237
-- Name: payments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.payments_id_seq', 6, true);


--
-- TOC entry 5179 (class 0 OID 0)
-- Dependencies: 231
-- Name: seats_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.seats_id_seq', 308, true);


--
-- TOC entry 5180 (class 0 OID 0)
-- Dependencies: 229
-- Name: showtimes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.showtimes_id_seq', 17, true);


--
-- TOC entry 5181 (class 0 OID 0)
-- Dependencies: 219
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 4, true);


--
-- TOC entry 4937 (class 2606 OID 18560)
-- Name: auth_tokens auth_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_tokens
    ADD CONSTRAINT auth_tokens_pkey PRIMARY KEY (id);


--
-- TOC entry 4939 (class 2606 OID 18562)
-- Name: auth_tokens auth_tokens_token_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_tokens
    ADD CONSTRAINT auth_tokens_token_key UNIQUE (token);


--
-- TOC entry 4966 (class 2606 OID 18689)
-- Name: bookings bookings_booking_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_booking_code_key UNIQUE (booking_code);


--
-- TOC entry 4968 (class 2606 OID 18687)
-- Name: bookings bookings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_pkey PRIMARY KEY (id);


--
-- TOC entry 4970 (class 2606 OID 18691)
-- Name: bookings bookings_showtime_id_seat_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_showtime_id_seat_id_key UNIQUE (showtime_id, seat_id);


--
-- TOC entry 4946 (class 2606 OID 18598)
-- Name: cinemas cinemas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas
    ADD CONSTRAINT cinemas_pkey PRIMARY KEY (id);


--
-- TOC entry 4948 (class 2606 OID 18611)
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- TOC entry 4944 (class 2606 OID 18580)
-- Name: otp_codes otp_codes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.otp_codes
    ADD CONSTRAINT otp_codes_pkey PRIMARY KEY (id);


--
-- TOC entry 4962 (class 2606 OID 18671)
-- Name: payment_methods payment_methods_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods
    ADD CONSTRAINT payment_methods_code_key UNIQUE (code);


--
-- TOC entry 4964 (class 2606 OID 18669)
-- Name: payment_methods payment_methods_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods
    ADD CONSTRAINT payment_methods_pkey PRIMARY KEY (id);


--
-- TOC entry 4977 (class 2606 OID 18721)
-- Name: payments payments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (id);


--
-- TOC entry 4958 (class 2606 OID 18652)
-- Name: seats seats_cinema_id_seat_row_seat_number_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seats_cinema_id_seat_row_seat_number_key UNIQUE (cinema_id, seat_row, seat_number);


--
-- TOC entry 4960 (class 2606 OID 18650)
-- Name: seats seats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seats_pkey PRIMARY KEY (id);


--
-- TOC entry 4953 (class 2606 OID 18627)
-- Name: showtimes showtimes_cinema_id_movie_id_show_date_show_time_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_cinema_id_movie_id_show_date_show_time_key UNIQUE (cinema_id, movie_id, show_date, show_time);


--
-- TOC entry 4955 (class 2606 OID 18625)
-- Name: showtimes showtimes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_pkey PRIMARY KEY (id);


--
-- TOC entry 4931 (class 2606 OID 18548)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 4933 (class 2606 OID 18544)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4935 (class 2606 OID 18546)
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- TOC entry 4940 (class 1259 OID 18734)
-- Name: idx_auth_tokens_expires_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_auth_tokens_expires_at ON public.auth_tokens USING btree (expires_at);


--
-- TOC entry 4941 (class 1259 OID 18733)
-- Name: idx_auth_tokens_token; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_auth_tokens_token ON public.auth_tokens USING btree (token);


--
-- TOC entry 4942 (class 1259 OID 18732)
-- Name: idx_auth_tokens_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_auth_tokens_user_id ON public.auth_tokens USING btree (user_id);


--
-- TOC entry 4971 (class 1259 OID 18739)
-- Name: idx_bookings_showtime_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bookings_showtime_id ON public.bookings USING btree (showtime_id);


--
-- TOC entry 4972 (class 1259 OID 18740)
-- Name: idx_bookings_status; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bookings_status ON public.bookings USING btree (status);


--
-- TOC entry 4973 (class 1259 OID 18738)
-- Name: idx_bookings_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bookings_user_id ON public.bookings USING btree (user_id);


--
-- TOC entry 4974 (class 1259 OID 18742)
-- Name: idx_payments_booking_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_payments_booking_id ON public.payments USING btree (booking_id);


--
-- TOC entry 4975 (class 1259 OID 18743)
-- Name: idx_payments_status; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_payments_status ON public.payments USING btree (status);


--
-- TOC entry 4956 (class 1259 OID 18741)
-- Name: idx_seats_cinema_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_seats_cinema_id ON public.seats USING btree (cinema_id);


--
-- TOC entry 4949 (class 1259 OID 18735)
-- Name: idx_showtimes_cinema_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_showtimes_cinema_id ON public.showtimes USING btree (cinema_id);


--
-- TOC entry 4950 (class 1259 OID 18736)
-- Name: idx_showtimes_movie_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_showtimes_movie_id ON public.showtimes USING btree (movie_id);


--
-- TOC entry 4951 (class 1259 OID 18737)
-- Name: idx_showtimes_show_date; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_showtimes_show_date ON public.showtimes USING btree (show_date);


--
-- TOC entry 4989 (class 2620 OID 18746)
-- Name: bookings update_bookings_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_bookings_updated_at BEFORE UPDATE ON public.bookings FOR EACH ROW EXECUTE FUNCTION public.update_updated_at_column();


--
-- TOC entry 4988 (class 2620 OID 18745)
-- Name: users update_users_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.update_updated_at_column();


--
-- TOC entry 4978 (class 2606 OID 18563)
-- Name: auth_tokens auth_tokens_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_tokens
    ADD CONSTRAINT auth_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- TOC entry 4983 (class 2606 OID 18702)
-- Name: bookings bookings_seat_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_seat_id_fkey FOREIGN KEY (seat_id) REFERENCES public.seats(id) ON DELETE CASCADE;


--
-- TOC entry 4984 (class 2606 OID 18697)
-- Name: bookings bookings_showtime_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_showtime_id_fkey FOREIGN KEY (showtime_id) REFERENCES public.showtimes(id) ON DELETE CASCADE;


--
-- TOC entry 4985 (class 2606 OID 18692)
-- Name: bookings bookings_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- TOC entry 4979 (class 2606 OID 18581)
-- Name: otp_codes otp_codes_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.otp_codes
    ADD CONSTRAINT otp_codes_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- TOC entry 4986 (class 2606 OID 18722)
-- Name: payments payments_booking_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_booking_id_fkey FOREIGN KEY (booking_id) REFERENCES public.bookings(id) ON DELETE CASCADE;


--
-- TOC entry 4987 (class 2606 OID 18727)
-- Name: payments payments_payment_method_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_payment_method_id_fkey FOREIGN KEY (payment_method_id) REFERENCES public.payment_methods(id);


--
-- TOC entry 4982 (class 2606 OID 18653)
-- Name: seats seats_cinema_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seats_cinema_id_fkey FOREIGN KEY (cinema_id) REFERENCES public.cinemas(id) ON DELETE CASCADE;


--
-- TOC entry 4980 (class 2606 OID 18628)
-- Name: showtimes showtimes_cinema_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_cinema_id_fkey FOREIGN KEY (cinema_id) REFERENCES public.cinemas(id) ON DELETE CASCADE;


--
-- TOC entry 4981 (class 2606 OID 18633)
-- Name: showtimes showtimes_movie_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id) ON DELETE CASCADE;


-- Completed on 2026-01-11 20:49:14