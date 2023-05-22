CREATE TABLE borrower
(
    id bigserial NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
	alamat character varying(255) NOT NULL,
	phone_number character varying(20) NOT NULL,
	created_at timestamptz NOT NULL,
    CONSTRAINT borrower_pkey PRIMARY KEY (id)
);

CREATE TABLE lender
(
    id bigserial NOT NULL,
    name character varying(255) NOT NULL,
	created_at timestamptz NOT NULL,
    CONSTRAINT lender_pkey PRIMARY KEY (id)
);

CREATE TABLE loan_product
(
    id bigserial NOT NULL,
    name character varying(255) NOT NULL,
	description character varying(255) NOT NULL,
	persyaratan character varying(255) NOT NULL,
	created_at timestamptz NOT NULL,
    CONSTRAINT loan_product_pkey PRIMARY KEY (id)
);

CREATE TABLE session
(
    id bigserial NOT NULL,
    username character varying(255) NOT NULL,
	token character varying(255) NOT NULL,
	expired_at timestamptz NOT NULL,
	created_at timestamptz NOT NULL,
    CONSTRAINT session_pkey PRIMARY KEY (id)
);

CREATE TABLE transaction 
(
  id bigserial PRIMARY KEY NOT NULL,
  id_borrower int REFERENCES borrower(id) NOT NULL,
  id_lender int REFERENCES lender(id) NOT NULL,
  id_loan_product int REFERENCES loan_product(id) NOT NULL,
  loan_amount decimal(10, 2) NOT NULL,
  transaction_date timestamptz NOT NULL,
  due_date timestamptz NOT NULL,
  transaction_status varchar(255) NOT NULL
);


CREATE TABLE loan_history
(
    id bigserial NOT NULL,
    id_transaksi int REFERENCES transaction(id) NOT NULL,
	history_state character varying(255) NOT NULL,
	information character varying(255) NOT NULL,
	change_date timestamptz NOT NULL,
	created_at timestamptz NOT NULL,
    CONSTRAINT loan_history_pkey PRIMARY KEY (id)
);

CREATE TABLE payment
(
    id bigserial NOT NULL,
    id_transaksi int REFERENCES transaction(id) NOT NULL,
	payment_amount character varying(255) NOT NULL,
	payment_date timestamptz NOT NULL,
	payment_method character varying(255) NOT NULL,
    CONSTRAINT payment_pkey PRIMARY KEY (id)
);