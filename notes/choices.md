# DB
## Mongodb
Pros
- Highly flexible document oriented db
- Easy to use and works well with JSON.
- Quering is simpler and just works

Cons
- Hard to make it HA
- Low performance

## Scylladb
Pros
- Highly performant when schema designed correctly
- Written in C++, no GC or JVM tuning
- High availability
- High latency guarantees

Cons
- Moderate devops
- Learning curve

## PostgreSQL with Citus Data
Pros
- Industry standard
- Power of schema and relations
- Highly stable
- Open source

Cons
- Making it highly available can be challenging

# Web server
## Node\ExpressJS
Pros
- Fast development
- Multiple middlewares, logging, db driver support
- Big community support

Cons
- Relatively low performance

## go fiber
Pros
- Highly performant
- ExpressJS like middleware and routing capabilities

Cons
- Relatively low number of out of box middleware
- Smaller community