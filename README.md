## SocialCops task
An efficient task runner and manipulator for SocialCops

<br/>

[![](https://img.shields.io/badge/docs%20-view%20API%20documentation-orange.svg?style=for-the-badge&logo=appveyor)](https://angadsharma1016.github.io/sc-task/)


[![demo](https://img.shields.io/badge/view%20demo-youtube-orange.svg?style=for-the-badge&logo=appveyor)](https://www.youtube.com/watch?v=Ldza4HUNZqg&feature=youtu.be)

<br/>
<br/>

<details open>

<summary>Technology stack</summary>

<br/>

- [X] Golang + MySQL
- [X] goroutines and channels for concurrency
- [X] NATS as a high availability messaging queue service
- [X] docker + docker-compose
- [X] apidoc documentation

### Why NATS?
---
[NATS](https://github.com/nats-io/go-nats.git) is an event sourcing tool which we will be using to publish logs and distribute related data between different services. The reason for NATS is:

* RabbitMQ is limited to HTTP and HTTPS natively

* NATS supports gRPC and works fast due to being type safe as well as removing extra overhead of marshalling and unmarshalling

* Publishing events on a different thread and subscribing from a different thread allows non-blocking log sourcing.

* NATS is very high throughput when it comes to requests per second

<br />

![NATS](./views/img/nats.png)

<br />
</details>

<br/>

#### Spawn process

<br/>

![spawn process](./views/img/spawn.png)

<br/>
<br/>

#### Process lifecycle

<br/>

![process lifecycle](./views/img/lifecycle.png)


<br/>

#### How to run
After running the commands below, wait for DB to build, then goto localhost:3000

```
$ docker-compose build
$ docker-compose up -d
```