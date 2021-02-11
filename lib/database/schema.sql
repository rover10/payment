-- Postgres database
create table users (id bigserial unique,
	upi_id varchar(50) unique,
	mob_number varchar(10),
	first_name varchar(50) not null,
	mid_name varchar(50),
	last_name varchar(50),
	country_id integer not null,
	primary_account_id integer,
	pan varchar(9),
	pin varchar(50) not null,
	create_on timestamp not null,
	max_transaction_limit numeric(15,4) not null
);

create table bank(
	id serial primary key,
	name varchar(200) not null unique,
	code varchar(100) not null unique,
	icon_url varchar(1000) not null unique
)

create table users_account(
	id bigserial unique,
    bank_id integer references bank(id) not null,
	account_number varchar(20),
	ifsc varchar(11),
	user_id bigserial references users(id),
	PRIMARY KEY(account_number, ifsc)
);

create table transaction(
	id bigserial unique,
	utr varchar(22) unique,
	amount numeric(15,4) not null check (amount > 0) ,
	from_account_id bigserial not null references users_account(id) check (from_account_id <> to_account_id),
	to_account_id bigserial not null references users_account(id),
	payment_time timestamp not null,
	payment_gateway varchar(50) not null,
	status varchar(50)
);

