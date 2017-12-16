import re

from .base import Base

class Source(Base):
    def __init__(self, vim):
        super().__init__(vim)

        self.__pattern = re.compile(r'data |resource ')

        self.filetypes = ['tf']
        self.min_pattern_length = 3
        self.mark = '[tf-args]'
        self.matchers = ['matcher_length', 'matcher_full_fuzzy']
        self.name = 'tf-args'

    def gather_candidates(self, context):
        return [{'word': k} for (k, v) in ARGS.items()]

    def get_complete_position(self, context):
        match = self.__pattern.search(context['input'])
        return match.start() if match is not None else -1

ARGS = {
'data "aws_acm_certificate"' :"",
'data "aws_ami"' :"",
'data "aws_ami_ids"' :"",
'data "aws_autoscaling_groups"' :"",
'data "aws_availability_zone"' :"",
'data "aws_availability_zones"' :"",
'data "aws_billing_service_account"' :"",
'data "aws_caller_identity"' :"",
'data "aws_canonical_user_id"' :"",
'data "aws_cloudformation_stack"' :"",
'data "aws_cloudtrail_service_account"' :"",
'data "aws_db_instance"' :"",
'data "aws_db_snapshot"' :"",
'data "aws_dynamodb_table"' :"",
'data "aws_ebs_snapshot"' :"",
'data "aws_ebs_snapshot_ids"' :"",
'data "aws_ebs_volume"' :"",
'data "aws_ecr_repository"' :"",
'data "aws_ecs_cluster"' :"",
'data "aws_ecs_container_definition"' :"",
'data "aws_ecs_task_definition"' :"",
'data "aws_eip"' :"",
'data "aws_elastic_beanstalk_solution_stack"' :"",
'data "aws_elasticache_cluster"' :"",
'data "aws_elasticache_replication_group"' :"",
'data "aws_elb"' :"",
'data "aws_elb_hosted_zone_id"' :"",
'data "aws_elb_service_account"' :"",
'data "aws_iam_account_alias"' :"",
'data "aws_iam_group"' :"",
'data "aws_iam_instance_profile"' :"",
'data "aws_iam_policy_document"' :"",
'data "aws_iam_role"' :"",
'data "aws_iam_server_certificate"' :"",
'data "aws_iam_user"' :"",
'data "aws_instance"' :"",
'data "aws_instances"' :"",
'data "aws_internet_gateway"' :"",
'data "aws_ip_ranges"' :"",
'data "aws_kinesis_stream"' :"",
'data "aws_kms_alias"' :"",
'data "aws_kms_ciphertext"' :"",
'data "aws_kms_secret"' :"",
'data "aws_lb"' :"",
'data "aws_lb_listener"' :"",
'data "aws_lb_target_group"' :"",
'data "aws_nat_gateway"' :"",
'data "aws_network_interface"' :"",
'data "aws_partition"' :"",
'data "aws_prefix_list"' :"",
'data "aws_rds_cluster"' :"",
'data "aws_redshift_service_account"' :"",
'data "aws_region"' :"",
'data "aws_route53_zone"' :"",
'data "aws_route_table"' :"",
'data "aws_s3_bucket"' :"",
'data "aws_s3_bucket_object"' :"",
'data "aws_security_group"' :"",
'data "aws_sns_topic"' :"",
'data "aws_ssm_parameter"' :"",
'data "aws_subnet"' :"",
'data "aws_subnet_ids"' :"",
'data "aws_vpc"' :"",
'data "aws_vpc_endpoint"' :"",
'data "aws_vpc_endpoint_service"' :"",
'data "aws_vpc_peering_connection"' :"",
'data "aws_vpn_gateway"' :"",
'data "efs_file_system"' :"",
'data "efs_mount_target"' :"",
}
