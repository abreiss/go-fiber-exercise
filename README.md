# go-fiber-exercise
To build an image:  
docker build -t userapi .

Confirm the image:  
docker images

Run as a container:  
docker run -p 8080:80 userapi

Then visit:  
http://localhost:8080

to remove a container first stop it, only do this if ran in background
otherwise cntrl c will stop
check all running
docker  ps
docker stop userapi

then remove it
docker rm userapi

