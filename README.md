해당 repo에서는 다양한 Module를 테스트 하는 목적을 가지고 있습니다.

Basic으로 간단한 CRUD를 gin을 통해서 구축해 두고 이후 redis, docker, Aws 추후에는 ETH, Klaytn등등 블록체인 까지 활용해볼 예정입니다.



# docker cli

`docker build -t go_gin_study_test .`
- -t ==  "name"
- . == path

`docker run -p ${입구 port}:8080 -d go_gin_study_test`
- 입구 포트로 들어오는 요청은 8080으로 전송하는 docker image띄우기
-- for nginx proxy