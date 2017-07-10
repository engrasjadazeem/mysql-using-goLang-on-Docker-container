#Read https://severalnines.com/blog/mysql-docker-containers-understanding-basics
########################################################################################################################
#Running an instance of mysql on terminal
#$ docker run --detach --name=test-mysql --env="MYSQL_ROOT_PASSWORD=<mypassword>" mysql
#To see logs of container
#$ docker logs test-mysql
#Install sql client
#$ apt-get install mysql-client
#Get your IP Address
#$ docker inspect test-mysql
#Connect to mysql
#$ mysql -uroot -p<mypassword> -h <IPAdress(172.17.0.02)> -P 3306
#Run new container linking to mysql
#docker run -dit --name test-link --link test-mysql:mysql linkcontainers
#Run Testdb
#mysql -uroot -p$mySQLPassword -h $mySQLIPAddress -P $mySQLIPPort --database=testdb

#Create Table in mysql

FROM golang:1.8.3
LABEL maintainer "Asjad"

RUN apt-get update && apt-get install -y mysql-client
EXPOSE 8080


ENV mySQLIPAddress 172.17.0.2
ENV mySQLIPPort 3306
ENV mySQLPassword mypassword

RUN echo $mySQLIPAddress
RUN echo $mySQLIPPort
RUN echo $mySQLPassword

RUN go get github.com/go-sql-driver/mysql
RUN mkdir /app
ADD . /app/
WORKDIR /app

#Go Execution
RUN go build -o main .
CMD ["/app/main"]