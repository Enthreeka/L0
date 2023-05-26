create table "order" (
    order_uid varchar(20),
    track_number varchar(50),
    entry varchar(20),
    locale varchar(4),
    internal_signature varchar(50),
    customer_id  varchar(30),
    delivery_service varchar(20),
    shardkey varchar(10),
    sm_id integer,
    date_created date,
    oof_shard varchar(10),
    primary key (order_uid)
);


create table delivery(
    order_uid varchar(20),
    phone varchar(12),
    name VARCHAR(70),
    zip varchar(15),
    city varchar(30),
    address varchar(30),
    region varchar(30),
    email varchar(30),
    primary key(phone),
    foreign key(order_uid)
        references "order" (order_uid)
);


create table payment(
    order_uid varchar(20),
    transaction varchar(30),
    request_id varchar(20),
    currency varchar(3),
    provider varchar(20),
    amount integer,
    payment_dt integer,
    bank varchar(20),
    delivery_cost integer,
    goods_total integer,
    custom_fee integer,
    primary key(transaction),
    foreign key(order_uid)
        references "order" (order_uid)
);

create table item(
      chrt_id integer,
      track_number varchar(20),
      order_uid varchar(20),
      price integer,
      rid varchar(30),
      name varchar(100),
      sale integer,
      size varchar(5),
      total_price integer,
      nm_id integer,
      brand varchar(50),
      status varchar(3),
      primary key(track_number),
      foreign key(order_uid)
        references "order" (order_uid)

); 

