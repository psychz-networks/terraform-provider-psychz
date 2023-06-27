# Psychz provider development

## Requirements

* [Terraform](https://www.terraform.io/downloads.html) 1.4.6 (to run tests)
* [Go](https://golang.org/doc/install) 1.20.4 (to build the provider plugin)

## Building the provider

*Note:* This project uses [Go Modules](https://blog.golang.org/using-go-modules)
 making it safe to work with it outside of your existing [GOPATH](http://golang.org/doc/code.html#GOPATH).
The instructions that follow assume a directory in your home directory outside of
the standard GOPATH (i.e `$HOME/development/`).

1. Clone Psychz Terraform Provider repository

   ```sh
   mkdir -p $HOME/development; cd $HOME/development 
   git clone https://github.com/psychz-networks/terraform-provider-psychz.git
   ```

2. Enter provider directory and compile it

   ```sh
   cd terraform-provider-psychz
   go build
   ```


### Testing the provider with Terraform

Once you've built the plugin binary (see [Developing the provider](#developing-the-provider) above), it can be incorporated within your Terraform environment using the `-plugin-dir` option. Subsequent runs of Terraform will then use the plugin from your development environment.

```sh
terraform init -plugin-dir $GOPATH/bin
```

## Manual provider installation

*Note:* manual provider installation is needed only for manual testing of custom
built Psychz provider plugin.

Manual installation process differs depending on Terraform version.
Run `terraform version` command to determine version of your Terraform installation.

### Terraform 1.4.6 and newer

1. Create `psychz.net/psychz/psychz/1.0.0/linux_amd64` directories
under:



   *Note:* adjust `linux_amd64` from above structure to match your *os_arch*

   ```sh
   mkdir -p ~/.terraform.d/plugins/psychz.net/psychz/psychz/1.0.0/linux_amd64
   ```

2. Copy Psychz provider **binary file** there.

   ```sh
   cp terraform-provider-psychz ~/.terraform.d/plugins/psychz.net/psychz/psychz/1.0.0/linux_amd64
   ```

3. In every Terraform template directory that uses Psychz provider, ship below
 `terraform.tf` file *(in addition to other Terraform files)*

   ```hcl
   terraform {
     required_providers {
       psychz = {
         source = "psychz.net/psychz/psychz"
         version = "1.0.0"
       }
     }
   }
   ```

4. **Done!**

   Local Psychz provider plugin will be used after `terraform init`
   command execution in Terraform template directory
