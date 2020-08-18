# Golang Job Queue Prometheus Exporter 
This repository consists of -
1. Golang exporter which collects job queue status information associated with the Slurm scheduler contained in the `exporter` directory. The exporter could be easily configured for different HPC scheddulers (e.g. TORQUE)
2. Golang tests & associated mock test data for the exporter contained in the `tests` directory.

## Exporter Configuration
The Golang Job Queue Exporter can be configured to run as a systemd service. This [tutorial](https://paulgorman.org/technical/blog/20171121184114.html) shows how to run a Golang package as a Linux Service which I have summarized below.
1. Download Exporter Package.
```
go get gitlab.com/surfprace/cathal-go/exporter
```

2. Navigate to the package location: `go get` will download packages to the `$HOME/go` directory if `$GOPATH` not set.
```
cd $HOME/go/bin

# OR

cd $GOPATH/bin
```

3. Copy `exporter` binary file to 'usr/local/bin' with `sudo` priviliges.
```
sudo cp exporter /usr/local/bin/
```

4. Write Golang Job Queue Exporter Service file.
```
sudo vim /etc/systemd/system/go-hpc-exporter.service
```
Copy in the following -
```
[Unit]
Description=HPC Scheduler Golang Exporter Service

[Service]
ExecStart=/usr/local/bin/exporter -command='squeue'

[Install]
WantedBy=multi-user.target
```
To ensure a real Slurm `squeue` command is running, it is recommended to change the `-command` flag in the above service file to -
```
squeue --all -h --format=%A,%j,%a,%g,%u,%P,%v,%D,%C,%T,%V,%M
```

5. Start and enable the Service at boot:
```
systemctl daemon-reload
systemctl enable go-hpc-exporter
systemctl start go-hpc-exporter
```

6. View running Service:
```
systemctl status go-hpc-exporter
```

## Exporter Description
The `exporter/exporter.go` file is the main exporter script which contains the `-command` flag to specify the Slurm job queue `squeue` command. By default, the script is running a dummy `squeue` command which will be replaced by the official squeue command. The recommended flag command to run is -
`squeue --all -h --format=%A,%j,%a,%g,%u,%P,%v,%D,%C,%T,%V,%M`

The `exporter/parser/parser.go` script is called by the main exporter which parses the above Slurm `squeue` command after the main exporter has run the `squeue` command with `subprocess.Popen()`.

A map of key-values pairs is returned to the main exporter where a Prometheus Gauge Metric is created and data is exposed over `http://localhost:8080/`. The metric is expored in the following format: `slurm_group{project_name=<project_name>, job_type=<job_type>}`.

## Testing Description
Testing is available under the `tests` directory which contains 3 Golang tests. The associated test mock data is contained in the `tests/tests_data` subdirectory. To run the above tests:
```
go test .
```
The testing ensures that the `parser.go` script functions correctly.

## Variations of Job Queue Exporter
This [repository](https://gitlab.com/surfprace/cathal) contains the exact same exporter written in Python3. The repository also includes an Ansible scipt on deploying the exporter with Prometheus, Grafana and NGINX reverse proxy completing a Time Series Monitoring System.

Unfortunatly, trying to deploy a Golang package with Ansible proved difficult due to path issues.

## Main Instructions to Deploy TSM System with Go Exporter
The `ansible` directory is nearly a direct copy of the Ansible script taken from the [Python3 Exporter](https://gitlab.com/surfprace/cathal) repository with the following changes.
1. No `hpc-exporter` deployment due to the above explained path issues trying to execute `go get` within the Ansible script.
2. The `prometheus.yml` file has been amended to include the Golang Exporter http endpoint at `http://localhost/8080` instead of `http://localhost/8000` where the Python3 Exporter is by default configured.

The Ansible playbook deploys Grafana, Pormetheus and NGINX reverse proxy. More details on the Ansible deployment of the system can be found [here](https://gitlab.com/surfprace/cathal).

Once the above instructions to configure the Golang Exporter have been completed, the remaining TSM system can be deployed with Ansible. The deployment of the Time Series Monitoring System occurs under the `localhost` domain.
1. Install Ansible & ansible-galaxy collection for Grafana:
```
sudo apt install ansible && ansible-galaxy collection install community.grafana
```
2. Clone Git Repository:
```
git clone https://gitlab.com/surfprace/cathal-go.git
```
4. From the `ansible` directory, run the ansible playbook & enter root user password when prompted:
```
ansible-playbook -K playbook.yml
```
5. Result:
* If you open your web browser and visit the following sites, metrics of each sub-system can be observed:
    * `http://localhost:80/grafana/metrics`
    * `http://localhost:19090/prometheus/metrics`
    * `http://localhost:8080/metrics` (go exporter)
* To view the main Prometheus page search `http://localhost:19090/prometheus/`.
* To view Grafana search `http://localhost:80/grafana/` where the Prometheus datasource and JSON dashboard have been preconfigured and graphs should be immediatly available.

### Squeue Mock Data for Golang Exporter
If you **do not** have Slurm configured and wish to test the deployment of the Golang Exporter with dummy `squeue` data, complete the following.
1. Install Python3 Exporter Package which will install a `squeue_dummy` Python `entry_point` under `/usr/local/bin` directory.
```
pip3 install job-queue-exporter
```
2. Change the `-command` flag in the `/etc/systemd/system/go-hpc-exporter.service` directory to `/usr/local/bin/squeue_dummy`.
