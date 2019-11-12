# docker 脚本

# 启动mysql
docker run --name mysql -v ~/Documents/docker/mysql:/var/lib/mysql  -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7