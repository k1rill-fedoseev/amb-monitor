DROP USER read_user;
REVOKE CONNECT ON DATABASE db FROM readonly;
REVOKE USAGE ON SCHEMA public FROM readonly;
REVOKE SELECT ON ALL TABLES IN SCHEMA public FROM readonly;
DROP ROLE readonly;
