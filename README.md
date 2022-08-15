# Deploy Stack - SingleVM 

This is a stack that deploys a single VM. 

It's meant to be a very simple example to illustrate a very small set of 
Terraform capabilities

![Single VM architecture](/architecture.png)

## Install
You can install this application using the `Open in Google Cloud Shell` button 
below. 

<a href="https://ssh.cloud.google.com/cloudshell/editor?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2FGoogleCloudPlatform%2Fdeploystack-single-vm&shellonly=true&cloudshell_image=gcr.io/ds-artifacts-cloudshell/deploystack_custom_image" target="_new">
    <img alt="Open in Cloud Shell" src="https://gstatic.com/cloudssh/images/open-btn.svg">
</a>

Once this opens up, you can install by: 
1. Creating a Google Cloud Project
1. Then typing `./deploystack install`

## Cleanup 
To remove all billing components from the project
1. Typing `./deploystack uninstall`


This is not an official Google product.