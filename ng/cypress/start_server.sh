docker compose down
docker volume rm delivrio_db_data
# Specify project name so the volume is deterministically named
docker compose -p delivrio up --detach
docker compose logs
