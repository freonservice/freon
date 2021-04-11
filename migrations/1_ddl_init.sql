-- +goose Up

CREATE SCHEMA IF NOT EXISTS public;

create extension if not exists ltree;

create table if not exists public.users
(
    id          SERIAL PRIMARY KEY,
    email       varchar(25)  not null,
    password    varchar(100) not null,
    first_name  varchar(255),
    second_name varchar(255),
    uuid_id     varchar(36)  not null,
    status      integer      not null       default 0,
    role        integer      not null       default 0,
    created_at  timestamp without time zone default (now() at time zone 'utc'),
    updated_at  timestamp without time zone
);

create unique index if not exists users_email_uindex
    on public.users (email);

create unique index if not exists users_uuid_id_uindex
    on public.users (uuid_id);

create table if not exists public.localizations
(
    id         SERIAL PRIMARY KEY,
    creator_id integer
        constraint localizations_users_user_id_fk
            references public.users,
    locale     varchar(10)  not null,
    lang_name  varchar(100) not null,
    status     integer      not null       default 0,
    icon       text,
    created_at timestamp without time zone default (now() at time zone 'utc'),
    updated_at timestamp without time zone
);

create unique index if not exists localizations_locale_uindex
    on public.localizations (locale);

create table if not exists public.categories
(
    id         SERIAL PRIMARY KEY,
    name       varchar(255) not null,
    created_at timestamp without time zone default (now() at time zone 'utc'),
    updated_at timestamp without time zone
);

create unique index if not exists categories_name_uindex
    on public.categories (name);

create table if not exists public.identifiers
(
    id           SERIAL PRIMARY KEY,
    creator_id   integer
        constraint identifiers_users_user_id_fk
            references public.users,
    category_id  integer
        constraint identifiers_categories_category_id_fk
            references public.categories on delete set null default null,
    parent_path  ltree,
    name         varchar(255) not null,
    description  text,
    example_text text,
    status       integer      not null                      default 0,
    platforms    varchar(255),
    named_list   text,
    created_at   timestamp without time zone                default (now() at time zone 'utc'),
    updated_at   timestamp without time zone
);

create unique index if not exists identifiers_name_uindex
    on public.identifiers (name);

create table if not exists public.localization_identifiers
(
    id              SERIAL PRIMARY KEY,
    localization_id integer
        constraint localization_identifiers_localization_id_fk
            references public.localizations,
    identifier_id   integer
        constraint localization_identifiers_identifier_id_fk
            references public.identifiers on delete cascade,
    status          integer not null            default 0,
    created_at      timestamp without time zone default (now() at time zone 'utc'),
    updated_at      timestamp without time zone
);

create table if not exists public.translations
(
    id              SERIAL PRIMARY KEY,
    localization_id integer
        constraint translations_localizations_localization_id_fk
            references public.localizations,
    identifier_id   integer
        constraint translations_identifiers_identifier_id_fk
            references public.identifiers on delete cascade,
    creator_id      integer
        constraint translations_users_creator_id_fk
            references public.users,
    text            text    not null,
    status          integer not null            default 0,
    created_at      timestamp without time zone default (now() at time zone 'utc'),
    updated_at      timestamp without time zone
);

create table if not exists public.user_sessions
(
    id         SERIAL PRIMARY KEY,
    user_id    integer
        constraint user_session_users_user_id_fk
            references public.users on delete cascade,
    token      text                                     not null,
    created_at timestamp without time zone default (now() at time zone 'utc'),
    active     bool                        default true not null
);

create unique index if not exists user_sessions_token_uindex
    on public.user_sessions (token);


-- +goose Down

DROP TABLE translations;

DROP TABLE categories;
DROP INDEX categories_name_uindex;

DROP TABLE identifiers;
DROP INDEX identifiers_name_uindex;

DROP TABLE localizations;
DROP INDEX localizations_locale_uindex;

DROP TABLE localization_identifiers;

DROP TABLE user_sessions;
DROP INDEX user_sessions_token_uindex;

DROP TABLE users;
DROP INDEX users_email_uindex;
DROP INDEX users_uuid_id_uindex;

DROP SCHEMA public;
