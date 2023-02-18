# go-im

## config

### mysql

解决`mysql`不支持datatime为0的情况 [链接](https://blog.csdn.net/sinat_40770656/article/details/101198274)

修改docker宿主机中`mysql`配置文件 `/etc/mysql/mysql.conf.d/mysqld.cnf`
```
[mysqld]
sql_mode=NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES
```