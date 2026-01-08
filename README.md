# Container for seeding a postgres container

This container will connect to the database configured and execute an SQL file mounted to `/data/seed.sql`

## Run the container
Using docker-compose:
```yaml
services:
  seed_postgres_db:
    container_name: seed-postgres-db
    image: ghcr.io/LightJack05/seed-postgres-db:latest
    # Change these environment variables to match your database configuration
    environment:
      - DATABASE_CONFIG_DB_HOST=your-db
      - DATABASE_CONFIG_DB_PORT=5432
      - DATABASE_CONFIG_DB_USER=seed_postgres_db
      - DATABASE_CONFIG_DB_PASS=4dfef616-220c-4e85-9991-485f9af486e3
      - DATABASE_CONFIG_DB_NAME=seed_postgres_db
      - DATABASE_CONFIG_DISABLE_SSL=true
    # Mount your SQL file here
    # This will be run on the database when the container starts
    volumes:
      - ./seed.sql:/data/seed.sql
    # Make sure the database is healthy before running this container
    # If your application runs migrations, you will have to wait for that to finish first
    depends_on:
      your-db:
        condition: service_healthy
```


