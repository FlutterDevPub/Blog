# Practical Distributed Systems: Implementing Raft Consensus, step by step - Part 1
## Introduction
We've already built an inconsistent - but highly available - database in my last blog post. Often, you need some level of consistency though. There are various consistency models for single operations in a database, and I won't go through them, as [Aphyr summed them up well here.][1] In a lot of databases you can actually choose the consistency level you need on a per-request basis. That's the case for [ZooKeeper][2] or [Consul][3]. 

In this blog post series though, we'll create a database which will provide *linearizable* operation consistency. What this means, is that we can create a timeline on which each interaction happened at some instant between the start and completion of that interaction (This may be the invocation and response of a web request for example). So one of the things this gets you, is that whenever you successfully write something to the database, all subsequent reads from any node will be up to date with that write. Another [post by Aphyr][4] describes strong consistency models in more detail.

We'll build this using a distributed commit log. So we'll have a log of operations which gets replicated between all database servers in the cluster. We will achieve *consensus* in regards to the data contained by the subsequent messages in the loge by using Raft Consensus. It's a well known consensus protocol used in, amongst others: [Consul][3], [etcd][5], [CockroachDB][7] or [RethinkDB][6].

## Raft Consensus
*"What's consensus?"* - you may ask yourself. We say that multiple parties achieve consensus when they all agree on some value. In our case, they'll agree on the values at each subsequent index in the log.

OK, we know what consensus is, but why use Raft? Well, it works and is relatively simple to reason about. It all began with [Paxos][8], a protocol widely known to be hard to understand. Opinions on that vary, but it surely is difficult to implement correctly, as described in [Paxos Made Life][9]. Alternatives emerged and one of the most widely used is Raft.

To get an intuition about how Raft works, watch this interactive visualization: http://thesecretlivesofdata.com/raft . You should also read the paper available [here][10] to understand Raft, but at least read page 4, as you won't understand the code otherwise. The variable names used will intentionally be similar to those described there.

Okay, ready to dive in? Let's go then!

## Preparations
First, we'll need to get a few packages we'll depend on. I'll go over each of them here:
#### Serf
Address: ```github.com/hashicorp/serf/serf```
Serf provides a cluster abstraction. It manages health checks and membership using the SWIM protocol.
#### gRPC
Address: ```google.golang.org/grpc```
gRPC is a remote procedure call library, which also autogenerates the clients for us.
Install the protobuf compiler as described in https://grpc.io/docs/quickstart/go.html.

We'll also use a gRPC connection cache I've written, which you can get here:
```github.com/cube2222/grpc-connection-cache```
I won't go over it as it wouldn't add anything meaningful here, but you can go over the code, as it's just one short file.
#### Others
We'll also use ```github.com/uber-go/atomic``` as a convenient wrapper over atomic variables, ```github.com/satori/go.uuid``` to generate uuid's and ```github.com/gorilla/mux```  as our router.

### Foreword
We'll take a bottom-up approach in the code. This is obviously not what I've done when I've writing this, but it's the only way to organize this in a sensible way in the form of a blog post.




[1]:https://github.com/aphyr/distsys-class
[2]:https://zookeeper.apache.org/
[3]:https://www.consul.io/
[4]:https://aphyr.com/posts/313-strong-consistency-models
[5]:https://github.com/coreos/etcd
[6]:https://www.rethinkdb.com/
[7]:https://www.cockroachlabs.com/
[8]:https://lamport.azurewebsites.net/pubs/lamport-paxos.pdf
[9]:https://research.google.com/archive/paxos_made_live.html
[10]:https://raft.github.io/raft.pdf
