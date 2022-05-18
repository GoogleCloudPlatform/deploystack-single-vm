variable "project_id" {
  type = string
}

variable "project_number" {
  type = string
}

variable "zone" {
  type = string
}

variable "basename" {
  type = string
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
resource "google_compute_instance" "sample" {
  name         = "${var.basename}-sample"
  machine_type = "n1-standard-1"
  zone         = var.zone
  project      = var.project_id


  boot_disk {
    auto_delete = true
    device_name = "${var.basename}-sample"
    initialize_params {
      image = "family/debian-10"
      size  = 200
      type  = "pd-standard"
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





