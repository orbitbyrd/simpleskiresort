
CREATE TABLE public.hills(
    id integer,
    length DOUBLE PRECISION,
    slope DOUBLE PRECISION
);

ALTER TABLE public.hills
    OWNER to postgres;

INSERT INTO hills values (1, 99.88, 0.22);
INSERT INTO hills values (2, 88.99, 0.33);