# kubernetes

이미지를 관리하기 위해 사용
로드 밸런싱이나, 이미지 재 시작을 담당
롤백 및 관리 시스템으로도 활용 가능

- 일단 docker에 들어가 설정창에서 `Enable Kubernetes`를 체크
- 이후 images를 확인

이후 brew를 활용하여 `minikube`를 설치

- 이것은 option

# cli 명령어

기본적으로 `kubectl`을 사용

1. `kubectl cluster-info`
   설치에 문제가 없다면 정보를 보여 준다.

## context

1. `kubectl config current-context`

   현재 사용하고 있는 컨텍스트를 보여 준다.

2. `kubectl config get-contexts`

   모든 컨텍스트들을 가져온다.

3. `kubectl config user-context [contextName]`

   컨텍스트를 바꾼다.

4. `kubectl config delete-context [contextName]`

   컨텍스트를 삭제한다.

5. `kubectl config rename-context [old-name] [new-name]`

   컨텍스트의 이름을 바꿉니다.

## deploy

ngnix기준 입니다.

1. `kubectl create deployment mynginx1 --image=ngnix`

   ngnix를 배포 합니다.

2. `kubectl get deploy`

   현재 배포된 배포 리스트를 가져 옵니다.

3. `kubectl create -f deploy-example.yaml`

   파일을 통해서 Deploy하는 형태 입니다.
