CREATE USER keycloak WITH ENCRYPTED PASSWORD 'keycloakpassword';
CREATE DATABASE keycloak;
GRANT ALL PRIVILEGES ON DATABASE keycloak TO keycloak;
CREATE USER scdf WITH ENCRYPTED PASSWORD 'scdfpassword';
CREATE DATABASE skipper;
GRANT ALL PRIVILEGES ON DATABASE skipper TO scdf;
CREATE DATABASE dataflow;
GRANT ALL PRIVILEGES ON DATABASE dataflow TO scdf;
