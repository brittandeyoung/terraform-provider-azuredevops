package acceptancetests

import (
	"fmt"
	"strconv"
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

func TestAccWorkItem_storyPoints(t *testing.T) {
	workItemTitle := testutils.GenerateResourceName()
	projectName := testutils.GenerateResourceName()
	tfNode := "azuredevops_workitem.test"
	storyPoints := 5.0
	storyPointsUpdate := 3.2
	itemType := "User Story"
	itemTypeAlternative := "Issue"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutils.PreCheck(t, nil) },
		ProviderFactories: testutils.GetProviderFactories(),
		CheckDestroy:      testutils.CheckProjectDestroyed,
		Steps: []resource.TestStep{
			{
				Config: workItemStoryPoints(projectName, workItemTitle, itemType, storyPoints),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", itemType),
					resource.TestCheckResourceAttr(tfNode, "state", "New"),
					resource.TestCheckResourceAttr(tfNode, "story_points", strconv.FormatFloat(storyPoints, 'f', -1, 64)),
				),
			},
			{
				Config: workItemStoryPoints(projectName, workItemTitle, itemType, storyPointsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", itemType),
					resource.TestCheckResourceAttr(tfNode, "state", "New"),
					resource.TestCheckResourceAttr(tfNode, "story_points", strconv.FormatFloat(storyPointsUpdate, 'f', -1, 64)),
				),
			},
			{
				ResourceName:      tfNode,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(tfNode),
			},
			{
				Config: workItemStoryPoints(projectName, workItemTitle, itemTypeAlternative, storyPointsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", itemTypeAlternative),
					resource.TestCheckResourceAttr(tfNode, "state", "Active"),
					resource.TestCheckResourceAttr(tfNode, "story_points", strconv.FormatFloat(storyPointsUpdate, 'f', -1, 64)),
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

func TestAccWorkItem_acceptanceCriteria(t *testing.T) {
	workItemTitle := testutils.GenerateResourceName()
	projectName := testutils.GenerateResourceName()
	tfNode := "azuredevops_workitem.test"
	acceptanceCriteria := testutils.GenerateResourceName()
	acceptanceCriteriaUpdate := testutils.GenerateResourceName()
	itemType := "User Story"
	itemTypeAlternative := "Issue"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutils.PreCheck(t, nil) },
		ProviderFactories: testutils.GetProviderFactories(),
		CheckDestroy:      testutils.CheckProjectDestroyed,
		Steps: []resource.TestStep{
			{
				Config: workItemAcceptanceCriteria(projectName, workItemTitle, itemType, acceptanceCriteria),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", itemType),
					resource.TestCheckResourceAttr(tfNode, "state", "New"),
					resource.TestCheckResourceAttr(tfNode, "acceptance_criteria", acceptanceCriteria),
				),
			},
			{
				Config: workItemAcceptanceCriteria(projectName, workItemTitle, itemType, acceptanceCriteriaUpdate),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", itemType),
					resource.TestCheckResourceAttr(tfNode, "state", "New"),
					resource.TestCheckResourceAttr(tfNode, "acceptance_criteria", acceptanceCriteriaUpdate),
				),
			},
			{
				ResourceName:      tfNode,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(tfNode),
			},
			{
				Config: workItemAcceptanceCriteria(projectName, workItemTitle, itemTypeAlternative, acceptanceCriteriaUpdate),
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttr(tfNode, "title", workItemTitle),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "url"),
					resource.TestCheckResourceAttr(tfNode, "type", itemTypeAlternative),
					resource.TestCheckResourceAttr(tfNode, "state", "Active"),
					resource.TestCheckResourceAttr(tfNode, "acceptance_criteria", acceptanceCriteriaUpdate),
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

func workItemBasic(projectNane string, title string) string {
	template := workItemTemplate(projectNane)
	return fmt.Sprintf(`
%s

resource "azuredevops_workitem" "test" {
  title      = "%s"
  project_id = azuredevops_project.project.id
  type       = "Issue"
}
`, template, title)
}

func workItemTagUpdate(projectNane string, title string) string {
	template := workItemTemplate(projectNane)
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

func workItemParent(projectNane string, title string) string {
	template := workItemTemplate(projectNane)
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

func workItemParentDelete(projectNane string, title string) string {
	template := workItemTemplate(projectNane)
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

func workItemParentUpdate(projectNane string, title string) string {
	template := workItemTemplate(projectNane)
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

func workItemStoryPoints(projectNane string, title string, itemType string, storyPoints float64) string {
	template := workItemTemplate(projectNane)
	return fmt.Sprintf(`
%s

resource "azuredevops_workitem" "test" {
  title      = "%s"
  project_id = azuredevops_project.project.id
  type       = "%s"
  story_points = %f
}
`, template, title, itemType, storyPoints)
}

func workItemDescription(projectNane string, title string, itemType string, description string) string {
	template := workItemTemplate(projectNane)
	return fmt.Sprintf(`
%s

resource "azuredevops_workitem" "test" {
  title      = "%s"
  project_id = azuredevops_project.project.id
  type       = "%s"
  description = "%s"
}
`, template, title, itemType, description)
}

func workItemAcceptanceCriteria(projectNane string, title string, itemType string, acceptanceCriteria string) string {
	template := workItemTemplate(projectNane)
	return fmt.Sprintf(`
%s

resource "azuredevops_workitem" "test" {
  title      = "%s"
  project_id = azuredevops_project.project.id
  type       = "%s"
  acceptance_criteria = "%s"
}
`, template, title, itemType, acceptanceCriteria)
}
