CREATE TABLE public."transactions" (
  "id" SERIAL PRIMARY KEY,
  "datetime" TIMESTAMPTZ,
  "amount" double precision
);
