/* version: "3.9"
services:
  postgres:
    image: postgres:latest
    environment:
      POSTGRES_DB: "tododb"
      POSTGRES_USER: "todoadmin"
      POSTGRES_PASSWORD: "tododo"
       PGDATA: "/var/lib/postgresql/data/pgdata"
    volumes:
      - ../db:/docker-entrypoint-initdb.d
      - .:/var/lib/postgresql/data
    
    ports:
      - "5432:5432" */