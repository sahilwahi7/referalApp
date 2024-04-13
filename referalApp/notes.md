1. If we are calling functions from outside the package it should be in upper case for same package it is lower case

2. The struct should be always called by references in other package

3. If we have to pass implememntation if a interface need to create struct(concrete implementation of the interface)

------>
next steps: instead of a static array will use a map as a databse storage to store and fetch jobs this will cover ViewJob and PostJob

onload of repo instance we need to prefetch all jobs that are there so we will use the fetchjobs function directly while starting the server

various on load api's will also be called once it is initialized