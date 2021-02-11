----- Postgres database
----- Create banks
INSERT INTO public.bank
("name", code, icon_url)
VALUES('State Bank of India', 'SBI', 'http://icon.mobile.matchmove/mobile/sbi.png'),
('Housing Development Finance Corporation Limited Bank', 'HDFC', 'http://icon.mobile.matchmove/mobile/hdfc.png'),
('Yes Bank', 'YESBANK', 'http://icon.mobile.matchmove/mobile/yesbank.png');

-----Create users
INSERT INTO public.users
(upi_id, mob_number, first_name, mid_name, last_name, country_id, primary_account_id, pan, pin, create_on, max_transaction_limit)
VALUES('rake@ybl', '9971588951', 'Rakesh', '', 'Kumar', 1, 0, 'DBLPK543R', '675476', now(), 100000000);

INSERT INTO public.users
(upi_id, mob_number, first_name, mid_name, last_name, country_id, primary_account_id, pan, pin, create_on, max_transaction_limit)
VALUES('sudhir@ybl', '9971999900', 'Rakesh', '', 'Kumar', 1, 0, 'DBLPK543B', '885476', now(), 100000000);

----Create users account
INSERT INTO public.users_account
(account_number, ifsc, user_id, bank_id )
VALUES('123456789', 'SBIN0001234', 1,1);

INSERT INTO public.users_account
(account_number, ifsc, user_id, bank_id  )
VALUES('123000009', 'SBIN0001234', 1,1);

INSERT INTO public.users_account
(account_number, ifsc, user_id, bank_id  )
VALUES('122000009', 'HDFC0001234', 2, 2);

INSERT INTO public.users_account
(account_number, ifsc, user_id, bank_id  )
VALUES('12200012', 'YESBANK0001', 2, 3);

----Create transactions
INSERT INTO public."transaction"
(utr, amount, payment_time, payment_gateway, status, from_account_id, to_account_id )
VALUES('UTR1200123456789045DSE', 1000, now(), 'razorpay', 'success',1, 2);

INSERT INTO public."transaction"
(utr, amount, payment_time, payment_gateway, status, from_account_id, to_account_id )
VALUES('UTR1888844446789045DSE', 500, now(), 'razorpay', 'success',1, 2);

INSERT INTO public."transaction"
(utr, amount, payment_time, payment_gateway, status, from_account_id, to_account_id )
VALUES('UTR1668844222678904DSE', 100, now(), 'razorpay', 'success',2, 1);

INSERT INTO public."transaction"
(utr, amount, payment_time, payment_gateway, status, from_account_id, to_account_id )
VALUES('UTR1668844222678904DZZ', 100, now(), 'razorpay', 'success',1, 3);

INSERT INTO public."transaction"
(utr, amount, payment_time, payment_gateway, status, from_account_id, to_account_id )
VALUES('UTR166884422267896YYZZ', 3500, now(), 'razorpay', 'success',2, 3);

INSERT INTO public."transaction"
(utr, amount, payment_time, payment_gateway, status, from_account_id, to_account_id )
VALUES('UTR166884422267896SEND', 1400, now(), 'razorpay', 'success',3, 2);