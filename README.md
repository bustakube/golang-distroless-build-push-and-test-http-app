# Purpose:

This is a template for creating distroless golang containers, making a build, test, push process easy.

# Usage notes for build and push scripts:

- The Dockerfile must have a "#dockerhub: " header that indicates registry, account, repo, version...

# Usage notes for test script:
- The test script is for http requests to / only for now.
- The Dockerfile must have a "#testport: " head that indicates the port number on which the HTTP server will run.

