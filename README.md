# MySQL Dojo

I recently realized that I have very little experience with MySQL, and I'd like to up my game a bit, so following the pattern of the wildly successful [mongo-dojo](https://github.com/lpmi-13/mongo-dojo), this will be a project to do just that!

## Aims/Goals

- see what a MySQL database looks like under regular load (write heavy for my use case, but also some reads).
- set up a scenario where we need to add/delete indexes
- see what the performance implications are for partitioned tables
- see what the performance implications are for missing indexes
- see what long-running queries look like in the systems graphs
- see what effect different iops constraints have on what happens when the database is trying to do a lot of writes.
- see what involves a metadata lock and what it looks like when that happens (https://www.alibabacloud.com/blog/generation-and-handling-of-metadata-locks-on-rds-for-mysql-tables_308797)
- see if we can trigger different types of problems with replication (https://aws.amazon.com/premiumsupport/knowledge-center/rds-mysql-high-replica-lag/), though this is going to complicate the setup a bit, but this is what a production workload would look like, so we should probably have a read replica anyway
- see if we can trigger a "waiting for table flush" state (https://www.thegeekdiary.com/troubleshooting-mysql-query-hung-waiting-for-table-flush/#:~:text=There%20are%20three%20ways%20to,Restart%20the%20server)

...and then also practice fixing all these things and making it faster.

## Additional Wrinkles

I'd like to also be able to see in real-time, what a slowly increasing load looks like in the metrics...so my current plan is to implement a simple web frontend that has some knobs to twist and config to specify, which then updates a k3s cluster running on the host machine to update pod deployments.

I'm assuming that going from 5 to 10 to 20 to 50 pods all trying to simultaneously write data into a fairly small database instance will show graphs that look "bad" (:tm:). Dynamiting a locally running database has no consequences, and it's probably gonna be fun anyway, so let's get to it!

## Local runs

At the moment, this is just a docker compose stack, and it doesn't involve any k3s yet, just until I get the grafana dashboards looking right. Then I'll add in the background load, and finally get some pod configuration updating on the fly.

```
docker-compose up -d
```

and your dashboards should be available at `localhost:3000`.

To insert some data, run the data insert container at `data-insert-background` via `docker-compose up --build`.
