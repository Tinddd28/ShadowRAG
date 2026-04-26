CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION IF NOT EXISTS age;

-- AGE требует явного LOAD в каждой сессии — вынесем в shared_preload_libraries
-- или добавим в search_path
LOAD 'age';
SET search_path = ag_catalog, "$user", public;