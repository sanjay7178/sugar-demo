### Sugar Demo : Complex App

---

Default Project Name : example 

Default Profile : production


create and start specific services
```bash
sugar --profile production  compose up  --services api,worker 
```
create and start all services
```bash
 sugar --profile production  compose up  
```
start specific services

```bash
sugar  --profile production compose-ext start --services api,worker
```

rebuild only certain services
```bash
sugar build --profile production --services api,database,cache --options "--no-cache"
```

swarm operations

```bash
# Deploy stack to swarm
sugar --profile production swarm deploy --stack complex-app 

# Scale services
sugar swarm node --update worker --options "--replicas 5"
```