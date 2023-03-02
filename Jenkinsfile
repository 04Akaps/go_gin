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
        sh 'which go'
      }
    }

  }
}