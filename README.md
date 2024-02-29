# Gedis

A simple clone of Redis built in Go.

## Overview

Gedis stands for "Go Redis" and offers a lightweight, efficient, and scalable key-value store solution. It's designed to be easy to use and integrate into your projects while maintaining compatibility with Redis commands and protocols.

## Features

- Key-Value Store: Gedis stores data in a key-value format, allowing fast retrieval and storage of information.
- Redis Protocol Compatible: Gedis supports the Redis protocol, allowing clients written for Redis to seamlessly communicate with Gedis without modifications.
- Lightweight: Implemented in Go, Gedis is lightweight and consumes fewer system resources compared to some other solutions.
- Docker Support: Gedis provides a Docker image for easy deployment and integration into containerized environments.

## Getting Started

To get started with Gedis, you can pull the public Docker image available on Docker Hub:

```bash
docker pull thalison/gedis:latest
```

After pulling the image, you can run Gedis using Docker:
```bash
docker run -p 6379:6379 thalison/gedis:latest
```

This command starts Gedis and maps port 6379 on the host to port 6379 within the Docker container.

## Usage

Once Gedis is running, you can connect to it using any Redis client that supports the Redis protocol. You can use familiar Redis commands to interact with Gedis.

For example, using the [redis-cli](https://redis.io/download/) command-line tool:

```bash
redis-cli
```

## Available Commands
Here's the list of available commands and how to call them:

1. **SET**
   - *Description*: Set the string value of a key.
   - *Syntax*: `SET key value`

2. **GET**
   - *Description*: Get the value of a key.
   - *Syntax*: `GET key`

3. **HSET**
   - *Description*: Set the value of a field in a hash stored at the key.
   - *Syntax*: `HSET key field value`

4. **HGET**
   - *Description*: Get the value associated with a field in a hash stored at the key.
   - *Syntax*: `HGET key field`

5. **HGETALL**
   - *Description*: Get all the fields and values in a hash stored at the key.
   - *Syntax*: `HGETALL key`
