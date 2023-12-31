pipeline{
	agent any

	stages{
		stage('build'){
			steps{
				sh 'maven clan package'
		
	}
	stage ( 'deploy to tomcat server'){
		post{
			success
			{
				echo "archiving the artifacts"
				archivingArtifacts artifacts: '**war/target/*.war'
			}
		}
	}
stage ('deploy to tomcat server'){
	steps{
		deploy adapters: deploy adapters: deploy adapters: [tomcat8(credentialsId: 'T', path: ' ' url: 'http://localhost:9090')], contextPath: null, war: '**/*.war'
	}


}
		}
	}
}
