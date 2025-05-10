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
execute commands
```bash
sugar exec --services database --cmd "psql -U postgres -c 'SELECT version();'"
```

start specific services

```bash
sugar  --profile production compose-ext start --services api,worker
```

monitor resources for api and worker services even in swarm cluster
```bash
sugar stats plot --services api,worker 
```

stop stop specific services 
```bash
sugar  --profile production compose-ext stop --services worker
```

check logs for specific services 
```bash
sugar  --profile production compose-ext logs --services api --options "--tail 10"
```


rebuild only certain services
```bash
sugar --profile production  build --services api,database,cache --options "--no-cache"
```

swarm operations

```bash
# Deploy stack to swarm
sugar --profile development swarm deploy --stack complex-app 
```

remove the swarm stack 
```bash
sugar --profile development swarm rm --stack complex-app
```