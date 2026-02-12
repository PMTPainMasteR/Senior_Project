import groovy.json.JsonOutput

def globalVariable(envName) {
    // !!!!----------------------------------------!!!! //
    // !!!!------------- Start to edit ------------!!!! //
    // !!!!----------------------------------------!!!! //
    env.project_group               = ""
    env.project_name                = ""
    env.project_version             = ""

    env.git_group_slug              = ""
    env.git_project_slug            = ""

    def application_language        = ["python":             false,
                                       "nodejs":             false,
                                       "golang":             true,
                                       "dotnet_core":        false,
                                       "java":               false,
                                       "php":                false,
                                       "dotnet_fw":          false,
                                       "nodejs_with_yarn":   false]

    def deploy_type                 = ["oc":                 false,
                                       "aks":                true,
                                       "eks":                false,
                                       "azure_function":     false,
                                       "appservice_srccode": false,
                                       "appservice_container": false]

    env.unit_test_base_image        = ""

    def automate_test               = ["api_test":           true,
                                       "ui_test":            true]

    def allow_failure               = ["trivy":              false,
                                       "sonarqube":          false,
                                       "blackduck":          false,
                                       "owasp_zap":          false,
                                       "coverity":           false,
                                       "performance_test":   false,
                                       "api_test":           false,
                                       "ui_test":            false]

    env.build_cmd                   = ""
    env.coverityID                  = "cov-jenkins"
    env.blackduckID                 = "blkduck-jenkins"

    def skip_stage                  = ["unit_test":          false,
                                       "quality_analysis":   true,
                                       "sca_black_duck":     true,
                                       "sast_coverity":      true,
                                       "image_scan_trivy":   false,
                                       "dast_owasp_zap":     true,
                                       "performance_test":   true,
                                       "health_check_dev":   true,
                                       "automate_test_dev":  true,
                                       "health_check_sit":   true,
                                       "automate_test_sit":  true,
                                       "health_check_uat":   true,
                                       "automate_test_uat":  true,
                                       "health_check_prd":   true]

    def image_registry_server       = ["acr":                true,
                                       "nexus":              false,
                                       "ecr":                false,
                                       "gar":                false,
                                       "gcr":                false]

    def container_os_platform       = ["windows":            false,
                                       "linux":              true]

    env.is_scan_src_code_only           = false
    env.is_build_with_internal_net      = false
    env.time_sleep_before_healt_check   = 0
    env.coverity_exclude_path           = ""

    // NPM Private Registry
    env.has_npm_private_reg             = false
    env.npm_private_reg_path            = ""
    env.npm_private_reg_token           = "${project_group}-npm-registry"

    // DEV
    url_env_1                           = ""
    url_root_path_env_1                 = ""
    url_health_check_path_env_1         = ""

    // UAT
    url_env_3                           = ""
    url_root_path_env_3                 = ""
    url_health_check_path_env_3         = ""

    // PRD
    url_env_4                           = ""
    url_root_path_env_4                 = ""
    url_health_check_path_env_4         = ""

    //! Azure Container Registry //
    acr_credentials_cicd                = ""
    // DEV
    acr_server_env_1                    = ""
    // UAT
    acr_server_env_3                    = ""
    // PRD
    acr_server_env_4                    = ""

    // !!!!----------------------------------------!!!! //
    // !!!!-------------- End to edit -------------!!!! //
    // !!!!----------------------------------------!!!! //

    // !!!!----------------------------------------!!!! //
    // !!!!-------------- Do not edit -------------!!!! //
    // !!!!----------------------------------------!!!! //

    env.application_language    = JsonOutput.toJson(application_language)
    env.deploy_type             = JsonOutput.toJson(deploy_type)
    env.automate_test           = JsonOutput.toJson(automate_test)
    env.allow_failure           = JsonOutput.toJson(allow_failure)
    env.skip_stage              = JsonOutput.toJson(skip_stage)
    env.image_registry_server   = JsonOutput.toJson(image_registry_server)
    env.container_os_platform   = JsonOutput.toJson(container_os_platform)
}