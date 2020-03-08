# QEMU demo

```bash
qemu-img create -f qcow2 ubuntu.qcow2 16G

qemu-system-x86_64 \
    -enable-kvm \
    -m 2048 \
    -smp cpus=4 \
    -nic user,model=virtio \
    -drive file=ubuntu.qcow2,media=disk,if=virtio \
    -cdrom /home/pierrez/Downloads/ubuntu-19.10-desktop-amd64.iso
```