include .env
.PHONY: init-dev plan-dev apply-dev destroy-dev

init-dev:
	cd environments/dev && \
	aws-vault exec $(AWS_ACCOUNT_NAME) -- terraform init \
	-backend-config=bucket="$(TF_BACKEND_BUCKET_DEV)" \
	-backend-config=key="$(TF_BUCKEND_KEY_DEV)" \

plan-dev:
	cd environments/dev && \
	aws-vault exec $(AWS_ACCOUNT_NAME) -- terraform plan \

apply-dev:
	cd environments/dev && \
	aws-vault exec $(AWS_ACCOUNT_NAME) -- terraform apply \
	-auto-approve

destroy-dev:
	cd environments/dev && \
	aws-vault exec $(AWS_ACCOUNT_NAME) -- terraform destroy \
	-auto-approve
