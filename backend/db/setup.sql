-- CREATE DATABASE nclouds_db;
/**/;
CREATE USER backenduser WITH PASSWORD 'backendpassword';
GRANT ALL PRIVILEGES ON DATABASE nclouds_db TO backenduser;
GRANT ALL PRIVILEGES ON TABLE users TO backenduser;