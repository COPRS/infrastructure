-- Copyright 2023 CS Group
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--     http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.

CREATE USER keycloak WITH ENCRYPTED PASSWORD '{{ keycloak.database.password }}';
CREATE DATABASE keycloak;
GRANT ALL PRIVILEGES ON DATABASE keycloak TO keycloak;
CREATE USER scdf WITH ENCRYPTED PASSWORD '{{ scdf.database.password }}';
CREATE DATABASE skipper;
GRANT ALL PRIVILEGES ON DATABASE skipper TO scdf;
CREATE DATABASE dataflow;
GRANT ALL PRIVILEGES ON DATABASE dataflow TO scdf;
