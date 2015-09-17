 
A demo program for creating a multi-process controller with Supervisor & Docker

================

 
Usage
================

### Build &  Demo

go build MS.go

##First, 
Create a master & slave program. The master part will listen port 9080.
The slave part will hold port 9081,9082,...,and so on depend on Supervisor's config.

##Second, 
Run supervisord will evoke a master program and slave programs be stopped.

##Finally, 
Through "http://127.0.0.1:9080/?start=yes&num=3" the master program will notify all 3 slave programs to start.
"http://127.0.0.1:9080/?close=yes&num=3" to stop.

#1.
    a. supervisord -c /home/apps/supervisord.conf -n
    
    b. curl "http://127.0.0.1:9080/?start=yes&num=3"

#OR

#2.
    a. docker build -t="centos/ms" . 
    
    b. docker run -p 9080:9080 -t -i centos/ms OR docker run --net=host -t -i centos/ms

 


 