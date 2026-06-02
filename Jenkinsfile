pipeline {
    agent {
	label 'go-builder'
    }

    environment {
	IMAGE_NAME = 'registry.lab.local/rme'
    }

    stages {
	stage('Prepare') {
	    steps {
		container('go-builder') {
		    script {
			env.VERSION = sh(
			    script: 'git rev-parse --short HEAD',
			    returnStdout: true
			).trim()

			currentBuild.displayName = "#${env.BUILD_NUMBER} ${env.VERSION}"
			currentBuild.description = "${env.IMAGE_NAME}:${env.VERSION}"
		    }
		}
	    }
	}


	stage('Build binary') {
	    steps {
		container('go-builder') {
		    sh 'CGO_ENABLED=0 GOOS=linux go build -o rme .'
		}
	    }
	}

	stage('Build and push image') {
	    steps {
		container('kaniko') {
		    sh '''
            /kaniko/executor \
              --context "${WORKSPACE}" \
              --dockerfile "${WORKSPACE}/Dockerfile" \
              --destination "${IMAGE_NAME}:${VERSION}" \
              --destination "${IMAGE_NAME}:latest" \
              --insecure \
              --skip-tls-verify \
              --insecure-registry registry.lab.local
          '''
		}
	    }
	}
    }
}
