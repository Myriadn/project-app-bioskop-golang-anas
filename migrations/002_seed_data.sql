-- Insert Payment Methods
INSERT INTO payment_methods (name, code, is_active) VALUES
    ('Credit Card', 'CREDIT_CARD', TRUE),
    ('Debit Card', 'DEBIT_CARD', TRUE),
    ('GoPay', 'GOPAY', TRUE),
    ('OVO', 'OVO', TRUE),
    ('DANA', 'DANA', TRUE),
    ('ShopeePay', 'SHOPEEPAY', TRUE),
    ('Bank Transfer', 'BANK_TRANSFER', TRUE);

-- Insert Cinemas
INSERT INTO cinemas (name, location, description) VALUES
    ('CGV Grand Indonesia', 'Jakarta Pusat', 'Cinema dengan fasilitas premium di pusat kota Jakarta'),
    ('XXI Plaza Senayan', 'Jakarta Selatan', 'Bioskop modern dengan teknologi terkini'),
    ('Cinepolis Lippo Mall Puri', 'Jakarta Barat', 'Cinema dengan konsep luxury dan VIP lounge'),
    ('CGV Blitz Megaplex', 'Jakarta Timur', 'Bioskop dengan berbagai pilihan film dan seat type'),
    ('XXI Pondok Indah Mall', 'Jakarta Selatan', 'Cinema dengan sound system Dolby Atmos');

-- Insert Movies
INSERT INTO movies (title, description, duration, genre, poster_url, rating) VALUES
    ('The Dark Knight', 'Batman melawan Joker dalam pertarungan epik untuk menyelamatkan Gotham City', 152, 'Action, Crime, Drama', 'https://example.com/dark-knight.jpg', '13+'),
    ('Inception', 'Seorang pencuri yang mencuri rahasia korporat melalui dream-sharing technology', 148, 'Action, Sci-Fi, Thriller', 'https://example.com/inception.jpg', '13+'),
    ('Interstellar', 'Sekelompok penjelajah menggunakan lubang cacing untuk melintasi dimensi ruang', 169, 'Adventure, Drama, Sci-Fi', 'https://example.com/interstellar.jpg', '13+'),
    ('Parasite', 'Keluarga miskin yang menyusup ke kehidupan keluarga kaya', 132, 'Comedy, Drama, Thriller', 'https://example.com/parasite.jpg', '17+'),
    ('Avengers: Endgame', 'Para Avengers berkumpul untuk mengalahkan Thanos sekali dan untuk selamanya', 181, 'Action, Adventure, Sci-Fi', 'https://example.com/endgame.jpg', '13+');

-- Insert Showtimes untuk 7 hari ke depan
INSERT INTO showtimes (cinema_id, movie_id, show_date, show_time, price) VALUES
    -- CGV Grand Indonesia
    (1, 1, CURRENT_DATE, '10:00:00', 50000),
    (1, 1, CURRENT_DATE, '14:00:00', 50000),
    (1, 1, CURRENT_DATE, '19:00:00', 75000),
    (1, 2, CURRENT_DATE, '11:30:00', 50000),
    (1, 2, CURRENT_DATE, '16:00:00', 50000),
    (1, 3, CURRENT_DATE, '13:00:00', 50000),
    (1, 3, CURRENT_DATE, '20:30:00', 75000),

    -- XXI Plaza Senayan
    (2, 4, CURRENT_DATE, '10:30:00', 55000),
    (2, 4, CURRENT_DATE, '15:00:00', 55000),
    (2, 5, CURRENT_DATE, '12:00:00', 60000),
    (2, 5, CURRENT_DATE, '17:30:00', 60000),
    (2, 5, CURRENT_DATE, '21:00:00', 80000),

    -- Untuk besok
    (1, 1, CURRENT_DATE + 1, '10:00:00', 50000),
    (1, 2, CURRENT_DATE + 1, '13:00:00', 50000),
    (1, 3, CURRENT_DATE + 1, '16:00:00', 50000),
    (2, 4, CURRENT_DATE + 1, '11:00:00', 55000),
    (2, 5, CURRENT_DATE + 1, '14:30:00', 60000);

-- Insert Seats untuk Cinema 1 (CGV Grand Indonesia)
-- Row A (10 seats - VIP)
INSERT INTO seats (cinema_id, seat_row, seat_number, seat_type)
SELECT 1, 'A', generate_series, 'vip'
FROM generate_series(1, 10);

-- Row B-D (30 seats each - Premium)
INSERT INTO seats (cinema_id, seat_row, seat_number, seat_type)
SELECT 1, seat_row, seat_number, 'premium'
FROM (
    SELECT 'B' as seat_row, generate_series(1, 10) as seat_number
    UNION ALL
    SELECT 'C', generate_series(1, 10)
    UNION ALL
    SELECT 'D', generate_series(1, 10)
) s;

-- Row E-H (40 seats each - Regular)
INSERT INTO seats (cinema_id, seat_row, seat_number, seat_type)
SELECT 1, seat_row, seat_number, 'regular'
FROM (
    SELECT 'E' as seat_row, generate_series(1, 10) as seat_number
    UNION ALL
    SELECT 'F', generate_series(1, 10)
    UNION ALL
    SELECT 'G', generate_series(1, 10)
    UNION ALL
    SELECT 'H', generate_series(1, 10)
) s;

-- Insert Seats untuk Cinema 2 (XXI Plaza Senayan)
INSERT INTO seats (cinema_id, seat_row, seat_number, seat_type)
SELECT 2, seat_row, seat_number, seat_type
FROM (
    SELECT 'A' as seat_row, generate_series(1, 8) as seat_number, 'vip' as seat_type
    UNION ALL
    SELECT 'B', generate_series(1, 8), 'premium'
    UNION ALL
    SELECT 'C', generate_series(1, 8), 'premium'
    UNION ALL
    SELECT 'D', generate_series(1, 8), 'regular'
    UNION ALL
    SELECT 'E', generate_series(1, 8), 'regular'
    UNION ALL
    SELECT 'F', generate_series(1, 8), 'regular'
) s;

-- Insert Seats untuk Cinema 3 (Cinepolis Lippo Mall Puri)
INSERT INTO seats (cinema_id, seat_row, seat_number, seat_type)
SELECT 3, seat_row, seat_num, 'regular'
FROM (
    SELECT 'A' as seat_row, generate_series(1, 10) as seat_num
    UNION ALL
    SELECT 'B', generate_series(1, 10)
    UNION ALL
    SELECT 'C', generate_series(1, 10)
    UNION ALL
    SELECT 'D', generate_series(1, 10)
    UNION ALL
    SELECT 'E', generate_series(1, 10)
    UNION ALL
    SELECT 'F', generate_series(1, 10)
) s;

-- Insert Seats untuk Cinema 4 (CGV Blitz Megaplex)
INSERT INTO seats (cinema_id, seat_row, seat_number, seat_type)
SELECT 4, seat_row, seat_num, 'regular'
FROM (
    SELECT 'A' as seat_row, generate_series(1, 10) as seat_num
    UNION ALL
    SELECT 'B', generate_series(1, 10)
    UNION ALL
    SELECT 'C', generate_series(1, 10)
    UNION ALL
    SELECT 'D', generate_series(1, 10)
    UNION ALL
    SELECT 'E', generate_series(1, 10)
    UNION ALL
    SELECT 'F', generate_series(1, 10)
) s;

-- Insert Seats untuk Cinema 5 (XXI Pondok Indah Mall)
INSERT INTO seats (cinema_id, seat_row, seat_number, seat_type)
SELECT 5, seat_row, seat_num, 'regular'
FROM (
    SELECT 'A' as seat_row, generate_series(1, 10) as seat_num
    UNION ALL
    SELECT 'B', generate_series(1, 10)
    UNION ALL
    SELECT 'C', generate_series(1, 10)
    UNION ALL
    SELECT 'D', generate_series(1, 10)
    UNION ALL
    SELECT 'E', generate_series(1, 10)
    UNION ALL
    SELECT 'F', generate_series(1, 10)
) s;
