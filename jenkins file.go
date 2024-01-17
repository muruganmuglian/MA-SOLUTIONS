pipeline{
	agent any

tools{
    maven 'local_maven'
}

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
		deploy adapters: [tomcat9(credentialsId: '49f0cb99-4b66-4c5a-a536-6ea4123a6514', path: '', url: 'http://localhost:2525/')], contextPath: null, war: '**/.war'
	}


}
		}
	}
}
