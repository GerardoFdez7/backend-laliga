#!/bin/sh

until pg_isready -h db -U postgres; do
  echo "Esperando por la base de datos..."
  sleep 2
done

echo "Base de datos lista, arrancando la API..."
#exec go run main.go
exec air -c .air.toml
