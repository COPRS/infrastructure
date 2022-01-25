CREATE USER {{ keycloak.database.user }} WITH ENCRYPTED PASSWORD '{{ keycloak.database.password }}';
CREATE DATABASE {{ keycloak.database.name }};
GRANT ALL PRIVILEGES ON DATABASE {{ keycloak.database.name }} TO {{ keycloak.database.user }};
CREATE USER scdf WITH ENCRYPTED PASSWORD 'scdfpassword';
CREATE DATABASE skipper;
GRANT ALL PRIVILEGES ON DATABASE skipper TO scdf;
CREATE DATABASE dataflow;
GRANT ALL PRIVILEGES ON DATABASE dataflow TO scdf;
