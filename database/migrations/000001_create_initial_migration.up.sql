create table links (
    id SERIAL PRIMARY KEY,
    url_code TEXT NOT NULL,
    destination TEXT NOT NULL UNIQUE
);

