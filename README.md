# mysql-using-goLang-on-Docker-container
mysql dbconectivity used in goLang to run db operations in Docker container using Docker file for basic CRUD operations

Reference
https://github.com/go-sql-driver/mysql

Pre-requisite
Sql Client should already be installed in machine.
apt-get update && apt-get install -y mysql-client

Introduction
1- First thing first, we need mysql running container in order to link it with our container. So we need to initiate a 
mysql container. Container runs in detached mode because we need it as a service.
docker run --detach --name=<containerName> --env="MYSQL_ROOT_PASSWORD=<mypassword>" mysql

To see logs of the running container use
docker logs <containerName>

2- Install sql client locally
apt-get install mysql-client

3- Test the current IP Address of mysql
docker inspect <containerName>

4- Using obtained IP Address type
mysql -uroot -p<mypassword> -h <IPAdress> -P 3306
By default port 3306 is assigned to mysql, it can be changed if needed

5- Run our docker container linking to this mysql container
docker run -dit --name <ourNewContainer> --link <containerNameOfSql>:mysql <ImageName>
This linking should be checked by executing in our container in Interactive mode and open "cat /etc/hosts" where you
should see mySql linking to our container with IP Address

6- Access mySql client this way
mysql -uroot -p<mypassword> -h <mySQLIPAddress> -P <mySQLIPPort> --database=<DatabaseName>

