-- A. Munculkan data country mana aja yang spend nya terbanyak (Query)
SELECT 
    t1.country,
    SUM(t2.total_buy) AS total_spend,
    COUNT(DISTINCT t2.id_user) AS total_users,
    COUNT(t2.id) AS total_transactions,
    AVG(t2.total_buy) AS avg_spend_per_transaction
FROM 
    public."user" t1 
INNER JOIN 
    public.belanja t2 ON t1.id = t2.id_user
GROUP BY 
    t1.country
ORDER BY 
    total_spend DESC
LIMIT 10;

-- B. Munculkan data jumlah tipe kartu kredit terbanyak (Query)
SELECT 
    credit_card_type,
    COUNT(*) AS jumlah_pengguna,
    COUNT(DISTINCT id) AS total_unique_users
FROM 
    public."user"
GROUP BY 
    credit_card_type
ORDER BY 
    jumlah_pengguna DESC;

-- C. Buatlah GET API untuk medapatkan data dari soal a dengan response menampilkan : Id, Country, Credit_card_type, Credit_card, First_name, Last_name
-- Dalam main.go

-- D. Buatlah POST API untuk mengirim data dengan body seperti dibawah ini :
-- {
--     "country”: " ",
--     "credit_card_type”: " ",
--     "credit_card”: " ",
--     "first_name”: " ",
--     "last_name”: " "
-- }
-- Dalam main.go