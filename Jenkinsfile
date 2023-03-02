pipeline {
  agent any
  stages {
    stage('check ls-al') {
      steps {
        sh 'ls -al'
      }
    }

    stage('Download dependencies') {
      steps {
        sh 'wget'
      }
    }

    stage('error') {
      steps {
        sh 'go mod download'
      }
    }

  }
}