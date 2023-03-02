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
        sh '''#!/bin/bash

# Download Go package
wget https://golang.org/dl/go1.17.5.linux-amd64.tar.gz

# Extract Go package
tar -C /usr/local -xzf go1.17.5.linux-amd64.tar.gz

# Add Go binary directory to the system PATH
echo \'export PATH=$PATH:/usr/local/go/bin\' >> /etc/profile

# Load the updated PATH variable
source /etc/profile

# Verify Go installation
go version'''
      }
    }

    stage('') {
      steps {
        sh 'go mod download'
      }
    }

  }
}