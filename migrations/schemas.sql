
CREATE TABLE public."transactions" (
  "id" SERIAL PRIMARY KEY,
  "datetime" TIMESTAMPTZ,
  "amount" NUMBER UNIQUE
);
