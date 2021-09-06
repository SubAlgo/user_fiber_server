docker run -d \
  --name=userFiber \
--restart=always -p 9000:5432 \
-e POSTGRES_USER=postgres \
-e POSTGRES_PASSWORD=test123456 \
-e POSTGRES_DB=userFiber \
postgres:13.1
