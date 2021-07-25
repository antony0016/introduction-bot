# introduction-bot

This line bot will show you all of my projects and skills, if you are interested in me, you can scan the QR Code below.

![line bot](https://i.imgur.com/fCpQAyt.png)

---

### Setup

1. rename **sample.env** to **.env**
2. fill all environment variables in .env file
3. build code to docker image
4. use **Docker run** command to create a container.

**!!please make sure your has a domain, and your server use https!!**

---

### Build Docker image:

```shell
docker build -t introduction-bot .
```

---

### Docker run:

```shell
docker run -d --name introduction-bot -p <your-port>:<your-port> introduction-bot
```

