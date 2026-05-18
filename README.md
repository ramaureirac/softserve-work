# SoftServe Work


## part 1
Simple URL malware dectector. Uses a hashmap for storing hostnames and insecure paths. Uses `gin-gonic` to expose this db as an API server.

requirements:

- go
- make
- docker (optional)

build & run:

    git clone https://github.com/ramaureirac/softserve-work && cd softserve-work
    make build # or make docker-build
    ./bin/softserve # or docker run -p 8080:8080 -it ramaureirac/softserve:latest

test:

    make test


usage:

    # check url
    curl http://localhost:8080/urlinfo/hecker.info/dolphin.exe

    # register url
    curl -X POST http://localhost:8080/urlinfo/hecker.info/dolphin2.exe


## part 2

As a thought exercise, please describe how you would accomplish the following:
 
- The size of the URL list could grow infinitely. How might you scale this beyond the memory capacity of the system? 

A: in this case scenario the new proposal should include support for a nonsql db cluster. ideally simething similar to mongo/cassandra/dynamo that way storage can continue growing up without relay in-memory db. if a cache layer is needed redis or memcache may help. some alternaties may include using go-bloom as filter or use save only a initial part of a hash of full url. this may help on compress some space.

- Assume that the number of requests will exceed the capacity of a single system, describe how might you solve this, and how might this change if you have to distribute this workload to an additional region, such as Europe. 

A: since this project support docker, it can be easly replicated on different sites using kubernetes or swarm. if needed you can configure your dns provider (let say cloudflare) to redirect using geolocation. that way users in americas can be redirected to a site in the US and users in Europe can be delivered to a different infrastructure. only change needed will be maintaining database in sync. in this case scenario the site on Europe may be configured as replica of main DB on the US.

- What are some strategies you might use to update the service with new URLs? Updates may be as much as 5 thousand URLs a day with updates arriving every 10 minutes.

A: this project includes a POST endpoint for adding new URLs as malware. You should put this behind an API key so customers can add their new url during day. The change that should be done is manage so the endpoint can receive a pool of URLs to save into DB. Ideally, this logic should be separated in a different project with write access to DB. since they are a lot of transactions a queue may help for not over sature the database, something similar to kafka and have a worker pool of machines taking care of attend any request.


- You’re woken up at 3am, what are some of the things you’ll look for in the app?

A: if it's for an incident with the API i should check for logs, then perform other activities inside the machine where this app is instaled to understand better what is going on.

- Does that change anything you’ve done in the app?

A: add more logs and some kind of healthcheck.

- What are some considerations for the lifecycle of the app?

A: deliver deployment strats to prevent interruption, add some obsevability latency, traffic, errors. manage /healthz and /readyz enpoints. maintain the code base updated, resolve CVEs in dependencies, code and OS.  

- You need to deploy a new version of this application. What would you do?

A: If is deployed on Kubernetes you can safely do a kubectl rollout. k8s should not redirect traffic until /readyz is OK.