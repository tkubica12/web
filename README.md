# Simple web app for container demonstrations

This is very simple web service written in Go to demonstrate cloud, load-balancing, containers. It is compiled to Linux binary package running on port 3000 by default, but PORT env can be set to different number. Web server generate random UUID on start a present this on page (so load-balancing to multiple nodes can be demonstrated). V2 is the same, but string Version 2 is added (this is to demonstrate A/B testing or releases).

Use Dockerfile to build Docker images or reference tkubica/web:1 and tkubica/web:2 on Docker Hub.

Tomas Kubica
linked.in/c/tkubica
