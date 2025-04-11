CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id),
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    total DECIMAL(10, 2) NOT NULL
);

-- Insert 50 random customers
INSERT INTO customers (name, email)
SELECT
    -- Generate name once and store it
    first_name || ' ' || last_name AS name,
    -- Use same name values for email in snake_case
    LOWER(first_name || '_' || last_name) || '_' || i || '@gmail.com' AS email
FROM (
    SELECT
        (ARRAY['John', 'Mary', 'James', 'Sarah', 'Michael', 'Emma', 'David', 'Lisa', 'Robert', 'Anna'])[floor(random() * 10 + 1)::integer] AS first_name,
        (ARRAY['Smith', 'Johnson', 'Williams', 'Brown', 'Jones', 'Miller', 'Davis', 'Wilson', 'Taylor', 'Clark'])[floor(random() * 10 + 1)::integer] AS last_name,
        i
    FROM generate_series(1, 50) AS i
) AS name_parts;

-- Insert 200 random orders linked to the customers
INSERT INTO orders (customer_id, order_date, total)
SELECT
    (random() * 49 + 1)::integer AS customer_id,
    CURRENT_TIMESTAMP - (random() * 365 || ' days')::interval AS order_date,
    (random() * 1000 + 10)::numeric(10,2) AS total
FROM generate_series(1, 200);

-- Setup users and permissions
CREATE USER strimq_readonly_user WITH PASSWORD 'strimq_readonly_user_password';
GRANT USAGE ON SCHEMA public TO strimq_readonly_user;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO strimq_readonly_user;

CREATE USER strimq_readwrite_user WITH PASSWORD 'strimq_readwrite_user_password';
GRANT USAGE ON SCHEMA public TO strimq_readwrite_user;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO strimq_readwrite_user;

-- Create publication
CREATE PUBLICATION strimq_publication FOR ALL TABLES;
SELECT PG_CREATE_LOGICAL_REPLICATION_SLOT('strimq_replication_slot', 'pgoutput');
