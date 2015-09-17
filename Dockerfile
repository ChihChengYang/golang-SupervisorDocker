# sudo docker build -t="centos/ms" .
# Base image  
FROM centos:centos7
MAINTAINER Jeff Yang

RUN yum update -y
RUN yum -y install python-setuptools python-pip
RUN easy_install supervisor
 
ENV MAJOR_PATH /home/apps/ 
RUN mkdir -p $MAJOR_PATH
 
COPY MS $MAJOR_PATH 

RUN mkdir -p /var/log/supervisor 
COPY supervisord.conf $MAJOR_PATH

COPY start.sh $MAJOR_PATH

RUN ln -s $MAJOR_PATH/supervisord.conf /etc/supervisord.conf
 
EXPOSE  9080 

CMD  $MAJOR_PATH/start.sh 
