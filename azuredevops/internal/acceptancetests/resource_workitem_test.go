package acceptancetests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/acceptancetests/testutils"
)

func TestAccWorkItem_basic(t *testing.T) {
	workItemTitle := testutils.GenerateResourceName()
	projectName := testutils.GenerateResourceName()
	tfNode := "azuredevops_workitem.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutils.PreCheck(t, nil) },
		ProviderFactories: testutils.GetProviderFactories(),
		CheckDestroy:      testutils.CheckProjectDestroyed,
		Steps: []resource.TestStep{
			{
				Config: workItemBasic(projectName, workItemTitle),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttr(tfNode, "type", "Issue"),
					resource.TestCheckResourceAttr(tfNode, "state", "Active"),
				),
			},
			{
				ResourceName:      tfNode,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(tfNode),
			},
		},
	})
}

func TestAccWorkItem_titleUpdate(t *testing.T) {
	workItemTitle := testutils.GenerateResourceName()
	workItemTitleUpdated := testutils.GenerateResourceName()
	projectName := testutils.GenerateResourceName()
	tfNode := "azuredevops_workitem.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutils.PreCheck(t, nil) },
		ProviderFactories: testutils.GetProviderFactories(),
		CheckDestroy:      testutils.CheckProjectDestroyed,
		Steps: []resource.TestStep{
			{
				Config: workItemBasic(projectName, workItemTitle),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
				),
			},
			{
				ResourceName:      tfNode,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(tfNode),
			},
			{
				Config: workItemBasic(projectName, workItemTitleUpdated),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitleUpdated),
				),
			},
			{
				ResourceName:      tfNode,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(tfNode),
			},
		},
	})
}

func TestAccWorkItem_tagUpdate(t *testing.T) {
	workItemTitle := testutils.GenerateResourceName()
	projectName := testutils.GenerateResourceName()
	tfNode := "azuredevops_workitem.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutils.PreCheck(t, nil) },
		ProviderFactories: testutils.GetProviderFactories(),
		CheckDestroy:      testutils.CheckProjectDestroyed,
		Steps: []resource.TestStep{
			{
				Config: workItemBasic(projectName, workItemTitle),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
				),
			},
			{
				ResourceName:      tfNode,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(tfNode),
			},
			{
				Config: workItemTagUpdate(projectName, workItemTitle),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
				),
			},
			{
				ResourceName:      tfNode,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(tfNode),
			},
			{
				Config: workItemBasic(projectName, workItemTitle),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
				),
			},
			{
				ResourceName:      tfNode,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(tfNode),
			},
		},
	})
}

func TestAccWorkItem_parent(t *testing.T) {
	workItemTitle := testutils.GenerateResourceName()
	projectName := testutils.GenerateResourceName()
	tfNode := "azuredevops_workitem.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutils.PreCheck(t, nil) },
		ProviderFactories: testutils.GetProviderFactories(),
		CheckDestroy:      testutils.CheckProjectDestroyed,
		Steps: []resource.TestStep{
			{
				Config: workItemParent(projectName, workItemTitle),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", "Issue"),
					resource.TestCheckResourceAttr(tfNode, "state", "Active"),
				),
			},
			{
				ResourceName:      tfNode,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(tfNode),
			},
		},
	})
}

func TestAccWorkItem_parentUpdate(t *testing.T) {
	workItemTitle := testutils.GenerateResourceName()
	projectName := testutils.GenerateResourceName()
	tfNode := "azuredevops_workitem.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutils.PreCheck(t, nil) },
		ProviderFactories: testutils.GetProviderFactories(),
		CheckDestroy:      testutils.CheckProjectDestroyed,
		Steps: []resource.TestStep{
			{
				Config: workItemParent(projectName, workItemTitle),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", "Issue"),
					resource.TestCheckResourceAttr(tfNode, "state", "Active"),
				),
			},
			{
				ResourceName:      tfNode,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(tfNode),
			},
			{
				Config: workItemParentUpdate(projectName, workItemTitle),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", "Issue"),
					resource.TestCheckResourceAttr(tfNode, "state", "Active"),
				),
			},
			{
				ResourceName:      tfNode,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(tfNode),
			},
		},
	})
}

func TestAccWorkItem_parentDelete(t *testing.T) {
	workItemTitle := testutils.GenerateResourceName()
	projectName := testutils.GenerateResourceName()
	tfNode := "azuredevops_workitem.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutils.PreCheck(t, nil) },
		ProviderFactories: testutils.GetProviderFactories(),
		CheckDestroy:      testutils.CheckProjectDestroyed,
		Steps: []resource.TestStep{
			{
				Config: workItemParent(projectName, workItemTitle),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", "Issue"),
					resource.TestCheckResourceAttr(tfNode, "state", "Active"),
				),
			},
			{
				ResourceName:      tfNode,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(tfNode),
			},
			{
				Config: workItemParentDelete(projectName, workItemTitle),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", "Issue"),
					resource.TestCheckResourceAttr(tfNode, "state", "Active"),
				),
			},
			{
				ResourceName:            tfNode,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       testutils.ComputeProjectQualifiedResourceImportID(tfNode),
				ImportStateVerifyIgnore: []string{"parent_id"},
			},
		},
	})
}

func TestAccWorkItem_additionalFieldsJson(t *testing.T) {
	workItemTitle := testutils.GenerateResourceName()
	projectName := testutils.GenerateResourceName()
	tfNode := "azuredevops_workitem.test"
	storyPoints := 5.00
	storyPointsUpdate := 3.2
	acceptanceCriteria := testutils.GenerateResourceName()
	acceptanceCriteriaUpdate := testutils.GenerateResourceName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutils.PreCheck(t, nil) },
		ProviderFactories: testutils.GetProviderFactories(),
		CheckDestroy:      testutils.CheckProjectDestroyed,
		Steps: []resource.TestStep{
			{
				Config: workItemAdditionalFields(projectName, workItemTitle, fmt.Sprintf("%f", storyPoints), acceptanceCriteria),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", "User Story"),
					resource.TestCheckResourceAttr(tfNode, "state", "New"),
					resource.TestCheckResourceAttrWith(tfNode, "additional_fields_json", func(value string) error {
						var m map[string]interface{}
						if err := json.Unmarshal([]byte(value), &m); err != nil {
							return err
						}

						if m["Microsoft.VSTS.Scheduling.StoryPoints"] != storyPoints {
							return fmt.Errorf("expected Microsoft.VSTS.Scheduling.StoryPoints %f got %f", storyPoints, m["Microsoft.VSTS.Scheduling.StoryPoints"])
						}

						if m["Microsoft.VSTS.Common.AcceptanceCriteria"] != acceptanceCriteria {
							return fmt.Errorf("expected Microsoft.VSTS.Common.AcceptanceCriteria %s, got %s", acceptanceCriteria, m["Microsoft.VSTS.Common.AcceptanceCriteria"])
						}

						return nil
					}),
				),
			},
			{
				Config: workItemAdditionalFields(projectName, workItemTitle, fmt.Sprintf("%f", storyPointsUpdate), acceptanceCriteriaUpdate),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", "User Story"),
					resource.TestCheckResourceAttr(tfNode, "state", "New"),
					resource.TestCheckResourceAttrWith(tfNode, "additional_fields_json", func(value string) error {
						var m map[string]interface{}
						if err := json.Unmarshal([]byte(value), &m); err != nil {
							return err
						}

						if m["Microsoft.VSTS.Scheduling.StoryPoints"] != storyPointsUpdate {
							return fmt.Errorf("expected Microsoft.VSTS.Scheduling.StoryPoints %f, got %f", storyPointsUpdate, m["Microsoft.VSTS.Scheduling.StoryPoints"])
						}

						if m["Microsoft.VSTS.Common.AcceptanceCriteria"] != acceptanceCriteriaUpdate {
							return fmt.Errorf("expected Microsoft.VSTS.Common.AcceptanceCriteria %s, got %s", acceptanceCriteriaUpdate, m["Microsoft.VSTS.Common.AcceptanceCriteria"])
						}

						return nil
					}),
				),
			},
			{
				ResourceName:      tfNode,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(tfNode),
				ImportStateVerifyIgnore: []string{
					// Since we filter fields based on config provided,
					// this is expected to have a Diff when import is run with no config
					"additional_fields_json",
				},
			},
		},
	})
}

func TestAccWorkItem_description(t *testing.T) {
	workItemTitle := testutils.GenerateResourceName()
	projectName := testutils.GenerateResourceName()
	tfNode := "azuredevops_workitem.test"
	description := testutils.GenerateResourceName()
	descriptionUpdate := testutils.GenerateResourceName()
	itemType := "User Story"
	itemTypeAlternative := "Issue"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutils.PreCheck(t, nil) },
		ProviderFactories: testutils.GetProviderFactories(),
		CheckDestroy:      testutils.CheckProjectDestroyed,
		Steps: []resource.TestStep{
			{
				Config: workItemDescription(projectName, workItemTitle, itemType, description),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", itemType),
					resource.TestCheckResourceAttr(tfNode, "state", "New"),
					resource.TestCheckResourceAttr(tfNode, "description", description),
				),
			},
			{
				Config: workItemDescription(projectName, workItemTitle, itemType, descriptionUpdate),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", itemType),
					resource.TestCheckResourceAttr(tfNode, "state", "New"),
					resource.TestCheckResourceAttr(tfNode, "description", descriptionUpdate),
				),
			},
			{
				ResourceName:      tfNode,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(tfNode),
			},
			{
				Config: workItemDescription(projectName, workItemTitle, itemTypeAlternative, descriptionUpdate),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", itemTypeAlternative),
					resource.TestCheckResourceAttr(tfNode, "state", "Active"),
					resource.TestCheckResourceAttr(tfNode, "description", descriptionUpdate),
				),
			},
		},
	})
}

func workItemTemplate(name string) string {
	return fmt.Sprintf(`
resource "azuredevops_project" "project" {
  name               = "%[1]s"
  description        = "%[1]s-description"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
}`, name)
}

func workItemBasic(projectName string, title string) string {
	template := workItemTemplate(projectName)
	return fmt.Sprintf(`
%s

resource "azuredevops_workitem" "test" {
  title      = "%s"
  project_id = azuredevops_project.project.id
  type       = "Issue"
}
`, template, title)
}

func workItemTagUpdate(projectName string, title string) string {
	template := workItemTemplate(projectName)
	return fmt.Sprintf(`
%s

resource "azuredevops_workitem" "test" {
  title      = "%s"
  project_id = azuredevops_project.project.id
  type       = "Issue"
  state      = "Active"
  tags       = ["tag1", "tag2"]
}
`, template, title)
}

func workItemParent(projectName string, title string) string {
	template := workItemTemplate(projectName)
	return fmt.Sprintf(`
%[1]s

resource "azuredevops_workitem" "parent" {
  title      = "%[2]s Parent"
  project_id = azuredevops_project.project.id
  type       = "Issue"
}

resource "azuredevops_workitem" "test" {
  title      = "%[2]s"
  project_id = azuredevops_project.project.id
  type       = "Issue"
  parent_id  = azuredevops_workitem.parent.id
}
`, template, title)
}

func workItemParentDelete(projectName string, title string) string {
	template := workItemTemplate(projectName)
	return fmt.Sprintf(`
%[1]s

resource "azuredevops_workitem" "parent" {
  title      = "%[2]s Parent"
  project_id = azuredevops_project.project.id
  type       = "Issue"
}

resource "azuredevops_workitem" "test" {
  title      = "%[2]s"
  project_id = azuredevops_project.project.id
  type       = "Issue"
}
`, template, title)
}

func workItemParentUpdate(projectName string, title string) string {
	template := workItemTemplate(projectName)
	return fmt.Sprintf(`
%[1]s

resource "azuredevops_workitem" "parent" {
  title      = "%[2]s Parent"
  project_id = azuredevops_project.project.id
  type       = "Issue"
}

resource "azuredevops_workitem" "parent2" {
  title      = "%[2]s Parent2"
  project_id = azuredevops_project.project.id
  type       = "Issue"
}

resource "azuredevops_workitem" "test" {
  project_id = azuredevops_project.project.id
  title      = "%[2]s"
  type       = "Issue"
  parent_id  = azuredevops_workitem.parent2.id
}
`, template, title)
}

func workItemDescription(projectName string, title string, itemType string, description string) string {
	template := workItemTemplate(projectName)
	return fmt.Sprintf(`
%s

resource "azuredevops_workitem" "test" {
  title       = "%s"
  project_id  = azuredevops_project.project.id
  type        = "%s"
  description = "%s"
}
`, template, title, itemType, description)
}

func workItemAdditionalFields(projectName string, title string, storyPoints string, acceptanceCriteria string) string {
	template := workItemTemplate(projectName)
	return fmt.Sprintf(`
%s

resource "azuredevops_workitem" "test" {
  title      = "%s"
  project_id = azuredevops_project.project.id
  type       = "User Story"
  additional_fields_json = jsonencode({
    "Microsoft.VSTS.Scheduling.StoryPoints"    = %s
    "Microsoft.VSTS.Common.AcceptanceCriteria" = "%s"
  })
}
`, template, title, storyPoints, acceptanceCriteria)
}
