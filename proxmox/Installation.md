# Make bootable USB disk

## Locate your USB disk (macOS)

```shell
ll /dev/disk*
```

disk4 is the usb disk in my case.

## Write ISO to the usb disk

```shell
cd ~/Downloads
sudo dd if=proxmox-ve_8.4-1.iso bs=4M of=/dev/disk4 status=progress oflag=sync
```

## Troubleshooting for "resource busy"

Some partitions of the USB disk may have been mounted to your MacOS. Use *Dis Utility* to review and unmount them.
![](diskutil.png)