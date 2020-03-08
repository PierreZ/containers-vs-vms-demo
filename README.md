# containers-vs-vms-demo
All the demo for the `Containers vs VMs` lecture

## Cheatsheets

```bash

# download gentoo stage
./init.sh

# become root

# mount things
mount --types proc /proc ./mychroot/proc
mount --rbind /sys ./mychroot/sys
mount --make-rslave ./mychroot/sys
mount --rbind /dev ./mychroot/dev
mount --make-rslave ./mychroot/dev

# chroot
chroot ./mychroot /bin/bash

# unmount things
umount -l ./mychroot/{proc,sys}
umount -l ./mychroot/sys
umount -l ./mychroot/dev

# systemd-nspawn
systemd-nspawn -D mychroot

# dockerlite
cd dockerlite
./dockerlite run /bin/bash
```
