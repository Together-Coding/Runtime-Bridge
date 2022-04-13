# runtime-bridge

# Requirements

> To be updated

# DB Migration - [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#with-go-toolchain)

1. Install golang-migrate CLI
    ```shell
    $ (root) curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
    $ (root) echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
    $ apt-get update
    $ apt-get install -y migrate
    ```
2. When you want to modify DB schema, create migration scripts
    ```shell
    $ migrate create -ext sql -dir db/migrations -seq <title>
    ```
3. Write DDL at the created up/down.sql files
4. Run a migration 
   ```shell
   $ migrate -verbose -database "mysql://<user>:<url_encoded_password>@tcp(<host>:<port>>)/<db_name>" -path db/migrations up  # or down
   ```
