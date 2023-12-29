CREATE TABLE sushi_users (
    id smallint not null auto_increment,
    phone varchar(255) not null,
    name varchar(255) not null,
    address JSON,
    primary key(id)
);


CREATE TABLE sushi_orders (
    id smallint not null auto_increment,
    phone varchar(255) not null,
    name varchar(255) not null,
    address varchar(255) not null,
    products JSON not null,
    price smallint not null,
    date timestamp,
    primary key(id)
);

CREATE TABLE sushi_products(
    id smallint not null auto_increment,
    name varchar(255) not null,
    image varchar(255) not null,
    category smallint not null,
    price smallint not null,
    weight smallint not null,
    description varchar(255) not null,
    nutrition JSON,
    primary key(id)
);

CREATE TABLE sushi_cart(
    userid smallint not null,
    productid smallint not null,
    foreign key (userid) references sushi_users(id) on delete cascade,
    foreign key (productid) references sushi_products(id) on delete cascade on update cascade
);


CREATE TABLE sushi_categories(
    id smallint not null,
    name varchar(255) not null,
    primary key(id)
);



