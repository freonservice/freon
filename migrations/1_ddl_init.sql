-- +goose Up

CREATE SCHEMA IF NOT EXISTS public;

create extension if not exists ltree;

create table if not exists public.users
(
    id          SERIAL PRIMARY KEY,
    email       varchar(50)  not null,
    password    varchar(100) not null,
    first_name  varchar(255),
    second_name varchar(255),
    uuid_id     varchar(36)  not null,
    status      smallint     not null       default 0,
    role        smallint     not null       default 0,
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
    status     smallint     not null       default 0,
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
    id            SERIAL PRIMARY KEY,
    creator_id    integer
        constraint identifiers_users_user_id_fk
            references public.users,
    category_id   integer
        constraint identifiers_categories_category_id_fk
            references public.categories on delete set null default null,
    parent_path   ltree,
    name          varchar(255) not null,
    description   text,
    text_singular text,
    text_plural   text,
    status        smallint     not null                     default 0,
    platforms     varchar(255),
    created_at    timestamp without time zone               default (now() at time zone 'utc'),
    updated_at    timestamp without time zone
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
    status          smallint not null           default 0,
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
    singular        text,
    plural          text,
    status          smallint not null           default 0,
    created_at      timestamp without time zone default (now() at time zone 'utc'),
    updated_at      timestamp without time zone
);

create table if not exists public.translation_files
(
    id              SERIAL PRIMARY KEY,
    localization_id integer
        constraint translation_files_localization_id_fk
            references public.localizations on delete cascade,
    creator_id      integer
        constraint translation_files_users_creator_id_fk
            references public.users,
    names           varchar(500) not null, -- storing list names with comma-separator for android/ios and without for web
    web_folder_path varchar(255) not null, -- folder for documents available from web
    fs_folder_path  varchar(255) not null, -- folder for documents available from fs
    s3_file_id      varchar(100),
    s3_bucket       varchar(100),
    platform        smallint     not null       default 0,
    status          smallint     not null       default 0,
    storage_type    smallint     not null       default 0,
    created_at      timestamp without time zone default (now() at time zone 'utc'),
    updated_at      timestamp without time zone
);

-- create unique index if not exists translation_files_web_folder_path_uindex
--     on public.translation_files (web_folder_path);

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

create table if not exists public.languages
(
    id     SERIAL PRIMARY KEY,
    locale varchar(5)   not null,
    name   varchar(100) not null
);

create unique index if not exists languages_locale_uindex
    on public.languages (locale);

INSERT INTO public.languages (id, name, locale)
VALUES (1, 'Abkhaz', 'ab');
INSERT INTO public.languages (id, name, locale)
VALUES (2, 'Afar', 'aa');
INSERT INTO public.languages (id, name, locale)
VALUES (3, 'Afrikaans', 'af');
INSERT INTO public.languages (id, name, locale)
VALUES (4, 'Akan', 'ak');
INSERT INTO public.languages (id, name, locale)
VALUES (5, 'Albanian', 'sq');
INSERT INTO public.languages (id, name, locale)
VALUES (6, 'Amharic', 'am');
INSERT INTO public.languages (id, name, locale)
VALUES (7, 'Arabic', 'ar');
INSERT INTO public.languages (id, name, locale)
VALUES (8, 'Aragonese', 'an');
INSERT INTO public.languages (id, name, locale)
VALUES (9, 'Armenian', 'hy');
INSERT INTO public.languages (id, name, locale)
VALUES (10, 'Assamese', 'as');
INSERT INTO public.languages (id, name, locale)
VALUES (11, 'Avaric', 'av');
INSERT INTO public.languages (id, name, locale)
VALUES (12, 'Avestan', 'ae');
INSERT INTO public.languages (id, name, locale)
VALUES (13, 'Aymara', 'ay');
INSERT INTO public.languages (id, name, locale)
VALUES (14, 'Azerbaijani', 'az');
INSERT INTO public.languages (id, name, locale)
VALUES (15, 'Bambara', 'bm');
INSERT INTO public.languages (id, name, locale)
VALUES (16, 'Bashkir', 'ba');
INSERT INTO public.languages (id, name, locale)
VALUES (17, 'Basque', 'eu');
INSERT INTO public.languages (id, name, locale)
VALUES (18, 'Belarusian', 'be');
INSERT INTO public.languages (id, name, locale)
VALUES (19, 'Bengali, Bangla', 'bn');
INSERT INTO public.languages (id, name, locale)
VALUES (20, 'Bihari', 'bh');
INSERT INTO public.languages (id, name, locale)
VALUES (21, 'Bislama', 'bi');
INSERT INTO public.languages (id, name, locale)
VALUES (22, 'Bosnian', 'bs');
INSERT INTO public.languages (id, name, locale)
VALUES (23, 'Breton', 'br');
INSERT INTO public.languages (id, name, locale)
VALUES (24, 'Bulgarian', 'bg');
INSERT INTO public.languages (id, name, locale)
VALUES (25, 'Burmese', 'my');
INSERT INTO public.languages (id, name, locale)
VALUES (26, 'Catalan', 'ca');
INSERT INTO public.languages (id, name, locale)
VALUES (27, 'Chamorro', 'ch');
INSERT INTO public.languages (id, name, locale)
VALUES (28, 'Chechen', 'ce');
INSERT INTO public.languages (id, name, locale)
VALUES (29, 'Chichewa, Chewa, Nyanja', 'ny');
INSERT INTO public.languages (id, name, locale)
VALUES (30, 'Chinese', 'zh');
INSERT INTO public.languages (id, name, locale)
VALUES (31, 'Chuvash', 'cv');
INSERT INTO public.languages (id, name, locale)
VALUES (32, 'Cornish', 'kw');
INSERT INTO public.languages (id, name, locale)
VALUES (33, 'Corsican', 'co');
INSERT INTO public.languages (id, name, locale)
VALUES (34, 'Cree', 'cr');
INSERT INTO public.languages (id, name, locale)
VALUES (35, 'Croatian', 'hr');
INSERT INTO public.languages (id, name, locale)
VALUES (36, 'Czech', 'cs');
INSERT INTO public.languages (id, name, locale)
VALUES (37, 'Danish', 'da');
INSERT INTO public.languages (id, name, locale)
VALUES (38, 'Divehi, Dhivehi, Maldivian', 'dv');
INSERT INTO public.languages (id, name, locale)
VALUES (39, 'Dutch', 'nl');
INSERT INTO public.languages (id, name, locale)
VALUES (40, 'Dzongkha', 'dz');
INSERT INTO public.languages (id, name, locale)
VALUES (41, 'English', 'en');
INSERT INTO public.languages (id, name, locale)
VALUES (42, 'Esperanto', 'eo');
INSERT INTO public.languages (id, name, locale)
VALUES (43, 'Estonian', 'et');
INSERT INTO public.languages (id, name, locale)
VALUES (44, 'Ewe', 'ee');
INSERT INTO public.languages (id, name, locale)
VALUES (45, 'Faroese', 'fo');
INSERT INTO public.languages (id, name, locale)
VALUES (46, 'Fijian', 'fj');
INSERT INTO public.languages (id, name, locale)
VALUES (47, 'Finnish', 'fi');
INSERT INTO public.languages (id, name, locale)
VALUES (48, 'French', 'fr');
INSERT INTO public.languages (id, name, locale)
VALUES (49, 'Fula, Fulah, Pulaar, Pular', 'ff');
INSERT INTO public.languages (id, name, locale)
VALUES (50, 'Galician', 'gl');
INSERT INTO public.languages (id, name, locale)
VALUES (51, 'Georgian', 'ka');
INSERT INTO public.languages (id, name, locale)
VALUES (52, 'German', 'de');
INSERT INTO public.languages (id, name, locale)
VALUES (53, 'Greek (modern)', 'el');
INSERT INTO public.languages (id, name, locale)
VALUES (54, 'Guaraní', 'gn');
INSERT INTO public.languages (id, name, locale)
VALUES (55, 'Gujarati', 'gu');
INSERT INTO public.languages (id, name, locale)
VALUES (56, 'Haitian, Haitian Creole', 'ht');
INSERT INTO public.languages (id, name, locale)
VALUES (57, 'Hausa', 'ha');
INSERT INTO public.languages (id, name, locale)
VALUES (58, 'Hebrew (modern)', 'he');
INSERT INTO public.languages (id, name, locale)
VALUES (59, 'Herero', 'hz');
INSERT INTO public.languages (id, name, locale)
VALUES (60, 'Hindi', 'hi');
INSERT INTO public.languages (id, name, locale)
VALUES (61, 'Hiri Motu', 'ho');
INSERT INTO public.languages (id, name, locale)
VALUES (62, 'Hungarian', 'hu');
INSERT INTO public.languages (id, name, locale)
VALUES (63, 'Interlingua', 'ia');
INSERT INTO public.languages (id, name, locale)
VALUES (64, 'Indonesian', 'id');
INSERT INTO public.languages (id, name, locale)
VALUES (65, 'Interlingue', 'ie');
INSERT INTO public.languages (id, name, locale)
VALUES (66, 'Irish', 'ga');
INSERT INTO public.languages (id, name, locale)
VALUES (67, 'Igbo', 'ig');
INSERT INTO public.languages (id, name, locale)
VALUES (68, 'Inupiaq', 'ik');
INSERT INTO public.languages (id, name, locale)
VALUES (69, 'Ido', 'io');
INSERT INTO public.languages (id, name, locale)
VALUES (70, 'Icelandic', 'is');
INSERT INTO public.languages (id, name, locale)
VALUES (71, 'Italian', 'it');
INSERT INTO public.languages (id, name, locale)
VALUES (72, 'Inuktitut', 'iu');
INSERT INTO public.languages (id, name, locale)
VALUES (73, 'Japanese', 'ja');
INSERT INTO public.languages (id, name, locale)
VALUES (74, 'Javanese', 'jv');
INSERT INTO public.languages (id, name, locale)
VALUES (75, 'Kalaallisut, Greenlandic', 'kl');
INSERT INTO public.languages (id, name, locale)
VALUES (76, 'Kannada', 'kn');
INSERT INTO public.languages (id, name, locale)
VALUES (77, 'Kanuri', 'kr');
INSERT INTO public.languages (id, name, locale)
VALUES (78, 'Kashmiri', 'ks');
INSERT INTO public.languages (id, name, locale)
VALUES (79, 'Kazakh', 'kk');
INSERT INTO public.languages (id, name, locale)
VALUES (80, 'Khmer', 'km');
INSERT INTO public.languages (id, name, locale)
VALUES (81, 'Kikuyu, Gikuyu', 'ki');
INSERT INTO public.languages (id, name, locale)
VALUES (82, 'Kinyarwanda', 'rw');
INSERT INTO public.languages (id, name, locale)
VALUES (83, 'Kyrgyz', 'ky');
INSERT INTO public.languages (id, name, locale)
VALUES (84, 'Komi', 'kv');
INSERT INTO public.languages (id, name, locale)
VALUES (85, 'Kongo', 'kg');
INSERT INTO public.languages (id, name, locale)
VALUES (86, 'Korean', 'ko');
INSERT INTO public.languages (id, name, locale)
VALUES (87, 'Kurdish', 'ku');
INSERT INTO public.languages (id, name, locale)
VALUES (88, 'Kwanyama, Kuanyama', 'kj');
INSERT INTO public.languages (id, name, locale)
VALUES (89, 'Latin', 'la');
INSERT INTO public.languages (id, name, locale)
VALUES (90, 'Luxembourgish, Letzeburgesch', 'lb');
INSERT INTO public.languages (id, name, locale)
VALUES (91, 'Ganda', 'lg');
INSERT INTO public.languages (id, name, locale)
VALUES (92, 'Limburgish, Limburgan, Limburger', 'li');
INSERT INTO public.languages (id, name, locale)
VALUES (93, 'Lingala', 'ln');
INSERT INTO public.languages (id, name, locale)
VALUES (94, 'Lao', 'lo');
INSERT INTO public.languages (id, name, locale)
VALUES (95, 'Lithuanian', 'lt');
INSERT INTO public.languages (id, name, locale)
VALUES (96, 'Luba-Katanga', 'lu');
INSERT INTO public.languages (id, name, locale)
VALUES (97, 'Latvian', 'lv');
INSERT INTO public.languages (id, name, locale)
VALUES (98, 'Manx', 'gv');
INSERT INTO public.languages (id, name, locale)
VALUES (99, 'Macedonian', 'mk');
INSERT INTO public.languages (id, name, locale)
VALUES (100, 'Malagasy', 'mg');
INSERT INTO public.languages (id, name, locale)
VALUES (101, 'Malay', 'ms');
INSERT INTO public.languages (id, name, locale)
VALUES (102, 'Malayalam', 'ml');
INSERT INTO public.languages (id, name, locale)
VALUES (103, 'Maltese', 'mt');
INSERT INTO public.languages (id, name, locale)
VALUES (104, 'M?ori', 'mi');
INSERT INTO public.languages (id, name, locale)
VALUES (105, 'Marathi', 'mr');
INSERT INTO public.languages (id, name, locale)
VALUES (106, 'Marshallese', 'mh');
INSERT INTO public.languages (id, name, locale)
VALUES (107, 'Mongolian', 'mn');
INSERT INTO public.languages (id, name, locale)
VALUES (108, 'Nauruan', 'na');
INSERT INTO public.languages (id, name, locale)
VALUES (109, 'Navajo, Navaho', 'nv');
INSERT INTO public.languages (id, name, locale)
VALUES (110, 'Northern Ndebele', 'nd');
INSERT INTO public.languages (id, name, locale)
VALUES (111, 'Nepali', 'ne');
INSERT INTO public.languages (id, name, locale)
VALUES (112, 'Ndonga', 'ng');
INSERT INTO public.languages (id, name, locale)
VALUES (113, 'Norwegian Bokmål', 'nb');
INSERT INTO public.languages (id, name, locale)
VALUES (114, 'Norwegian Nynorsk', 'nn');
INSERT INTO public.languages (id, name, locale)
VALUES (115, 'Norwegian', 'no');
INSERT INTO public.languages (id, name, locale)
VALUES (116, 'Nuosu', 'ii');
INSERT INTO public.languages (id, name, locale)
VALUES (117, 'Southern Ndebele', 'nr');
INSERT INTO public.languages (id, name, locale)
VALUES (118, 'Occitan', 'oc');
INSERT INTO public.languages (id, name, locale)
VALUES (119, 'Ojibwe, Ojibwa', 'oj');
INSERT INTO public.languages (id, name, locale)
VALUES (120, 'Old Church Slavonic, Church Slavonic, Old Bulgarian', 'cu');
INSERT INTO public.languages (id, name, locale)
VALUES (121, 'Oromo', 'om');
INSERT INTO public.languages (id, name, locale)
VALUES (122, 'Oriya', 'or');
INSERT INTO public.languages (id, name, locale)
VALUES (123, 'Ossetian, Ossetic', 'os');
INSERT INTO public.languages (id, name, locale)
VALUES (124, 'Panjabi, Punjabi', 'pa');
INSERT INTO public.languages (id, name, locale)
VALUES (125, 'P?li', 'pi');
INSERT INTO public.languages (id, name, locale)
VALUES (126, 'Persian (Farsi)', 'fa');
INSERT INTO public.languages (id, name, locale)
VALUES (127, 'Polish', 'pl');
INSERT INTO public.languages (id, name, locale)
VALUES (128, 'Pashto, Pushto', 'ps');
INSERT INTO public.languages (id, name, locale)
VALUES (129, 'Portuguese', 'pt');
INSERT INTO public.languages (id, name, locale)
VALUES (130, 'Quechua', 'qu');
INSERT INTO public.languages (id, name, locale)
VALUES (131, 'Romansh', 'rm');
INSERT INTO public.languages (id, name, locale)
VALUES (132, 'Kirundi', 'rn');
INSERT INTO public.languages (id, name, locale)
VALUES (133, 'Reunionese, Reunion Creole', 'rc');
INSERT INTO public.languages (id, name, locale)
VALUES (134, 'Romanian', 'ro');
INSERT INTO public.languages (id, name, locale)
VALUES (135, 'Russian', 'ru');
INSERT INTO public.languages (id, name, locale)
VALUES (136, 'Sanskrit', 'sa');
INSERT INTO public.languages (id, name, locale)
VALUES (137, 'Sardinian', 'sc');
INSERT INTO public.languages (id, name, locale)
VALUES (138, 'Sindhi', 'sd');
INSERT INTO public.languages (id, name, locale)
VALUES (139, 'Northern Sami', 'se');
INSERT INTO public.languages (id, name, locale)
VALUES (140, 'Samoan', 'sm');
INSERT INTO public.languages (id, name, locale)
VALUES (141, 'Sango', 'sg');
INSERT INTO public.languages (id, name, locale)
VALUES (142, 'Serbian', 'sr');
INSERT INTO public.languages (id, name, locale)
VALUES (143, 'Scottish Gaelic, Gaelic', 'gd');
INSERT INTO public.languages (id, name, locale)
VALUES (144, 'Shona', 'sn');
INSERT INTO public.languages (id, name, locale)
VALUES (145, 'Sinhalese, Sinhala', 'si');
INSERT INTO public.languages (id, name, locale)
VALUES (146, 'Slovak', 'sk');
INSERT INTO public.languages (id, name, locale)
VALUES (147, 'Slovene', 'sl');
INSERT INTO public.languages (id, name, locale)
VALUES (148, 'Somali', 'so');
INSERT INTO public.languages (id, name, locale)
VALUES (149, 'Southern Sotho', 'st');
INSERT INTO public.languages (id, name, locale)
VALUES (150, 'Spanish', 'es');
INSERT INTO public.languages (id, name, locale)
VALUES (151, 'Sundanese', 'su');
INSERT INTO public.languages (id, name, locale)
VALUES (152, 'Swahili', 'sw');
INSERT INTO public.languages (id, name, locale)
VALUES (153, 'Swati', 'ss');
INSERT INTO public.languages (id, name, locale)
VALUES (154, 'Swedish', 'sv');
INSERT INTO public.languages (id, name, locale)
VALUES (155, 'Tamil', 'ta');
INSERT INTO public.languages (id, name, locale)
VALUES (156, 'Telugu', 'te');
INSERT INTO public.languages (id, name, locale)
VALUES (157, 'Tajik', 'tg');
INSERT INTO public.languages (id, name, locale)
VALUES (158, 'Thai', 'th');
INSERT INTO public.languages (id, name, locale)
VALUES (159, 'Tigrinya', 'ti');
INSERT INTO public.languages (id, name, locale)
VALUES (160, 'Tibetan Standard, Tibetan, Central', 'bo');
INSERT INTO public.languages (id, name, locale)
VALUES (161, 'Turkmen', 'tk');
INSERT INTO public.languages (id, name, locale)
VALUES (162, 'Tagalog', 'tl');
INSERT INTO public.languages (id, name, locale)
VALUES (163, 'Tswana', 'tn');
INSERT INTO public.languages (id, name, locale)
VALUES (164, 'Tonga (Tonga Islands)', 'to');
INSERT INTO public.languages (id, name, locale)
VALUES (165, 'Turkish', 'tr');
INSERT INTO public.languages (id, name, locale)
VALUES (166, 'Tsonga', 'ts');
INSERT INTO public.languages (id, name, locale)
VALUES (167, 'Tatar', 'tt');
INSERT INTO public.languages (id, name, locale)
VALUES (168, 'Twi', 'tw');
INSERT INTO public.languages (id, name, locale)
VALUES (169, 'Tahitian', 'ty');
INSERT INTO public.languages (id, name, locale)
VALUES (170, 'Uyghur', 'ug');
INSERT INTO public.languages (id, name, locale)
VALUES (171, 'Ukrainian', 'uk');
INSERT INTO public.languages (id, name, locale)
VALUES (172, 'Urdu', 'ur');
INSERT INTO public.languages (id, name, locale)
VALUES (173, 'Uzbek', 'uz');
INSERT INTO public.languages (id, name, locale)
VALUES (174, 'Venda', 've');
INSERT INTO public.languages (id, name, locale)
VALUES (175, 'Vietlang_namese', 'vi');
INSERT INTO public.languages (id, name, locale)
VALUES (176, 'Volapük', 'vo');
INSERT INTO public.languages (id, name, locale)
VALUES (177, 'Walloon', 'wa');
INSERT INTO public.languages (id, name, locale)
VALUES (178, 'Welsh', 'cy');
INSERT INTO public.languages (id, name, locale)
VALUES (179, 'Wolof', 'wo');
INSERT INTO public.languages (id, name, locale)
VALUES (180, 'Western Frisian', 'fy');
INSERT INTO public.languages (id, name, locale)
VALUES (181, 'Xhosa', 'xh');
INSERT INTO public.languages (id, name, locale)
VALUES (182, 'Yiddish', 'yi');
INSERT INTO public.languages (id, name, locale)
VALUES (183, 'Yoruba', 'yo');
INSERT INTO public.languages (id, name, locale)
VALUES (184, 'Zhuang, Chuang', 'za');
INSERT INTO public.languages (id, name, locale)
VALUES (185, 'Zulu', 'zu');

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

DROP TABLE translation_files;
DROP INDEX translation_files_path_uindex;

DROP TABLE languages;
DROP INDEX languages_locale_uindex;

DROP SCHEMA public;
