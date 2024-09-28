CREATE TABLE videos
(
    video_id                             UUID PRIMARY KEY,
    v_pub_datetime                       TIMESTAMP,
    v_total_comments                     INTEGER,
    v_year_views                         INTEGER,
    v_month_views                        INTEGER,
    v_week_views                         INTEGER,
    v_day_views                          INTEGER,
    v_likes                              INTEGER,
    v_dislikes                           INTEGER,
    v_duration                           FLOAT,
    v_cr_click_like_7_days               FLOAT,
    v_cr_click_dislike_7_days            FLOAT,
    v_cr_click_long_view_7_days          FLOAT,
    v_cr_click_comment_7_days            FLOAT,
    v_is_hidden                          BOOLEAN,
    v_is_deleted                         BOOLEAN,
    v_avg_watchtime_1_day                FLOAT,
    v_avg_watchtime_7_day                FLOAT,
    v_avg_watchtime_30_day               FLOAT,
    v_frac_avg_watchtime_7_day_duration  FLOAT,
    v_category_popularity_percent_7_days FLOAT,
    v_long_views_7_days                  INTEGER,
    title                                TEXT,
    description                          TEXT,
    category_id                          UUID,
    author_id                            UUID
);
