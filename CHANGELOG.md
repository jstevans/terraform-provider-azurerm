## 2.0.0 (Unreleased)

NOTES:

* **Major Version:** Version 2.0 of the Azure Provider is a major version - some deprecated fields/resources have been removed - please [refer to the 2.0 upgrade guide for more information](https://www.terraform.io/docs/providers/azurerm/guides/2.0-upgrade-guide.html).
* **Provider Block:** The Azure Provider now requires that a `features` block is specified within the Provider block, which can be used to alter the behaviour of certain resources - [more information on the `features` block can be found in the documentation](https://www.terraform.io/docs/providers/azurerm/index.html#features).
* **Terraform 0.10/0.11:** Version 2.0 of the Azure Provider no longer supports Terraform 0.10 or 0.11 - you must upgrade to Terraform 0.12 to use version 2.0 of the Azure Provider.

FEATURES:

* **Custom Timeouts:** - all resources within the Azure Provider now allow configuring custom timeouts - please [see Terraform's Timeout documentation](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) and the documentation in each data source resource for more information.
* **Requires Import:** The Azure Provider now checks for the presence of an existing resource prior to creating it - which means that if you try and create a resource which already exists (without importing it) you'll be prompted to import this into the state.
* **New Resource:** `azurerm_linux_virtual_machine` [GH-5705]
* **New Resource:** `azurerm_linux_virtual_machine_scale_set` [GH-5705]
* **New Resource:** `azurerm_virtual_machine_scale_set_extension` [GH-5705]
* **New Resource:** `azurerm_windows_virtual_machine` [GH-5705]
* **New Resource:** `azurerm_windows_virtual_machine_scale_set` [GH-5705]

BREAKING CHANGES:

* Data Source: `azurerm_scheduler_job_collection` - This data source has been removed since it was deprecated [GH-5712]
* `azurerm_container_service` - This resource has been removed since it was deprecated [GH-5709]
* `azurerm_scheduler_job` - This resource has been removed since it was deprecated [GH-5712]
* `azurerm_scheduler_job_collection` - This resource has been removed since it was deprecated [GH-5712]

IMPROVEMENTS:

* Data Source: `azurerm_app_service_plan` - the deprecated `properties` block has been removed since these properties have been moved to the top level [GH-5717]
* Data Source: `azurerm_kubernetes_service_version` - support for filtering of preview releases [GH-5662]
* `azurerm_app_service_plan` - the deprecated `properties` block has been removed since these properties have been moved to the top level [GH-5717]
* `azurerm_availability_set` - updating the default value for `managed` from `false` to `true` [GH-5724]
* `azurerm_windows_virtual_machine` - fixing a bug when provisioning from a Shared Gallery image [GH-5661]

---

For information on v1.44.0 and prior releases, please see [the v1.44.0 changelog](https://github.com/terraform-providers/terraform-provider-azurerm/blob/v1.44.0/CHANGELOG.md).
