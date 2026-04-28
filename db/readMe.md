```bash

# 安装migrate  golang-migrate/migrate

# https://github.com/golang-migrate/migrate/tree/master/cmd/migrate



# 创建数据表
migrate create -ext sql -dir db/migrations -seq create_users_table
# 生成数据表
migrate -database "mysql://root:123456@tcp(localhost:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true" -path db/migrations up
# 删除数据表 **会删除整个数据库**
migrate -database "mysql://root:123456@tcp(localhost:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true" -path db/migrations drop
# 删除数据表 **会删除整个数据库**
migrate -database "mysql://root:123456@tcp(localhost:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true" -path db/migrations down
# 查看当前版本
migrate -database "mysql://root:123456@tcp(localhost:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true" -path db/migrations version
# 回退到上一个版本
migrate -database "mysql://root:123456@tcp(localhost:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true" -path db/migrations down 1







```