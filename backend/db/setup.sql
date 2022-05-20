-- CREATE DATABASE nclouds_db;
/**/;
CREATE USER backenduser WITH PASSWORD 'backendpassword';
GRANT ALL ON DATABASE nclouds_db TO backenduser;