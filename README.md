# chuck
If you visit http://www.icndb.com/the-jokes-2/ you’ll find a good amount of chuck norris jokes. 
This Go code will pull in jokes using the http://www.icndb.com/api/ api, store them into a cache instance, write them to a file while running within a docker
container. The application pull in one joke at a time every 3 seconds and add it to a caching system. 
After a minute it flushes the cache and write all the jokes to a file.
