CREATE TABLE user (
    id int constraint pk_user primary key autoincrement,
    username varchar constraint un_user_username unique,
    password varchar constraint nn_user_password not null,
    role varchar constraint nn_user_role not null
);