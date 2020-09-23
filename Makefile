output:
	go get && go build

deploy:
	doctl compute droplet create --image ubuntu-20-04-x64 --size s-1vcpu-1gb --region nyc1 --ssh-keys ${SSHKEYS} winnie

ssh:
	doctl compute ssh winnie --ssh-port ${SSHPORT}

destroy:
	doctl compute droplet delete winnie

docker:
	docker build . -t winnie:0.1.0
	docker run --name winnie -p 22:22 winnie:0.1.0

clean:
	rm winnie-the-pot

clean-docker:
	docker stop winnie && docker rm winnie