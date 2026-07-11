cpu
Architecture:                x86_64
  CPU op-mode(s):            32-bit, 64-bit
  Address sizes:             39 bits physical, 48 bits virtual
  Byte Order:                Little Endian
CPU(s):                      4
  On-line CPU(s) list:       0-3
Vendor ID:                   GenuineIntel
  Model name:                Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
    CPU family:              6
    Model:                   60
    Thread(s) per core:      1
    Core(s) per socket:      4
    Socket(s):               1
    Stepping:                3
    CPU(s) scaling MHz:      57%
    CPU max MHz:             4000.0000
    CPU min MHz:             800.0000
    BogoMIPS:                6999.30
    Flags:                   fpu vme de pse tsc msr pae mce cx8 apic sep m
                             trr pge mca cmov pat pse36 clflush dts acpi m
                             mx fxsr sse sse2 ss ht tm pbe syscall nx pdpe
                             1gb rdtscp lm constant_tsc arch_perfmon pebs 
                             bts rep_good nopl xtopology nonstop_tsc cpuid
                              aperfmperf pni pclmulqdq dtes64 monitor ds_c
                             pl vmx est tm2 ssse3 sdbg fma cx16 xtpr pdcm 
                             pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_de
                             adline_timer aes xsave avx f16c rdrand lahf_l
                             m abm cpuid_fault pti ssbd ibrs ibpb stibp tp
                             r_shadow flexpriority ept vpid ept_ad fsgsbas
                             e tsc_adjust bmi1 avx2 smep bmi2 erms invpcid
                              xsaveopt dtherm ida arat pln pts vnmi md_cle
                             ar flush_l1d
Virtualization features:     
  Virtualization:            VT-x
Caches (sum of all):         
  L1d:                       128 KiB (4 instances)
  L1i:                       128 KiB (4 instances)
  L2:                        1 MiB (4 instances)
  L3:                        6 MiB (1 instance)
NUMA:                        
  NUMA node(s):              1
  NUMA node0 CPU(s):         0-3
Vulnerabilities:             
  Gather data sampling:      Not affected
  Ghostwrite:                Not affected
  Indirect target selection: Not affected
  Itlb multihit:             KVM: Mitigation: Split huge pages
  L1tf:                      Mitigation; PTE Inversion; VMX conditional ca
                             che flushes, SMT disabled
  Mds:                       Mitigation; Clear CPU buffers; SMT disabled
  Meltdown:                  Mitigation; PTI
  Mmio stale data:           Not affected
  Old microcode:             Not affected
  Reg file data sampling:    Not affected
  Retbleed:                  Not affected
  Spec rstack overflow:      Not affected
  Spec store bypass:         Mitigation; Speculative Store Bypass disabled
                              via prctl
  Spectre v1:                Mitigation; usercopy/swapgs barriers and __us
                             er pointer sanitization
  Spectre v2:                Mitigation; Retpolines; IBPB conditional; IBR
                             S_FW; STIBP disabled; RSB filling; PBRSB-eIBR
                             S Not affected; BHI Not affected
  Srbds:                     Mitigation; Microcode
  Tsa:                       Not affected
  Tsx async abort:           Not affected
  Vmscape:                   Mitigation; IBPB before exit to userspace

ram
              total        used        free      shared  buff/cache   available
Mem:            15Gi       2.0Gi        10Gi        26Mi       3.7Gi        13Gi
Swap:          4.0Gi          0B       4.0Gi


storage

NAME   MAJ:MIN RM   SIZE RO TYPE MOUNTPOINTS
loop0    7:0    0 245.1M  1 loop /snap/firefox/6565
loop1    7:1    0  11.1M  1 loop /snap/firmware-updater/167
loop2    7:2    0   516M  1 loop /snap/gnome-42-2204/202
loop3    7:3    0     4K  1 loop /snap/bare/5
loop4    7:4    0  73.9M  1 loop /snap/core22/2045
loop5    7:5    0  91.7M  1 loop /snap/gtk-common-themes/1535
loop6    7:6    0  10.8M  1 loop /snap/snap-store/1270
loop7    7:7    0  49.3M  1 loop /snap/snapd/24792
loop8    7:8    0   576K  1 loop /snap/snapd-desktop-integration/315
loop9    7:9    0  63.8M  1 loop /snap/core20/2866
loop10   7:10   0 450.6M  1 loop /snap/code/249
sda      8:0    0   1.8T  0 disk 
├─sda1   8:1    0     1G  0 part /boot/efi
└─sda2   8:2    0   1.8T  0 part /


lspci

00:00.0 Host bridge: Intel Corporation 4th Gen Core Processor DRAM Controller (rev 06)
00:01.0 PCI bridge: Intel Corporation Xeon E3-1200 v3/4th Gen Core Processor PCI Express x16 Controller (rev 06)
00:14.0 USB controller: Intel Corporation 9 Series Chipset Family USB xHCI Controller
00:16.0 Communication controller: Intel Corporation 9 Series Chipset Family ME Interface #1
00:1a.0 USB controller: Intel Corporation 9 Series Chipset Family USB EHCI Controller #2
00:1b.0 Audio device: Intel Corporation 9 Series Chipset Family HD Audio Controller
00:1c.0 PCI bridge: Intel Corporation 9 Series Chipset Family PCI Express Root Port 1 (rev d0)
00:1c.3 PCI bridge: Intel Corporation 9 Series Chipset Family PCI Express Root Port 4 (rev d0)
00:1d.0 USB controller: Intel Corporation 9 Series Chipset Family USB EHCI Controller #1
00:1f.0 ISA bridge: Intel Corporation Z97 Chipset LPC Controller
00:1f.2 SATA controller: Intel Corporation 9 Series Chipset Family SATA Controller [AHCI Mode]
00:1f.3 SMBus: Intel Corporation 9 Series Chipset Family SMBus Controller
01:00.0 VGA compatible controller: NVIDIA Corporation GM107 [GeForce GTX 750] (rev a2)
01:00.1 Audio device: NVIDIA Corporation GM107 High Definition Audio Controller [GeForce 940MX] (rev a1)
03:00.0 Ethernet controller: Qualcomm Atheros Killer E220x Gigabit Ethernet Controller (rev 13)


full system info

Static hostname: U-Server
       Icon name: computer-desktop
         Chassis: desktop 🖥️
      Machine ID: abfbf9290b114cd4a5eb269a416ec0d5
         Boot ID: 7d1f9638fcf6424f94ae7face3ff6ca5
Operating System: Ubuntu 24.04.4 LTS              
          Kernel: Linux 6.17.0-35-generic
    Architecture: x86-64
 Hardware Vendor: MSI
  Hardware Model: MS-7917
Firmware Version: V1.8
   Firmware Date: Thu 2014-11-06
    Firmware Age: 11y 8month 4d      