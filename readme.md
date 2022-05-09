## Redigo Example
Basic redis in REST-API with redigo & echo framework

### Setup Project
1. Clone this project
    ```
    git clone https://github.com/ellashella24/redigo-example.git redigo-example
    ```
2. Create config file based on `config-example.toml` 
3. Create Postgre Database for apps
4. Create local redis service
    ```
    docker run --name redis -p 6379:6379 -d redis
    ```
5. If you not use the seeder, add comment on line 24 to 29 in file `postgre.go` at driver directory

### Run Project
1. Run this project 
    ```go
    go run main.go
    ```
2. To get response time from endpoint get all data using redis
    ```
    curl -o /dev/null -w "\n%{time_total} seconds\n" http://localhost:8000/withredis
    ```
3. To get response time from endpoint get all data without redis
    ```
    curl -o /dev/null -w "\n%{time_total} seconds\n" http://localhost:8000/withoutredis
    ```

### Acknowledgment
1. Redha Juanda on Medium - [Implementasi Server Side Caching dengan Redis](https://medium.com/redhajuanda/implementasi-server-side-caching-dengan-redis-part-iii-3f45a5c30bd5)
2. pete911 on Github - [example redigo github repository](https://github.com/pete911/examples-redigo)
3. furqonzt99 on Github - [news-redis github repository](https://github.com/furqonzt99/news-redis)
4. Redigo on go package - [redigo package documentation](https://pkg.go.dev/github.com/gomodule/redigo@v1.8.8/redis) 
5. Redis commands - [redis commands documentation](https://redis.io/commands/)