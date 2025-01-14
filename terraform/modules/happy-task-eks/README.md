<!-- START -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 4.45 |
| <a name="requirement_kubernetes"></a> [kubernetes](#requirement\_kubernetes) | >= 2.16 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | >= 4.45 |
| <a name="provider_kubernetes"></a> [kubernetes](#provider\_kubernetes) | >= 2.16 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [kubernetes_cron_job.task_definition](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/cron_job) | resource |
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_backoff_limit"></a> [backoff\_limit](#input\_backoff\_limit) | kubernetes\_cron\_job backoff\_limit | `number` | `2` | no |
| <a name="input_cmd"></a> [cmd](#input\_cmd) | Command to run | `list(string)` | `[]` | no |
| <a name="input_cpu"></a> [cpu](#input\_cpu) | CPU shares (1cpu=1000m) per pod | `string` | `"100m"` | no |
| <a name="input_deployment_stage"></a> [deployment\_stage](#input\_deployment\_stage) | The name of the deployment stage of the Application | `string` | n/a | yes |
| <a name="input_failed_jobs_history_limit"></a> [failed\_jobs\_history\_limit](#input\_failed\_jobs\_history\_limit) | kubernetes\_cron\_job failed jobs history limit | `number` | `5` | no |
| <a name="input_image"></a> [image](#input\_image) | Image name | `string` | n/a | yes |
| <a name="input_k8s_namespace"></a> [k8s\_namespace](#input\_k8s\_namespace) | K8S namespace for this task | `string` | n/a | yes |
| <a name="input_memory"></a> [memory](#input\_memory) | Memory in megabits per pod | `string` | `"100Mi"` | no |
| <a name="input_remote_dev_prefix"></a> [remote\_dev\_prefix](#input\_remote\_dev\_prefix) | S3 storage path / db schema prefix | `string` | `""` | no |
| <a name="input_stack_name"></a> [stack\_name](#input\_stack\_name) | Happy Path stack name | `string` | n/a | yes |
| <a name="input_starting_deadline_seconds"></a> [starting\_deadline\_seconds](#input\_starting\_deadline\_seconds) | kubernetes\_cron\_job starting\_deadline\_seconds | `number` | `30` | no |
| <a name="input_successful_jobs_history_limit"></a> [successful\_jobs\_history\_limit](#input\_successful\_jobs\_history\_limit) | kubernetes\_cron\_job successful\_jobs\_history\_limit | `number` | `5` | no |
| <a name="input_task_name"></a> [task\_name](#input\_task\_name) | Happy Path task name | `string` | n/a | yes |
| <a name="input_ttl_seconds_after_finished"></a> [ttl\_seconds\_after\_finished](#input\_ttl\_seconds\_after\_finished) | kubernetes\_cron\_job ttl\_seconds\_after\_finished | `number` | `10` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_task_definition_arn"></a> [task\_definition\_arn](#output\_task\_definition\_arn) | Task definition name |
<!-- END -->