# select image
FROM mysql

# 環境変数
ENV MYSQL_ROOT_PASSWORD root
ENV MYSQL_DATABASE items

# 初期実行スクリプト置き場に配置
COPY backend-api/init/init.sql /docker-entrypoint-initdb.d/init.sql

# create mysql container image and run image
#$ docker build -t itemdb_mysql .
#$ docker run -p 3306:3306 --name=itemdb_mysql -d itemdb_mysql
# 接続確認
#$ mysql