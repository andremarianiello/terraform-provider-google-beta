// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceMonitoringGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitoringGroupCreate,
		Read:   resourceMonitoringGroupRead,
		Update: resourceMonitoringGroupUpdate,
		Delete: resourceMonitoringGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMonitoringGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				Description: `A user-assigned name for this group, used only for display
purposes.`,
			},
			"filter": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The filter used to determine which monitored resources
belong to this group.`,
			},
			"is_cluster": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `If true, the members of this group are considered to be a
cluster. The system can perform additional analysis on
groups that are clusters.`,
			},
			"parent_name": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkRelativePaths,
				Description: `The name of the group's parent, if it has one. The format is
"projects/{project_id_or_number}/groups/{group_id}". For
groups with no parent, parentName is the empty string, "".`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `A unique identifier for this group. The format is
"projects/{project_id_or_number}/groups/{group_id}".`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceMonitoringGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	parentNameProp, err := expandMonitoringGroupParentName(d.Get("parent_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent_name"); !isEmptyValue(reflect.ValueOf(parentNameProp)) && (ok || !reflect.DeepEqual(v, parentNameProp)) {
		obj["parentName"] = parentNameProp
	}
	isClusterProp, err := expandMonitoringGroupIsCluster(d.Get("is_cluster"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("is_cluster"); !isEmptyValue(reflect.ValueOf(isClusterProp)) && (ok || !reflect.DeepEqual(v, isClusterProp)) {
		obj["isCluster"] = isClusterProp
	}
	displayNameProp, err := expandMonitoringGroupDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	filterProp, err := expandMonitoringGroupFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !isEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}

	lockName, err := replaceVars(d, config, "stackdriver/groups/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/projects/{{project}}/groups")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Group: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate), isMonitoringRetryableError)
	if err != nil {
		return fmt.Errorf("Error creating Group: %s", err)
	}
	if err := d.Set("name", flattenMonitoringGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Group %q: %#v", d.Id(), res)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
	}
	d.Set("name", name.(string))
	d.SetId(name.(string))

	return resourceMonitoringGroupRead(d, meta)
}

func resourceMonitoringGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil, isMonitoringRetryableError)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("MonitoringGroup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}

	if err := d.Set("parent_name", flattenMonitoringGroupParentName(res["parentName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("name", flattenMonitoringGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("is_cluster", flattenMonitoringGroupIsCluster(res["isCluster"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("display_name", flattenMonitoringGroupDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("filter", flattenMonitoringGroupFilter(res["filter"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}

	return nil
}

func resourceMonitoringGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	parentNameProp, err := expandMonitoringGroupParentName(d.Get("parent_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, parentNameProp)) {
		obj["parentName"] = parentNameProp
	}
	isClusterProp, err := expandMonitoringGroupIsCluster(d.Get("is_cluster"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("is_cluster"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, isClusterProp)) {
		obj["isCluster"] = isClusterProp
	}
	displayNameProp, err := expandMonitoringGroupDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	filterProp, err := expandMonitoringGroupFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}

	lockName, err := replaceVars(d, config, "stackdriver/groups/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Group %q: %#v", d.Id(), obj)
	_, err = sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutUpdate), isMonitoringRetryableError)

	if err != nil {
		return fmt.Errorf("Error updating Group %q: %s", d.Id(), err)
	}

	return resourceMonitoringGroupRead(d, meta)
}

func resourceMonitoringGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "stackdriver/groups/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Group %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete), isMonitoringRetryableError)
	if err != nil {
		return handleNotFoundError(err, d, "Group")
	}

	log.Printf("[DEBUG] Finished deleting Group %q: %#v", d.Id(), res)
	return nil
}

func resourceMonitoringGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<project>[^ ]+) (?P<name>[^ ]+)", "(?P<name>[^ ]+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenMonitoringGroupParentName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringGroupName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringGroupIsCluster(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringGroupDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringGroupFilter(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandMonitoringGroupParentName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringGroupIsCluster(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringGroupDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringGroupFilter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
