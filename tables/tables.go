// This file is generated by GoAdmin CLI adm.
package tables

import (
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "env_config" => http://localhost:9033/admin/info/env_config
// "api_definition" => http://localhost:9033/admin/info/api_definition
// "api_relation" => http://localhost:9033/admin/info/api_relation
// "api_test_data" => http://localhost:9033/admin/info/api_test_data
// "api_fuzzing_data" => http://localhost:9033/admin/info/api_fuzzing_data
// "api_test_detail" => http://localhost:9033/admin/info/api_test_detail
// "api_test_result" => http://localhost:9033/admin/info/api_test_result
// "api_id_count" => http://localhost:9033/admin/info/api_id_count
// "product_count" => http://localhost:9033/admin/info/product_count
// "parameter_definition" => http://localhost:9033/admin/info/parameter_definition
// "fuzzing_definition" => http://localhost:9033/admin/info/fuzzing_definition
// "scene_data" => http://localhost:9033/admin/info/scene_data
// "playbook" => http://localhost:9033/admin/info/playbook
// "product" => http://localhost:9033/admin/info/product
// "schedule" => http://localhost:9033/admin/info/schedule
// "sys_parameter" => http://localhost:9033/admin/info/sys_parameter
// "scene_test_history" => http://localhost:9033/admin/info/scene_test_history
// "scene_data_test_history" => http://localhost:9033/admin/info/scene_data_test_history
// "test_case" => http://localhost:9033/admin/info/test_case
// "app_api_changelog" => http://localhost:9033/admin/info/app_api_changelog
// "assert_template" => http://localhost:9033/admin/info/assert_template
//
// example end
//
var Generators = map[string]table.Generator{
	//"user":                    datamodel.GetUserTable,
	"env_config":              GetEnvConfigTable,
	"api_definition":          GetApiDefinitionTable,
	"api_relation":            GetApiRelationTable,
	"api_test_data":           GetApiTestDataTable,
	"api_fuzzing_data":        GetApiFuzzingDataTable,
	"api_test_detail":         GetApiTestDetailTable,
	"api_test_result":         GetApiTestResultTable,
	"api_id_count":            GetApiIdCountTable,
	"product_count":           GetProductCountTable,
	"parameter_definition":    GetParameterDefinitionTable,
	"fuzzing_definition":      GetFuzzingDefinitionTable,
	"scene_data":              GetSceneDataTable,
	"playbook":                GetPlaybookTable,
	"product":                 GetProductTable,
	"schedule":                GetScheduleTable,
	"sys_parameter":           GetSysParameterTable,
	"scene_test_history":      GetSceneTestHistoryTable,
	"scene_data_test_history": GetSceneDataTestHistoryTable,
	"test_case":               GetTestCaseTable,
	"api_custom":              GetApiCustomDefinitionTable,
	"app_api_changelog":       GetAppApiChangelogTable,
	"assert_template":         GetAssertTemplateTable,
	"ai_case":                 GetAiCaseTable,
	"ai_data":                 GetAiDataTable,
	"ai_issue":                GetAiIssueTable,
	"ai_playbook":             GetAiPlaybookTable,
	"ai_report":               GetAiReportTable,
	"ai_task":                 GetAiTaskTable,
	"ai_template":             GetAiTemplateTable,
	"ai_create":               GetAiCreateTable,
	"ai_optimize":             GetAiOptimizeTable,
	// generators end
}
