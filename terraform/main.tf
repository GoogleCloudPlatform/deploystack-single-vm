/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

resource "google_project_service" "all" {
  for_each                   = toset(var.gcp_service_list)
  project                    = var.project_number
  service                    = each.key
  disable_dependent_services = false
  disable_on_destroy         = false
}

resource "google_compute_network" "main" {
  provider                = google-beta
  name                    = "${var.basename}-network"
  auto_create_subnetworks = true
  project                 = var.project_id
}

resource "google_compute_firewall" "default-allow-ssh" {
  name    = "deploystack-allow-ssh"
  project = var.project_number
  network = google_compute_network.main.id

  allow {
    protocol = "tcp"
    ports    = ["22"]
  }

  source_ranges = ["0.0.0.0/0"]

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
    network = google_compute_network.main.id
    access_config {
      // Ephemeral public IP
    }
  }

  depends_on = [google_project_service.all]
}
