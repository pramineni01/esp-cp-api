# Enterprise SaaS Platform Collaborative Planning APIs

### Build and Deploy process
#### Pack the db sql templates into a binary
1. cd to main directory
2. packr2 build -v cmd/cp-api/server.go --ignore-imports
3. // modify resolver-packr.go file with the correct import

#### Build
4. Make
5. docker build -t esp-cp-api .

#### Docker tagging and pushing to registry
6. docker tag esp-cp-api:latest registry-dev.espdev.antuits.com/backend/esp-cp-api:v1.1.5
7. // connect to vpn
8. docker login registry-dev.espdev.antuits.com
9. docker push registry-dev.espdev.antuits.com/backend/esp-cp-api:v1.1.5
10. curl --user "devops:dh7as8gf9fh" -X GET https://registry-dev.espdev.antuits.com/v2/backend/esp-cp-api/tags/list

#### Deploy
11. kb set image deployment/esp-cp-api cp-api=registry-dev.espdev.antuits.com/backend/esp-cp-api:v1.1.5
12. kb rollout restart deployment.apps/esp-cp-api

NOTE: 'kb' is an alias to "kubectl --kubeconfig=/Users/praveen.ramineni/kubeconfig/esp-dev.conf -n backend"