-- テーブルの作成: authors
CREATE TABLE authors (
  id integer PRIMARY KEY autoincrement,
  name VARCHAR(255) NOT NULL,
  country VARCHAR(255) NOT NULL
);

-- テーブルの作成: books
CREATE TABLE books (
  id integer  PRIMARY KEY autoincrement,
  title VARCHAR(255) NOT NULL,
  author_id INTEGER REFERENCES authors(id),
  publication_year INTEGER,
  available BOOLEAN DEFAULT TRUE
);

-- テーブルの作成: users
CREATE TABLE users (
  id integer PRIMARY KEY autoincrement,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  birth_date DATE NOT NULL
);

-- テーブルの作成: loans
CREATE TABLE loans (
  id integer PRIMARY KEY autoincrement,
  user_id INTEGER REFERENCES users(id),
  book_id INTEGER REFERENCES books(id),
  loan_date DATE NOT NULL,
  due_date DATE NOT NULL,
  returned BOOLEAN DEFAULT FALSE
);

-- データの挿入: authors
INSERT INTO authors (name, country) VALUES ('夏目 漱石', 'Japan');
INSERT INTO authors (name, country) VALUES ('芥川 龍之介', 'Japan');
INSERT INTO authors (name, country) VALUES ('宮沢 賢治', 'Japan');

-- データの挿入: books
INSERT INTO books (title, author_id, publication_year) VALUES ('吾輩は猫である', 1, 1905);
INSERT INTO books (title, author_id, publication_year) VALUES ('羅生門', 2, 1915);
INSERT INTO books (title, author_id, publication_year) VALUES ('銀河鉄道の夜', 3, 1934);

-- データの挿入: users
INSERT INTO users (name, email, birth_date) VALUES ('山田 太郎', 'taro.yamada@example.com', '2000-01-01');
INSERT INTO users (name, email, birth_date) VALUES ('鈴木 次郎', 'jiro.suzuki@example.com', '1995-02-14');
INSERT INTO users (name, email, birth_date) VALUES ('佐藤 三郎', 'saburo.sato@example.com', '1980-12-31');

-- データの挿入: loans
INSERT INTO loans (user_id, book_id, loan_date, due_date) VALUES (1, 1, '2022-01-01', '2022-01-15');
INSERT INTO loans (user_id, book_id, loan_date, due_date) VALUES (1, 2, '2022-01-01', '2022-01-15');
INSERT INTO loans (user_id, book_id, loan_date, due_date) VALUES (2, 3, '2022-01-05', '2022-01-19');
