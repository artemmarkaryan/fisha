create index recommendations_user_id_shown_rank_index
    on recommendations (user_id, rank)
    where (shown = false);
