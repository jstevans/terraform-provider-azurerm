package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
)

func TestAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskBasic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_windows_virtual_machine_scale_set", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMWindowsVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskBasic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
		},
	})
}

func TestAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskCaching(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_windows_virtual_machine_scale_set", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMWindowsVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskCaching(data, "None"),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			), {
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskCaching(data, "ReadOnly"),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
			{
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskCaching(data, "ReadWrite"),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
		},
	})
}

func TestAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskDiskEncryptionSet(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_windows_virtual_machine_scale_set", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMWindowsVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				// TODO: After applying soft-delete and purge-protection in keyVault, this extra step can be removed.
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDisk_diskEncryptionSetDependencies(data),
				Check: resource.ComposeTestCheckFunc(
					enableSoftDeleteAndPurgeProtectionForKeyVault("azurerm_key_vault.test"),
				),
			},
			{
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDisk_diskEncryptionSet(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
		},
	})
}

func TestAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskResizing(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_windows_virtual_machine_scale_set", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMWindowsVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				// 30GB
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskResize(data, 30),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
			{
				// 60GB
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskResize(data, 60),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
		},
	})
}

func TestAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskMultiple(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_windows_virtual_machine_scale_set", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMWindowsVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskMultiple(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
		},
	})
}

func TestAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskRemove(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_windows_virtual_machine_scale_set", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMWindowsVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskBasic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
			{
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_authPassword(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
		},
	})
}

func TestAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskScaling(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_windows_virtual_machine_scale_set", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMWindowsVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				// no disks
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_authPassword(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
			{
				// one disk
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskBasic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
			{
				// two disks
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskMultiple(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
			{
				// no disks
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_authPassword(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
		},
	})
}

func TestAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskStorageAccountTypeStandardLRS(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_windows_virtual_machine_scale_set", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMWindowsVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskStorageAccountType(data, "Standard_LRS"),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
		},
	})
}

func TestAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskStorageAccountTypeStandardSSDLRS(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_windows_virtual_machine_scale_set", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMWindowsVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskStorageAccountType(data, "StandardSSD_LRS"),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
		},
	})
}

func TestAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskStorageAccountTypePremiumLRS(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_windows_virtual_machine_scale_set", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMWindowsVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskStorageAccountType(data, "Premium_LRS"),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
		},
	})
}

func TestAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskStorageAccountTypeUltraSSDLRS(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_windows_virtual_machine_scale_set", "test")

	// Are supported in East US 2, SouthEast Asia, and North Europe, in two availability zones per region
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMWindowsVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskStorageAccountTypeUltraSSDLRS(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
		},
	})
}

func TestAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskWriteAcceleratorEnabled(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_windows_virtual_machine_scale_set", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMWindowsVirtualMachineScaleSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskWriteAcceleratorEnabled(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMWindowsVirtualMachineScaleSetExists(data.ResourceName),
				),
			},
			data.ImportStep(
				"admin_password",
			),
		},
	})
}

func testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskBasic(data acceptance.TestData) string {
	template := testAccAzureRMWindowsVirtualMachineScaleSet_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_windows_virtual_machine_scale_set" "test" {
  name                = local.vm_name
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "Standard_F2"
  instances           = 1
  admin_username      = "adminuser"
  admin_password      = "P@ssword1234!"

  source_image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2019-Datacenter"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }

  data_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
    disk_size_gb         = 10
    lun                  = 10
  }

  network_interface {
    name    = "example"
    primary = true

    ip_configuration {
      name      = "internal"
      primary   = true
      subnet_id = azurerm_subnet.test.id
    }
  }
}
`, template)
}

func testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskCaching(data acceptance.TestData, caching string) string {
	template := testAccAzureRMWindowsVirtualMachineScaleSet_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_windows_virtual_machine_scale_set" "test" {
  name                = local.vm_name
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "Standard_F2"
  instances           = 1
  admin_username      = "adminuser"
  admin_password      = "P@ssword1234!"

  source_image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2019-Datacenter"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }

  data_disk {
    storage_account_type = "Standard_LRS"
    caching              = %q
    disk_size_gb         = 10
    lun                  = 10
  }

  network_interface {
    name    = "example"
    primary = true

    ip_configuration {
      name      = "internal"
      primary   = true
      subnet_id = azurerm_subnet.test.id
    }
  }
}
`, template, caching)
}

func testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskResize(data acceptance.TestData, diskSizeGb int) string {
	template := testAccAzureRMWindowsVirtualMachineScaleSet_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_windows_virtual_machine_scale_set" "test" {
  name                = local.vm_name
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "Standard_F2"
  instances           = 1
  admin_username      = "adminuser"
  admin_password      = "P@ssword1234!"

  source_image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2019-Datacenter"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }

  data_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
    lun                  = 10
    disk_size_gb         = %d
  }

  network_interface {
    name    = "example"
    primary = true

    ip_configuration {
      name      = "internal"
      primary   = true
      subnet_id = azurerm_subnet.test.id
    }
  }
}
`, template, diskSizeGb)
}

func testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDisk_diskEncryptionSetDependencies(data acceptance.TestData) string {
	// whilst this is in Preview it's only supported in: West Central US, Canada Central, North Europe
	// TODO: switch back to default location
	location := "westus2"

	name := testAccAzureRMWindowsVirtualMachineScaleSet_vmName(data)
	return fmt.Sprintf(`
locals {
  vm_name = "%s"
}
data "azurerm_client_config" "current" {}

resource "azurerm_resource_group" "test" {
  name     = "acctestrg-%d"
  location = "%s"
}

resource "azurerm_key_vault" "test" {
  name                        = "acctestkv%s"
  location                    = azurerm_resource_group.test.location
  resource_group_name         = azurerm_resource_group.test.name
  tenant_id                   = data.azurerm_client_config.current.tenant_id
  sku_name                    = "premium"
  enabled_for_disk_encryption = true
}

resource "azurerm_key_vault_access_policy" "service-principal" {
  key_vault_id = azurerm_key_vault.test.id
  tenant_id    = data.azurerm_client_config.current.tenant_id
  object_id    = data.azurerm_client_config.current.service_principal_object_id

  key_permissions = [
    "create",
    "delete",
    "get",
    "update",
  ]

  secret_permissions = [
    "get",
    "delete",
    "set",
  ]
}

resource "azurerm_key_vault_key" "test" {
  name         = "examplekey"
  key_vault_id = azurerm_key_vault.test.id
  key_type     = "RSA"
  key_size     = 2048

  key_opts = [
    "decrypt",
    "encrypt",
    "sign",
    "unwrapKey",
    "verify",
    "wrapKey",
  ]

  depends_on = ["azurerm_key_vault_access_policy.service-principal"]
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestnw-%d"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_subnet" "test" {
  name                 = "internal"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefix       = "10.0.2.0/24"
}
`, name, data.RandomInteger, location, data.RandomString, data.RandomInteger)
}

func testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDisk_diskEncryptionSetResource(data acceptance.TestData) string {
	template := testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDisk_diskEncryptionSetDependencies(data)
	return fmt.Sprintf(`
%s

resource "azurerm_disk_encryption_set" "test" {
  name                = "acctestdes-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  key_vault_key_id    = azurerm_key_vault_key.test.id

  identity {
    type = "SystemAssigned"
  }
}

resource "azurerm_key_vault_access_policy" "disk-encryption" {
  key_vault_id = azurerm_key_vault.test.id

  key_permissions = [
    "get",
    "wrapkey",
    "unwrapkey",
  ]

  tenant_id = azurerm_disk_encryption_set.test.identity.0.tenant_id
  object_id = azurerm_disk_encryption_set.test.identity.0.principal_id
}

resource "azurerm_role_assignment" "disk-encryption-read-keyvault" {
  scope                = azurerm_key_vault.test.id
  role_definition_name = "Reader"
  principal_id         = azurerm_disk_encryption_set.test.identity.0.principal_id
}
`, template, data.RandomInteger)
}

func testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDisk_diskEncryptionSet(data acceptance.TestData) string {
	template := testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDisk_diskEncryptionSetResource(data)
	return fmt.Sprintf(`
%s

resource "azurerm_windows_virtual_machine_scale_set" "test" {
  name                = local.vm_name
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "Standard_F2"
  instances           = 1
  admin_username      = "adminuser"
  admin_password      = "P@ssword1234!"

  source_image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2019-Datacenter"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }

  data_disk {
    storage_account_type   = "Standard_LRS"
    caching                = "ReadWrite"
    disk_encryption_set_id = azurerm_disk_encryption_set.test.id
    disk_size_gb           = 10
    lun                    = 10
  }

  network_interface {
    name    = "example"
    primary = true

    ip_configuration {
      name      = "internal"
      primary   = true
      subnet_id = azurerm_subnet.test.id
    }
  }

  depends_on = [
    "azurerm_role_assignment.disk-encryption-read-keyvault",
    "azurerm_key_vault_access_policy.disk-encryption",
  ]
}
`, template)
}

func testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskMultiple(data acceptance.TestData) string {
	template := testAccAzureRMWindowsVirtualMachineScaleSet_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_windows_virtual_machine_scale_set" "test" {
  name                = local.vm_name
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "Standard_F2"
  instances           = 1
  admin_username      = "adminuser"
  admin_password      = "P@ssword1234!"

  source_image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2019-Datacenter"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }

  data_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
    disk_size_gb         = 10
    lun                  = 10
  }

  data_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
    disk_size_gb         = 10
    lun                  = 20
  }

  network_interface {
    name    = "example"
    primary = true

    ip_configuration {
      name      = "internal"
      primary   = true
      subnet_id = azurerm_subnet.test.id
    }
  }
}
`, template)
}

func testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskStorageAccountType(data acceptance.TestData, storageAccountType string) string {
	template := testAccAzureRMWindowsVirtualMachineScaleSet_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_windows_virtual_machine_scale_set" "test" {
  name                = local.vm_name
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "Standard_F2s_v2"
  instances           = 1
  admin_username      = "adminuser"
  admin_password      = "P@ssword1234!"

  source_image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2019-Datacenter"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }

  data_disk {
    storage_account_type = %q
    caching              = "None"
    disk_size_gb         = 10
    lun                  = 10
  }

  network_interface {
    name    = "example"
    primary = true

    ip_configuration {
      name      = "internal"
      primary   = true
      subnet_id = azurerm_subnet.test.id
    }
  }
}
`, template, storageAccountType)
}

func testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskStorageAccountTypeUltraSSDLRS(data acceptance.TestData) string {
	template := testAccAzureRMWindowsVirtualMachineScaleSet_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_windows_virtual_machine_scale_set" "test" {
  name                = local.vm_name
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "Standard_D2s_v3"
  instances           = 1
  admin_username      = "adminuser"
  admin_password      = "P@ssword1234!"
  zones               = [1, 2, 3]

  source_image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2019-Datacenter"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }

  data_disk {
    storage_account_type = "UltraSSD_LRS"
    caching              = "None"
    disk_size_gb         = 10
    lun                  = 10
  }

  additional_capabilities {
    ultra_ssd_enabled = true
  }

  network_interface {
    name    = "example"
    primary = true

    ip_configuration {
      name      = "internal"
      primary   = true
      subnet_id = azurerm_subnet.test.id
    }
  }
}
`, template)
}

func testAccAzureRMWindowsVirtualMachineScaleSet_disksDataDiskWriteAcceleratorEnabled(data acceptance.TestData) string {
	template := testAccAzureRMWindowsVirtualMachineScaleSet_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_windows_virtual_machine_scale_set" "test" {
  name                = local.vm_name
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  sku                 = "Standard_M8ms"
  instances           = 1
  admin_username      = "adminuser"
  admin_password      = "P@ssword1234!"

  source_image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2019-Datacenter"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Premium_LRS"
    caching              = "None"
  }

  data_disk {
    storage_account_type      = "Premium_LRS"
    caching                   = "None"
    disk_size_gb              = 10
    lun                       = 10
    write_accelerator_enabled = true
  }

  network_interface {
    name    = "example"
    primary = true

    ip_configuration {
      name      = "internal"
      primary   = true
      subnet_id = azurerm_subnet.test.id
    }
  }
}
`, template)
}
