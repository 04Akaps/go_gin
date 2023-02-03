해당 repo에서는 다양한 Module를 테스트 하는 목적을 가지고 있습니다.

Basic으로 간단한 CRUD를 gin을 통해서 구축해 두고 이후 redis, docker, Aws 추후에는 ETH, Klaytn등등 블록체인 까지 활용해볼 예정입니다.

# docker cli

`docker build -t go_gin_study_test .`

- -t == "name"
- . == path

`docker run -p ${입구 port}:8080 -d go_gin_study_test`

- 입구 포트로 들어오는 요청은 8080으로 전송하는 docker image띄우기
  -- for nginx proxy

`docker stop $(docker ps -a -q)`

- 모든 컨테이너 중지

`docker rmi $(docker images -q)`

- 모든 이미지 삭제

# Ubuntu EC2

`ssh -i 'key' ubuntu@ipaddr`

- EC2에 ssh연결
- 이떄 key의 권한에 대해서 에러가 발생하기 떄문에
- `chmod 400 or 600 'key'`를 입력하여 권한 변경

`sudo apt install kubuntu-desktop`

- GUI를 위해 설치
- 시간이 좀 걸림

X2go를 활용할 것이기 떄문에

`sudo apt install x2goserver x2goserver-xsession`

`sudo passwd ubuntu`

- 패스 워드 설정

`sudo reboot`

- 반드시 XQuartz를 설치 해야 함

만약 `unable to execute: startkde`라는 에러가 발생 한다면,

- ubuntu에서 `sudo apt install kde-full` 실행

만약 `unable to execute: xfce4-session`라는 에러가 발생을 한다면

- ubuntu에서 `sudo apt install xfce4` 실행

# golang path 설정

1. `sudo apt-get install wget`

2. `wget https://golang.org/dl/go1.16.5.linux-amd64.tar.gz`

3. `sudo tar –xvf go1.16.5.linux-amd64.tar.gz`

- 압축 파일 해제 후 압축 파일 지워주면 됨

4. Path 설정

```
sudo vi $HOME/.profile

export GOROOT="go패키지 경로"
export GOPATH="go작업 공간"
export PATH="$GOPATH/bin:$GOROOT/bin:$PATH"
```

5. . ./.profile

- profile파일 리로딩

6. go version 확인
