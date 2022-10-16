# System requirements:
- Host - CPU 7, RAM 7 GiB, Memory 40 GiB 
  - p1 - CPU 2, RAM 2 GiB
    - vagrant vm1 - CPU 1, RAM 1 GiB (k3s server)
    - vagrant vm2 - CPU 1, RAM 1 GiB (k3s agent)
  - p2 - CPU 3, RAM 3 GiB
    - vagrant vm - CPU 3, RAM 3 GiB (k3s server)
  - p3 - CPU 3, RAM 3 GiB
    - Docker - RAM 1 GiB (k3d requirement)
    - ArgoCD - CPU 1, RAM 1 GiB (namespace)
    - app - CPU 1, RAM 1 GiB (namespace)
  - bonus - CPU 7, RAM 7 GiB
    - Docker - RAM 1 GiB (k3d requirement)
    - ArgoCD - CPU 1, RAM 1 GiB (namespace)
    - GitLab - CPU 4, RAM 4 GiB (namespace)
    - app - CPU 1, RAM 1 GiB (namespace)

# Vagrant commands:
	-vagrant up - запускает виртуальную машину
	-vagrant halt - останавливает виртуальную машину
	-vagrant destroy - удаляет виртуальную машину
	-vagrant suspend - "замораживает" виртуальную машину
	-vagrant global-status - выводит список всех ранее созданных виртуальных машин в хост-системе
	-vagrant ssh - подключается к виртуальной машине по SSH
