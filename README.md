# Driver-Passenger-Matching


docker-compose up

or

want to run local golang project, just need mongodb instance in docker-compose and isLocal variable must change to true in config/config.yml 

swag init

go mod tiny

-------------------------

localhost:8080/swagger

localhost:8081 mongodb ui

-----------------

[PUT] /api/driver insert all drivers from Coordinates.csv file to database

[POST] /api/driver insert a new driver 

[POST] /api/passenger insert a new passenger 

[PUT] /api/matcher/findDriver find the nearest driver to passenger 

[PUT] /api/matcher match a passenger with a driver


![image](https://user-images.githubusercontent.com/46713778/162590279-e1de0d67-4ab0-44f2-9b78-c98a19ec45bf.png)
