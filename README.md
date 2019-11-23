Heroku: https://floating-lake-41216.herokuapp.com/

OpenStack: http://10.212.139.170:5632/project/v1/weather/?city=Gjøvik

OpenStack uses Docker

Original plan: Easy to find and use data related to traveling, you can look up the weather as well as the best attractions in an area, in the end the only thing missing from the project was a little more substance

What went well: Technically sound, especially openstack, json parsing and docker

What went badly: abscense of time to work on the "meat" of the project

Hard aspects: It was very hard thinking of a good original plan, in the end  I am very happy with it.

Learned: In-depth API-usage, .json file parsing, docker-images

Total work time: ~30hours (One person project)

Endpoints:

/project/v1/weather?{city=CITY_NAME} 

or, it is possible to ask for the weather by passing coordinates, like this

/project/v1/weather?{lon=LONGITUDE&lat=LATITUDE} 

/project/v1/places?city=CITY_NAME

/project/v1/status/
