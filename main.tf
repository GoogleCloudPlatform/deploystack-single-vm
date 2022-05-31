variable "project_id" {
  type = string
}

variable "project_number" {
  type = string
}

variable "zone" {
  type = string
}

variable "region" {
  type = string
}

variable "basename" {
  type = string
}

variable "instance-disksize" {
  type = string
}
variable "instance-disktype" {
  type = string
}

variable "instance-image" {
  type = string
}

variable "instance-machine-type" {
  type = string
}

variable "instance-name" {
  type = string
}

variable "instance-tags" {
  type = list(string)
}





# Enabling services in your GCP project
variable "gcp_service_list" {
  description = "The list of apis necessary for the project"
  type        = list(string)
  default = [
    "compute.googleapis.com",
  ]
}

resource "google_project_service" "all" {
  for_each                   = toset(var.gcp_service_list)
  project                    = var.project_number
  service                    = each.key
  disable_dependent_services = false
  disable_on_destroy         = false
}

# Create Instance
resource "google_compute_instance" "instance" {
  name         = var.instance-name
  machine_type = var.instance-machine-type
  zone         = var.zone
  project      = var.project_id
  tags         = var.instance-tags


  boot_disk {
    auto_delete = true
    device_name = var.instance-name
    initialize_params {
      image = var.instance-image
      size  = var.instance-disksize
      type  = var.instance-disktype
    }
  }

  network_interface {
    network = "default"
    access_config {
      // Ephemeral public IP
    }
  }

  depends_on = [google_project_service.all]
}


output "cmd"{
  value = "gcloud compute ssh ${var.instance-name}"
}


