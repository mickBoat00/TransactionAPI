-- +goose Up
INSERT INTO currencies (id, name, code, createdAt, updatedAt) 
VALUES 
('16a74e66-0285-4cb8-a8f4-dc642c855b56', 'US Dollar', 'USD', '2024-01-20 15:39:01.621648236', '2024-01-20 15:39:01.621648236'),
('8adfa6ab-b0f4-4a56-8dbe-20df3dd69caf', 'Euro', 'EUR', '2024-01-20 15:39:01.621648236', '2024-01-20 15:39:01.621648236'),
('fbb3902b-0dad-49db-9f25-c48ef99acfa6', 'Great British Pound', 'GBP', '2024-01-20 15:39:01.621648236', '2024-01-20 15:39:01.621648236'),
('48c30d44-86d1-4765-b9b7-fca181a280a6', 'Ghanaian Cedis', 'GHS', '2024-01-20 15:39:01.621648236', '2024-01-20 15:39:01.621648236'),
('3d39310c-cfc6-4931-a425-a1b0a036bf9f', 'Japanesse Yen', 'Yen', '2024-01-20 15:39:01.621648236', '2024-01-20 15:39:01.621648236'),
('a99c1aab-3dba-4318-a910-f57118d1ae5a', 'South African Rand', 'ZAR', '2024-01-20 15:39:01.621648236', '2024-01-20 15:39:01.621648236');

-- +goose Down
DELETE FROM currencies;