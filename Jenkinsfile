pipeline {
  agent any
  stages {
    stage('Checkout Code') {
      steps {
        git(url: 'https://github.com/jjimgo/go_gin', branch: 'main')
      }
    }

    stage('check ls-al') {
      steps {
        sh 'ls -al'
      }
    }

    stage('Download dependencies') {
      steps {
        sh '''dependencies=$(go list -f \'{{if not (or .Main .Indirect)}}{{.Path}}@{{.Version}}{{end}}\' -m all)

for dependency in $dependencies
do
  go get $dependency
done'''
      }
    }

  }
}