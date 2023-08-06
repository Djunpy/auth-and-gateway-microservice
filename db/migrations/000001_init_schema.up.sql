CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(80) NOT NULL,
    email VARCHAR(120) NOT NULL,
    verified_email BOOL DEFAULT false,
    photo VARCHAR(255),
    password VARCHAR(120) NOT NULL,
    first_name VARCHAR(120),
    last_name VARCHAR(120),
    is_staff BOOL DEFAULT false,
    is_active BOOL DEFAULT false,
    is_superuser BOOL DEFAULT false,
    is_deleted BOOL DEFAULT false,
    auth_source VARCHAR(120) NOT NULL,
    update_at TIMESTAMPTZ,
    date_joined TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE phones (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    number INTEGER NOT NULL,
    country_code VARCHAR(30) NOT NULL,
    verified BOOL DEFAULT false,
    create_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id),
    UNIQUE (number)
);

CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    codename VARCHAR(255)
);

CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE user_groups (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (group_id) REFERENCES groups(id),
    create_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE (user_id, group_id)
);

CREATE TABLE user_permissions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    permission_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (permission_id) REFERENCES permissions(id),
    create_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE (user_id, permission_id)
);

CREATE TABLE chats (
   id SERIAL PRIMARY KEY,
   is_deleted BOOL DEFAULT false,
   create_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE chat_participants (
   id SERIAL PRIMARY KEY,
   user_id INTEGER,
   chat_id INTEGER NOT NULL,
   is_deleted BOOL DEFAULT false,
   FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
   UNIQUE (user_id, chat_id)
);

CREATE TABLE messages (
  id SERIAL PRIMARY KEY,
  chat_id INTEGER NOT NULL,
  sender_id INTEGER,
  text_message TEXT NOT NULL,
  is_deleted BOOL DEFAULT false,
  sent_at TIMESTAMPTZ DEFAULT NOW(),
  FOREIGN KEY (chat_id) REFERENCES chats(id),
  FOREIGN KEY (sender_id) REFERENCES users(id)
);

CREATE TABLE message_status (
    id SERIAL PRIMARY KEY,
    message_id INTEGER NOT NULL,
    is_read BOOL DEFAULT false,
    is_delivered BOOL DEFAULT false,
    FOREIGN KEY (message_id) REFERENCES messages(id),
    UNIQUE (message_id)
);

CREATE TABLE files (
   id SERIAL PRIMARY KEY,
   user_id INTEGER NOT NULL,
   filename VARCHAR(100),
   filepath VARCHAR(255) NOT NULL,
   FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE message_files (
    id SERIAL PRIMARY KEY,
    message_id INTEGER NOT NULL,
    file_id INTEGER NOT NULL,
    UNIQUE (message_id, file_id)
);

CREATE TABLE user_ratings (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  rater_id INTEGER,
  rating_value INTEGER CHECK ( rating_value >= 1 AND rating_value <= 5 ) NOT NULL,
  comment TEXT NOT NULL,
  is_deleted BOOL DEFAULT false,
  create_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE reply_to_ratings (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  rating_id INTEGER NOT NULL,
  reply_text TEXT NOT NULL,
  create_at TIMESTAMPTZ DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (rating_id) REFERENCES user_ratings(id)
);