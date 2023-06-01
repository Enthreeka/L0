INSERT INTO "order" (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES ('order1', 'track1', 'entry1', 'en', 'signature1', 'customer1', 'service1', 'shard1', 1, CURRENT_DATE, 'oof1');
-- Добавьте еще тестовые данные по аналогии
INSERT INTO delivery (order_uid, phone, name, zip, city, address, region, email)
VALUES ('order1', '1234567890', 'John Doe', '12345', 'City', 'Address', 'Region', 'john@example.com');
-- Добавьте еще тестовые данные по аналогии
    INSERT INTO payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
VALUES ('order1', 'transaction1', 'request1', 'USD', 'provider1', 100, 1622515200, 'bank1', 10, 90, 5);
-- Добавьте еще тестовые данные по аналогии
INSERT INTO item (chrt_id, track_number, order_uid, price, rid, name, sale, size, total_price, nm_id, brand, status)
VALUES (12222, 'track1', 'order1', 50, 'rid1', 'Product A', 10, 'M', 500, 1, 'Brand X', 1);
-- Добавьте еще тестовые данные по аналогии
