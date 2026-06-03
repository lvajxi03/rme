pipeline {
    agent {
	kubernetes {
	    yaml '''
apiVersion: v1
kind: Pod
spec:
  hostAliases:
      - ip: "10.10.10.24"
        hostnames:
          - "registry.lab.local"
  containers:
    - name: rme-go-builder
      image: registry.lab.local/go-builder:1.26.1
      command: ["sleep"]
      args: ["infinity"]
    - name: jnlp
      image: jenkins/inbound-agent:latest-jdk21
    - name: rme-kaniko
      image: registry.lab.local/kaniko-builder:latest
      command: ["sleep"]
      args: ["infinity"]
      volumeMounts:
        - name: docker-config
          mountPath: /kaniko/.docker
  volumes:
    - name: docker-config
      configMap:
        name: kaniko-docker-config
'''
	    
	}
    }

    environment {
	IMAGE_NAME = 'registry.lab.local/rme'
    }

    stages {
	stage('Prepare') {
	    steps {
		container('rme-go-builder') {
		    script {
			env.VERSION = env.GIT_COMMIT.take(7)			

			currentBuild.displayName = "#${env.BUILD_NUMBER} ${env.VERSION}"
			currentBuild.description = "${env.IMAGE_NAME}:${env.VERSION}"
		    }
		}
	    }
	}


	stage('Build binary') {
	    steps {
		container('rme-go-builder') {
		    sh '''
		        git config --global --add safe.directory "$WORKSPACE"
		        CGO_ENABLED=0 GOOS=linux go build -o rme .
                    '''
		}
	    }
	}

	stage('Build and push image') {
	    steps {
		container('rme-kaniko') {
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
