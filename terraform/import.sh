terraform import google_compute_network.default ${TF_VAR_GOOGLE_PROJECT_ID}-network
terraform import google_compute_subnetwork.default ${TF_VAR_GOOGLE_PROJECT_ID}-subnetwork

terraform import google_compute_global_address.static_region_ip ${TF_VAR_GOOGLE_PROJECT_ID}/${TF_VAR_GOOGLE_PROJECT_ID}-region-address
terraform import google_compute_global_address.static_internal_ip ${TF_VAR_GOOGLE_PROJECT_ID}/${TF_VAR_GOOGLE_PROJECT_ID}-internal-address

terraform import google_vpc_access_connector.default ${TF_VAR_GOOGLE_PROJECT_ID}/${TF_VAR_GOOGLE_COMPUTE_REGION}/access-connector
terraform import google_compute_instance.default ${TF_VAR_GOOGLE_PROJECT_ID}-instance
terraform import google_compute_firewall.default ${TF_VAR_GOOGLE_PROJECT_ID}-firewall

terraform import google_app_engine_application.default ${TF_VAR_GOOGLE_PROJECT_ID}
true
