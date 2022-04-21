alter table user_interest add constraint user_interest__rank_boundaries check ( rank >= -1 and rank <= 1 );
