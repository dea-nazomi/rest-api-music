-- Table: public.music_list

-- DROP TABLE public.music_list;

CREATE TABLE IF NOT EXISTS public.music_list
(
    music_id character varying(36) COLLATE pg_catalog."default" NOT NULL,
    music_name character(50) COLLATE pg_catalog."default",
    music_album character(50) COLLATE pg_catalog."default",
    music_album_art character(50) COLLATE pg_catalog."default",
    music_singer character(50) COLLATE pg_catalog."default",
    music_publish_date date,
    music_created_at timestamp without time zone DEFAULT now(),
    music_updated_at timestamp without time zone DEFAULT now(),
    CONSTRAINT music_list_pkey PRIMARY KEY (music_id)
)

TABLESPACE pg_default;

ALTER TABLE public.music_list
    OWNER to postgres;