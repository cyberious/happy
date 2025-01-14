variable "cpu" {
  type        = string
  description = "CPU shares (1cpu=1000m) per pod"
  default     = "100m"
}

variable "memory" {
  type        = string
  description = "Memory in megabits per pod"
  default     = "100Mi"
}

variable "image" {
  type        = string
  description = "Image name"
}

variable "service_port" {
  type        = number
  description = "What ports does this service run on?"
  default     = 80
}

variable "desired_count" {
  type        = number
  description = "How many instances of this task should we run across our cluster?"
  default     = 2
}

variable "host_match" {
  type        = string
  description = "Host header to match for target rule. Leave empty to match all requests"
}

variable "stack_name" {
  type        = string
  description = "Happy Path stack name"
}

variable "cloud_env" {
  type = object({
    public_subnets : list(string),
    private_subnets : list(string),
    database_subnets : list(string),
    database_subnet_group : string,
    vpc_id : string,
    vpc_cidr_block : string,
  })
  description = "Typically data.terraform_remote_state.cloud-env.outputs"
}

variable "path" {
  type        = string
  description = "The path to register with the Application Load Balancer"
  default     = "/*"
}

variable "deployment_stage" {
  type        = string
  description = "The name of the deployment stage of the Application"
  default     = "dev"
}

variable "health_check_path" {
  type        = string
  description = "path to use for health checks"
  default     = "/"
}

variable "wait_for_steady_state" {
  type        = bool
  description = "Whether Terraform should block until the service is in a steady state before exiting"
  default     = true
}

variable "k8s_namespace" {
  type        = string
  description = "K8S namespace for this service"
}

variable "certificate_arn" {
  type        = string
  description = "ACM certificate ARN to attach to the load balancer listener"
}

variable "oauth_certificate_arn" {
  type        = string
  description = "Oauth Proxy ACM certificate ARN to attach to the load balancer listener"
}

variable "container_name" {
  type        = string
  description = "The name of the container"
}

variable "service_endpoints" {
  type        = map(string)
  default     = {}
  description = "Service endpoints to be injected for service discovery"
}

variable "service_name" {
  type        = string
  description = "Service name to be deployed"
}

variable "service_type" {
  type        = string
  description = "The type of the service to deploy. Supported types include 'EXTERNAL', 'INTERNAL', and 'PRIVATE'"
}

variable "period_seconds" {
  type        = number
  default     = 3
  description = "The period in seconds used for the liveness and readiness probes."
}

variable "initial_delay_seconds" {
  type        = number
  default     = 30
  description = "The initial delay in seconds for the liveness and readiness probes."
}

variable "success_codes" {
  type        = string
  default     = "200-499"
  description = "The range of success codes that are used by the ALB ingress controller."
}

variable "aws_iam_policy_json" {
  type        = string
  default     = ""
  description = "The AWS IAM policy to give to the pod."
}


variable "eks_cluster" {
  type = object({
    cluster_id : string,
    cluster_arn : string,
    cluster_endpoint : string,
    cluster_ca : string,
    cluster_oidc_issuer_url : string,
    cluster_security_group : string,
    cluster_iam_role_name : string,
    cluster_version : string,
    worker_iam_role_name : string,
    kubeconfig : string,
    worker_security_group : string,
    oidc_provider_arn : string,
  })
  description = "eks-cluster module output"
}

variable "additional_env_vars" {
  type        = map(string)
  description = "Additional environment variables to add to the task definition"
  default     = {}
}
